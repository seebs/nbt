package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/seebs/gogetopt"
	"github.com/seebs/nbt"
	"github.com/seebs/quotable"
)

func main() {
	opts, files, err := gogetopt.GetOpt(os.Args[1:], "u")
	if err != nil {
		log.Fatalf("invalid args: %s", err)
	}

	if len(files) != 1 {
		log.Fatalf("usage: nbtdb [-u] file")
	}

	load := nbt.Load
	if opts.Seen("u") {
		load = nbt.LoadUncompressed
	}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("open: fatal: %s\n", err)
			os.Exit(1)
		}
		t, err := load(f)
		if err != nil {
			fmt.Printf("load: fatal: %s\n", err)
			os.Exit(1)
		}
		editor(t)
	}
}

var tagStack []nbt.Tag

func prompt() {
	for _, p := range tagStack {
		fmt.Printf("%q > ", p.Name)
	}
}

func editor(t nbt.Tag) {
	doEdit(t)
}

type handler func([]string) error

var cmds = map[string]handler{
	"ls": doLs,
}

func doLs(args []string) error {
	if len(args) == 0 {
		listContents(tagStack[len(tagStack)-1])
		return nil
	}
	return nil
}

func listContents(t nbt.Tag) {
	switch t.Type {
	case nbt.TypeList:
		_, ok := t.GetList()
		if !ok {
			fmt.Printf("not a list\n")
		}
	case nbt.TypeCompound:
		c, ok := t.GetCompound()
		if !ok {
			fmt.Printf("not a compound\n")
		}
		for k := range c {
			fmt.Printf("%s\n", k)
		}
	default:
		fmt.Printf("unlistable node %v", t.Type)
	}
}

// doEdit runs the actual editing interface
func doEdit(t nbt.Tag) {
	// add tag to stack
	tagStack = append(tagStack, t)
	defer func() {
		// and remove the tag when we're done
		if len(tagStack) > 0 {
			tagStack = tagStack[:len(tagStack)-1]
		}
	}()
	// do editing
	gotEOF := false
	for {
		prompt()
		cmd, err := input()
		if err == io.EOF {
			gotEOF = true
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "unexpected error from input: %s\n", err)
			return
		}
		if cmd == "" {
			break
		}
		words, err := quotable.Split(cmd, &quotable.Options{FancyBackslash: true})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error in split: %s [ignored]\n", err)
		}
		var args []string
		cmd = words[0]
		if len(words) > 1 {
			args = words[1:]
		}
		fmt.Printf("%s %v\n", cmd, args)
		fn, ok := cmds[cmd]
		if ok {
			err = fn(args)
			if err != nil {
				fmt.Printf("cmd failed: %s\n", err)
				break
			}
		}
		if gotEOF {
			break
		}
	}
}

func input() (string, error) {
	buf := make([]byte, 512)
	n, err := os.Stdin.Read(buf)
	if n != 0 {
		return string(buf[0:n]), nil
	}
	return "", err
}
