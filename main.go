package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/evilsocket/wax/schema"
)

var (
	err      error
	docPath  string
	doc      schema.Document
	descPath string
	desc     schema.Descriptor
)

func init() {
	flag.StringVar(&docPath, "document", "", "Path of the document to use as a baseline template.")
	flag.StringVar(&descPath, "schema", "", "Path of the schema descriptor of the input document.")
	flag.Parse()
}

func fatal(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func precheck() {
	if docPath == "" {
		fatal("no --document specified")
	} else if descPath == "" {
		fatal("no --schema specified")
	}
}

func main() {
	precheck()

	if doc, err = schema.LoadDocument(docPath); err != nil {
		fatal("error while loading document %s: %s\n", docPath, err)
	} else if desc, err = schema.LoadDescriptor(descPath); err != nil {
		fatal("error while loading schema descriptor %s: %s\n", descPath, err)
	}

	fmt.Printf("all good buddy\n")
}
