package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// search finds all the files that are called `.notes` in the current
// path recursively and tries to find the expression in each one.
// The results are printed to standard out.
func search(s string) {
	fs, err := files()
	if err != nil {
		log.Fatal(err)
	}
	for _, path := range fs {
		res, err := searchIn(path, s)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(path)
		for _, l := range res {
			fmt.Printf("\t%s\n", l)
		}
	}
}

func searchIn(path string, s string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	s = strings.ToLower(s)

	res := make([]string, 0)
	line := 1
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), s) {
			res = append(res, scanner.Text())
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		return res, err
	}
	return res, nil
}

var ignoredFiles = map[string]struct{}{
	".git": {},
}

func files() ([]string, error) {
	var fs []string

	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		_, ignored := ignoredFiles[f.Name()]
		switch {
		case ignored:
			return filepath.SkipDir
		case !f.IsDir() && f.Name() == ".notes":
			fs = append(fs, path)
		}

		return nil
	})

	return fs, err
}
