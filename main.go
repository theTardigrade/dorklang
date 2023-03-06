package main

import (
	"os"
)

func main() {
	fileContents, err := os.ReadFile(*fileFlag)
	if err != nil {
		panic(err)
	}

	lexemes, err := produceLexemes(fileContents)
	if err != nil {
		panic(err)
	}

	tree, err := produceTree(lexemes)
	if err != nil {
		panic(err)
	}

	if err = tree.Run(); err != nil {
		panic(err)
	}
}
