package files

import (
	"log"
	"os"
	"path/filepath"
)

// ignoredFiles
var ignoredFiles = map[string]struct{}{
	".git": {},
}

const fsChanSize = 100

func RecursiveFind(path, filename string) <-chan string {
	fsChan := make(chan string, fsChanSize)

	go func() {
		err := walk(path, filename, ignoredFiles, fsChan)
		if err != nil {
			log.Fatal(err)
		}
		close(fsChan)
	}()

	return fsChan
}

func Find(path, filename string) <-chan string {
	fsChan := make(chan string, 1)

	p, err := filepath.Abs(filepath.Join(path, filename))
	if err != nil {
		log.Fatal(err)
	}
	fsChan <- p
	close(fsChan)

	return fsChan
}

func walk(path, filename string, ignoredFiles map[string]struct{}, fsChan chan string) error {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if err != nil || f == nil {
			return filepath.SkipDir
		}
		if f.IsDir() {
			return nil
		}
		_, ignored := ignoredFiles[f.Name()]
		switch {
		case ignored:
			return filepath.SkipDir
		case f.Name() == filename:
			fsChan <- path
		}

		return nil
	})
	return err
}
