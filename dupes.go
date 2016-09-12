
package main

import (
	"bufio";
	"fmt";
	"os";
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func countLinesInFile (files []string, counts map[string]int) {
	for _, arg := range files {
        f, err := os.Open(arg)

        if err != nil {
            fmt.Fprintf(os.Stderr, "Error file %s: %v\n", f, err)
			continue
		}

        countLines(f, counts)
		f.Close()
    }
}

func printCounts(counts map[string]int) {
	for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
	}
}

func main () {
    files := os.Args[1:]
    counts := make(map[string]int)
    if len(files) == 0 {
        countLines(os.Stdin, counts)

    } else {
    	countLinesInFile(files, counts)
    }

    printCounts(counts)
}