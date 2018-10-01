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
		t, _, err := load(f)
		if err != nil {
			fmt.Printf("load: fatal: %s\n", err)
			os.Exit(1)
		}
		editor := newEditor(t)
		editor.run()
	}
}

type editor struct {
	path nbt.Path
	cmds map[string]handler
}

type handler func([]string) error

func newEditor(t nbt.Tag) *editor {
	e := &editor{path: nbt.NewPath(t)}
	e.cmds = make(map[string]handler)
	e.cmds["ls"] = e.doLs
	e.cmds["cd"] = e.doCd
	return e
}

func (e *editor) prompt() {
	fmt.Printf("%v > ", e.path)
}

func (e *editor) doLs(args []string) error {
	if len(args) == 0 {
		listContents(e.path.Current())
		return nil
	}
	return nil
}

func (e *editor) doCd(args []string) error {
	_, err := e.path.Cd(nbt.String(args[0]))
	return err
}

func (e *editor) input() (string, error) {
	buf := make([]byte, 512)
	n, err := os.Stdin.Read(buf)
	if n != 0 {
		return string(buf[0:n]), nil
	}
	return "", err
}

func (e *editor) run() {
	// do editing
	gotEOF := false
	for {
		e.prompt()
		cmd, err := e.input()
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
		fn, ok := e.cmds[cmd]
		if ok {
			err = fn(args)
			if err != nil {
				fmt.Printf("cmd failed: %s\n", err)
			}
		} else {
			fmt.Printf("unknown cmd %q\n", cmd)
		}
		if gotEOF {
			break
		}
	}
}

func listContents(t nbt.Tag) {
	switch tag := t.(type) {
	case nbt.List:
		tag.Iterate(func(i int, t nbt.Tag) error { fmt.Printf("%d\n", i); return nil })
	case nbt.Compound:
		for k := range tag {
			fmt.Printf("%s\n", k)
		}
	default:
		fmt.Printf("unlistable node %v\n", t.Type())
	}
}
