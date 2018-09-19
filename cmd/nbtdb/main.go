package main

import (
	"fmt"
	"log"
	"os"
	"github.com/seebs/gogetopt"
	"github.com/seebs/nbt"
)

func main() {
	opts, files, err := gogetopt.GetOpt(os.Args[1:], "u")
	if err != nil {
		log.Fatalf("invalid args: %s", err)
	}
	load := nbt.Load
	if opts.Seen("u") {
		load = nbt.LoadUncompressed
	}


	for _, file := range(files) {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("open: fatal: %s\n", err)
			os.Exit(1)
		}
		n, err := load(f)
		if err != nil {
			fmt.Printf("load: fatal: %s\n", err)
			os.Exit(1)
		}
		n.PrintIndented(os.Stdout)
	}
}
