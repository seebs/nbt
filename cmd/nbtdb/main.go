package main

import (
	"fmt"
	"os"
	"github.com/seebs/nbt"
)

func main() {
	f, err := os.Open("/tmp/player.dat")
	if err != nil {
		fmt.Printf("fatal: %s\n", err)
		os.Exit(1)
	}
	n, err := nbt.Load(f)
	if err != nil {
		fmt.Printf("fatal: %s\n", err)
		os.Exit(1)
	}
	n.PrintIndented(os.Stdout)
}
