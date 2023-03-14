package dorklang

import (
	"io"
	"os"
)

const (
	interpretCodeOptionsSaveStackMaxLen = 1 << 20 // 1_048_576
)

type InterpretCodeOptions struct {
	WorkingDir          string
	DebugMode           bool
	SkipClean           bool
	Input               io.Reader
	Output              io.Writer
	initialCurrentValue memoryCell
	saveStackIndex      int
	saveStacks          [2]memoryCellCollection
}

var (
	InterpretCodeDefaultOptions = InterpretCodeOptions{
		DebugMode:           false,
		SkipClean:           false,
		Input:               os.Stdin,
		Output:              os.Stdout,
		initialCurrentValue: 0,
		saveStackIndex:      0,
	}
)

func (options InterpretCodeOptions) Clone() InterpretCodeOptions {
	return InterpretCodeOptions{
		WorkingDir:          options.WorkingDir,
		DebugMode:           options.DebugMode,
		SkipClean:           options.SkipClean,
		Input:               options.Input,
		Output:              options.Output,
		initialCurrentValue: options.initialCurrentValue,
		saveStackIndex:      options.saveStackIndex,
		saveStacks:          options.saveStacks,
	}
}
