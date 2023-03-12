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

	tokens, err := produceTokens(fileContents)
	if err != nil {
		panic(err)
	}

	if !*flagSkipClean {
		cleanTokens(tokens)
	}

	if *flagDebug {
		for _, t := range tokens {
			log.Printf("lexeme: %s\n", t.lex)
		}
	}

	tree, err := produceTree(tokens)
	if err != nil {
		panic(err)
	}

	if err = tree.Run(); err != nil {
		panic(err)
	}
}
