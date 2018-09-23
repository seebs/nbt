// The taggen command exists to generate some generic code from templates, because
// there's some methods which can be generated entirely automatically from a
// single known word.
package main

import (
	"flag"
	"log"
	"os"
	"text/template"
)

func main() {
	infile := flag.String("in", "", "input file name (containing template)")
	outfile := flag.String("out", "", "output file name")
	flag.Parse()
	if infile == nil {
		log.Fatal("-in must be specified")
	}
	tmpl, err := template.ParseFiles(*infile)
	if err != nil {
		log.Fatalf("could not read input file '%s': %s", *infile, err)
	}
	output := os.Stdout
	if outfile != nil && *outfile != "" {
		output, err = os.Create(*outfile)
		if err != nil {
			log.Fatalf("could not open output file '%s': %s", *outfile, err)
		}
		defer output.Close()
	}
	err = tmpl.Execute(output, flag.Args())
	if err != nil {
		if outfile != nil && *outfile != "" {
			os.Remove(*outfile)
		}
		log.Fatalf("template failed: %s", err)
	}
}
