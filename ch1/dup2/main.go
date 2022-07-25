// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named args.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type linemap = map[string]map[string]bool

func main() {
	lines := make(linemap)
	args := os.Args[1:]
	if len(args) == 0 {
		groupFilesByLine(os.Stdin, lines, "<stdin>")
	} else {
		for _, arg := range args {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			groupFilesByLine(f, lines, arg)
			f.Close()
		}
	}
	for line, files := range lines {
		if len(files) > 1 {
			fmt.Printf("%s\n", line)
			for match := range files {
				fmt.Printf("\t%s\n", match)
			}
		}
	}
}

func groupFilesByLine(f *os.File, lines linemap, file string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		files, ok := lines[line]
		if !ok {
			files = make(map[string]bool)
			lines[line] = files
		}
		files[file] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}
