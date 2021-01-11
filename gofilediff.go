package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	// Version number
	VERSION = "0.1"
)

// command line arguments
var (
	flagVersion bool = false
)

func version() {
	fmt.Printf("Version %s\n", VERSION)
}

func usage() {
	fmt.Fprint(os.Stderr, "A text file comparison to remove fileToClean lines already contained in fileMatch. Result will be saved in fileToClean.cleaned\n\n")
	fmt.Fprint(os.Stderr, "usage: gofilediff <options> <fileMatch> <fileToClean>\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func openFile() {
	file, err := os.Open("file.go") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	_ = file
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}

func readFile(filename string) []string {
	// check file type
	_, err := os.Stat(filename)
	check(err)

	// Read the file
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())

	return lines
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) {
	file, err := os.Create(path)
	check(err)

	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	check(w.Flush())
}

func cleanup(f1Lines []string, f2Lines []string) []string {

	var cleannedFile []string
	for _, line := range f2Lines {
		if !stringInSlice(line, f1Lines) {
			cleannedFile = append(cleannedFile, line)
		}
	}
	return cleannedFile
}

// check if a string is in the slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {
	flag.Usage = usage
	flag.BoolVar(&flagVersion, "v", flagVersion, "Print version information")
	flag.Parse()

	if flagVersion {
		version()
		os.Exit(0)
	}

	// get command line args
	args := flag.Args()
	if len(args) == 0 {
		usage()
		os.Exit(0)
	}

	if len(args) < 2 {
		log.Printf("Missing files")
		os.Exit(1)
	}

	if len(args) > 2 {
		log.Printf("Too many files")
		os.Exit(1)
	}

	// Read the files
	fileMatch := readFile(args[0])
	fileToClean := readFile(args[1])

	// Remove fileToClean line's already contained in fileMatch and return cleaned result
	cleannedFile := cleanup(fileMatch, fileToClean)

	// write the cleaned result
	writeLines(cleannedFile, fmt.Sprintf("%s.%s", args[1], "cleaned"))
}
