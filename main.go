package main

import (
	"fmt"
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
	fmt.Println(lexemes)

	tree, err := produceTree(lexemes)
	if err != nil {
		panic(err)
	}

	if err = tree.Run(); err != nil {
		panic(err)
	}
}
