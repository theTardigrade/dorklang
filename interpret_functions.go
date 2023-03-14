package dorklang

import (
	"os"
)

func InterpretCode(input []byte, options InterpretCodeOptions) (output uint64, err error) {
	initialDir, err := os.Getwd()
	if err != nil {
		return
	}

	err = os.Chdir(options.WorkingDir)
	if err != nil {
		return
	}

	tokens, err := produceTokens(input)
	if err != nil {
		return
	}

	if !options.SkipClean {
		if err = cleanTokens(tokens); err != nil {
			return
		}
	}

	if options.DebugMode {
		tokens.log()
	}

	tree, err := produceTree(tokens, options)
	if err != nil {
		return
	}

	outputMemoryCell, err := tree.Run()
	if err != nil {
		return
	}
	output = outputMemoryCell.Uint64()

	err = os.Chdir(initialDir)
	if err != nil {
		return
	}

	return
}

func InterpretCodeWithDefaultOptions(input []byte) (output uint64, err error) {
	output, err = InterpretCode(input, InterpretCodeDefaultOptions)

	return
}
