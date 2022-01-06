package searcher

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/jnszkr/note/internal/reader"

	"github.com/jnszkr/note/internal/formatter"
)

type Searcher interface {
	Search(s string)
}

func New(path string, out io.Writer) Searcher {
	return &searcher{
		path: path,
		out:  out,
	}
}

type searcher struct {
	path string
	out  io.Writer
}

// Search finds all the files that are called `.notes` in the current
// path recursively and tries to find the expression in each one.
// The results are written to io.Writer.
func (s *searcher) Search(exp string) {
	fs, err := s.files()
	if err != nil {
		log.Fatal(err)
	}

	exp = strings.ToLower(exp)

	for _, path := range fs {
		res, err := searchIn(path, exp)
		if err != nil {
			log.Fatal(err)
		}

		if len(res) > 0 {
			fmt.Fprintln(s.out, s.topicDisplay(path))
			fmt.Fprint(s.out, res)
		}
	}
}

func (s *searcher) topicDisplay(path string) string {
	re := regexp.MustCompilePOSIX(s.path + "/(.*)/.notes")
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
		return "", err
	}

	res := formatter.FormatWith(r, "   ")
	// TODO: add this to formatter
	re := regexp.MustCompile(fmt.Sprintf("(?i)(%s)", s))
	res = re.ReplaceAllString(res, formatter.Red("$1"))

	return res, nil
}

// ignoredFiles
var ignoredFiles = map[string]struct{}{
	".git": {},
}

func (s *searcher) files() ([]string, error) {
	var fs []string

	err := filepath.Walk(s.path, func(path string, f os.FileInfo, err error) error {
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
