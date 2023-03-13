package dorklang

import (
	"log"
)

func InterpretCode(input []byte, options InterpretCodeOptions) (output uint64, err error) {
	tokens, err := produceTokens(input)
	if err != nil {
		return
	}

	if !options.SkipClean {
		cleanTokens(tokens)
	}

	if options.DebugMode {
		for _, t := range tokens {
			log.Printf("lexeme: %s\n", t.lex)
		}
	}

	tree, err := produceTree(tokens)
	if err != nil {
		return
	}

	outputMemoryCell, err := tree.Run()
	if err != nil {
		return
	}
	output = outputMemoryCell.Uint64()

	return
}

func InterpretCodeWithDefaultOptions(input []byte) (output uint64, err error) {
	output, err = InterpretCode(input, InterpretCodeDefaultOptions)

	return
}
