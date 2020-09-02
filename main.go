package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/getwild"
)

func mains(args []string) error {
	items, err := readXmlFiles(args)
	if err != nil {
		return err
	}
	for _, item1 := range items {
		println("###", shrink(item1.Case))
		println()
		println(shrink(item1.Operation))
		if item1.Status != "" {
			println()
			println(item1.Status)
		}
		println()
	}
	return nil
}

func main() {
	if err := mains(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
