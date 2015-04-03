package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"golang.org/x/net/html"
)

var (
	exitCode = 0
)

func main() {
	annotateMain()
	os.Exit(exitCode)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: annotate\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func annotateMain() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		processFile(".", nil)
	}
}

func processFile(filename string, in io.Reader) {
	z := html.NewTokenizer(in)
	fmt.Printf("%+v\n", z)
}

func isHtmlFile(f os.FileInfo) bool {
	name := f.Name()
	return !f.IsDir() && !strings.HasPrefix(name, ".") &&
		strings.HasSuffix(name, "html")
}
