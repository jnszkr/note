package searcher

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jnszkr/note/color"
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
	for _, path := range fs {
		res, err := searchIn(path, exp)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(s.out, s.topicDisplay(path))
		for _, l := range res {
			fmt.Fprintf(s.out, "\t%s\n", l)
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

func searchIn(path string, s string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	res := make([]string, 0)
	for scanner.Scan() {
		t := scanner.Text()
		if strings.Contains(strings.ToLower(t), s) {
			re := regexp.MustCompile(fmt.Sprintf("(?i)(%s)", s))
			t = re.ReplaceAllString(t, color.Red("$1"))
			res = append(res, t)
		}
	}

	if err := scanner.Err(); err != nil {
		return res, err
	}
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
