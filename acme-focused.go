package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"9fans.net/go/acme"
)

// name of the text file generated
const logName = "acme-focused"

type focusedWin struct {
	id int
	mu sync.Mutex
}

// open acme's log and save it's $winid to fw.id everytime a "focus" event happens
func (fw *focusedWin) readLog() {
	alog, err := acme.Log()
	if err != nil {
		log.Fatalf("failed to open acmelog: %v\n", err)
	}
	defer alog.Close()
	for {
		time.Sleep(2 * time.Second)
		ev, err := alog.Read()
		if err != nil {
			log.Fatalf("failed to read log: %v\n", err)
		}
		if ev.Op == "focus" {
			fw.mu.Lock()
			fw.id = ev.ID
			fw.mu.Unlock()
		}
	}
}

// makes the final path of the temporary file
func makeFilePath(path *string) {
	sepIndex := strings.LastIndex(*path, "/")
	if sepIndex != len(*path)-1 {
		*path += "/"
	}
	*path += logName
}

// returns the current window ID
func (fw *focusedWin) ID() int {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	return fw.id
}

// func writeId(path string, id int) {
func writeId(path string, fw *focusedWin) {
	lastVal := 0
	makeFilePath(&path)
	for {

		if lastVal != fw.ID() {
			lastVal = fw.ID()

			/*
			   use os.WriteFile because it overwrites the previous text entered,
			   as opposed to a, say, os.Create and fmt.Fprintf(), which would
			   just append strings to the previous data entered
			*/
			err := os.WriteFile(path, []byte(strconv.Itoa(fw.ID())+"\n"), 0666)
			if err != nil {
				log.Fatalf("couldn't open/write file at '%s': %s", path, err)
			}
		}
	}
}

func usage() {
	fmt.Printf("Usage:\nacme-focused [-h, --h] [path]\n")
	fmt.Printf("-h, --h: prints this message\n")
	fmt.Printf("path: directory to store the file, defaults to /tmp/ if left blank\n")
	os.Exit(1)
}

func main() {
	var fw focusedWin
	var path string

	argLen := len(os.Args)
	if argLen > 2 {
		usage()
	} else if argLen <= 1 {
		path = os.TempDir()
	} else if os.Args[1] == "-h" || os.Args[1] == "--h" {
		usage()
	} else {
		path = os.Args[1]
	}
	_, err := os.ReadDir(path) // check if directory exists
	if err != nil {
		log.Fatalf("failed to open directory: '%s'\n", err)
	}

	go fw.readLog()

	writeId(path, &fw)
}
