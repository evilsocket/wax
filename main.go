package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/evilsocket/wax/schema"
)

var (
	err      error
	docPath  string
	doc      schema.Document
	descPath string
	desc     schema.Descriptor
)

func adios(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func init() {
	flag.StringVar(&docPath, "document", "", "Path of the document to use as a baseline template.")
	flag.StringVar(&descPath, "schema", "", "Path of the schema descriptor of the input document.")
	flag.Parse()

	if docPath == "" {
		adios("no --document specified")
	} else if descPath == "" {
		adios("no --schema specified")
	}
}

func main() {
	if doc, err = schema.LoadDocument(docPath); err != nil {
		adios("error while loading document %s: %s\n", docPath, err)
	} else if desc, err = schema.LoadDescriptor(descPath); err != nil {
		adios("error while loading schema descriptor %s: %s\n", descPath, err)
	}

	start := time.Now()
	if err = desc.Test(doc); err != nil {
		adios("%v\n", err)
	}
	elapsed := time.Now().Sub(start)

	fmt.Printf("loaded %d bytes of document raw data\n", doc.Size)
	fmt.Printf("loaded schema `%s` (%d atoms tested in %dns)\n", desc.Name, len(desc.Atoms), elapsed/time.Nanosecond)

	fmt.Println()
	fmt.Printf("-------------------------------------\n%s\n-------------------------------------\n", doc.Data)
	for _, atom := range desc.Atoms {
		off, size, _ := atom.Locator.Find(doc)
		fmt.Printf("%s\n  %s[%d:%d] -> '%s'\n", atom.Name, atom.Type, off, size, doc.Data[off:off+size])
	}
}
