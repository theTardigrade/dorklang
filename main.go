package main

import (
	"log"
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

	if *flagDebug {
		for _, l := range lexemes {
			log.Printf("lexeme: %s\n", l)
		}
	}

	tree, err := produceTree(lexemes)
	if err != nil {
		panic(err)
	}

	if err = tree.Run(); err != nil {
		panic(err)
	}
}
