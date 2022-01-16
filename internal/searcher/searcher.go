package searcher

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/jnszkr/note/internal/reader"

	"github.com/jnszkr/note/internal/formatter"
)

const FileName = ".notes"

var debug = false

func init() {
	debug = os.Getenv("DEBUG") != ""
}

type Searcher interface {
	Search(s string, recursive bool)
}

func New(path string, out io.Writer) Searcher {
	return &searcher{
		path: path,
		out:  out,
	}
}

type searcher struct {
	path  string
	out   io.Writer
	stats stats
}

type stats struct {
	numberOfFiles int
}

// Search finds all the files that are called `.notes` in the current
// path recursively and tries to find the expression in each one.
// The results are written to io.Writer.
func (s *searcher) Search(exp string, recursive bool) {
	ts := time.Now()
	defer func() {
		if debug {
			fmt.Printf("Files found: %d\n", s.stats.numberOfFiles)
			fmt.Printf("Time       : %v\n", time.Since(ts))
		}
	}()

	fs := s.files(recursive)

	exp = strings.ToLower(exp)

	resultChan := make(chan string, 10)
	printDone := make(chan struct{})
	go func() {
		for res := range resultChan {
			fmt.Fprint(s.out, res)
		}
		close(printDone)
	}()

	wg := &sync.WaitGroup{}
	for path := range fs {
		s.stats.numberOfFiles++

		wg.Add(1)
		go func(path string) {
			defer wg.Done()

			res, err := searchIn(path, exp)
			if err != nil {
				log.Fatal(err)
			}
			if len(res) > 0 {
				resultChan <- fmt.Sprintf("%s\n%s", s.topicDisplay(path), res)
			}
		}(path)
	}
	wg.Wait()
	close(resultChan)
	<-printDone
}

func (s *searcher) topicDisplay(path string) string {
	re := regexp.MustCompilePOSIX(s.path + "/(.*)/" + FileName)
	subs := re.FindAllStringSubmatch(path, -1)
	if subs == nil {
		return " • "
	}
	topics := strings.Split(subs[0][1], "/")
	return " • " + strings.Join(topics, " • ")
}

func searchIn(path string, s string) (string, error) {
	r, err := reader.NewWith(path, func(t *time.Time, note string) bool {
		return strings.Contains(strings.ToLower(note), s)
	})
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			return "", nil // return with nil if file does not exist
		}
		return "", err
	}
	defer r.Close()

	res := formatter.FormatWith(r, "   ")
	return formatter.Highlight(res, s, formatter.Red), nil
}

// ignoredFiles
var ignoredFiles = map[string]struct{}{
	".git": {},
}

const fsChanSize = 100

func (s *searcher) files(recursive bool) <-chan string {
	fsChan := make(chan string, fsChanSize)

	if recursive {
		go func() {
			err := s.walk(fsChan)
			if err != nil {
				log.Fatal(err)
			}
			close(fsChan)
		}()
		return fsChan
	}

	p, err := filepath.Abs(filepath.Join(s.path, FileName))
	if err != nil {
		log.Fatal(err)
	}
	fsChan <- p
	close(fsChan)

	return fsChan
}

func (s *searcher) walk(fsChan chan string) error {
	err := filepath.Walk(s.path, func(path string, f os.FileInfo, err error) error {
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
		case !f.IsDir() && f.Name() == FileName:
			fsChan <- path
		}

		return nil
	})
	return err
}
