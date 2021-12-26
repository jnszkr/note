package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

func display(path string) string {
	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return ""
		}
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	res := make([]string, 0)
	for scanner.Scan() {
		t := scanner.Text()
		res = append(res, t)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return strings.Join(res, "\n")
}
