package main

import (
	"fmt"
	"os"
)

func main() {
	sourceFileContents, err := os.ReadFile(*fileFlag)
	if err != nil {
		panic(err)
	}

	lexemes, err := produceLexemes(sourceFileContents)
	if err != nil {
		panic(err)
	}
	fmt.Println(lexemes)

	tree, err := produceTree(lexemes)
	if err != nil {
		panic(err)
	}

	if err = tree.Run(); err != nil {
		panic(err)
	}
}
