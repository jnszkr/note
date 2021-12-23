package main

import (
	"bytes"
	"log"
	"os"
	"time"
)

func add(args []string) {
	// if the file does not exist, create it, or append to the file
	f, err := os.OpenFile(".notes", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// read input
	b := bytes.Buffer{}
	b.WriteString(time.Now().Local().Format(time.RFC3339))
	b.WriteString(" ")
	for _, arg := range args {
		b.WriteString(arg)
		b.WriteString(" ")
	}
	b.WriteString("\n")

	// append input
	_, err = f.Write(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	// close file
	f.Close()
}
