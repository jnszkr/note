package searcher

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/jnszkr/note/internal/files"

	"github.com/jnszkr/note/internal/reader"

	"github.com/jnszkr/note/internal/formatter"
)

const FileName = ".notes"

type Searcher interface {
	// Search finds all the files that are called `.notes` in the current
	// path recursively and tries to find the expression in each one.
	// The results are written to io.Writer.
	Search(s string, recursive bool)
}

func New(path string, out io.Writer) Searcher {
	return &searcher{
		path: path,
		out:  out,
		done: make(chan struct{}),
	}
}

type searcher struct {
	path  string
	out   io.Writer
	stats stats
	done  chan struct{}
}

type stats struct {
	numberOfFiles int
}

func (s *searcher) Search(exp string, recursive bool) {
	ts := time.Now()
	defer func() {
		if os.Getenv("DEBUG") != "" {
			fmt.Printf("Files found: %d\n", s.stats.numberOfFiles)
			fmt.Printf("Time       : %v\n", time.Since(ts))
		}
	}()

	var fs <-chan string
	if recursive {
		fs = files.RecursiveFind(s.path, FileName)
	} else {
		fs = files.Find(s.path, FileName)
	}

	exp = strings.ToLower(exp)

	resultChan := make(chan string, 10)
	go func() {
		for res := range resultChan {
			fmt.Fprint(s.out, res)
		}
		close(s.done)
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
	<-s.done
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
