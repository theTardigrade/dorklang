package dorklang

import (
	"io"
	"os"
)

type InterpretCodeOptions struct {
	WorkingDir string
	DebugMode  bool
	SkipClean  bool
	Input      io.Reader
	Output     io.Writer
}

var (
	InterpretCodeDefaultOptions = InterpretCodeOptions{
		DebugMode: false,
		SkipClean: false,
		Input:     os.Stdin,
		Output:    os.Stdout,
	}
)

func (options InterpretCodeOptions) Clone() InterpretCodeOptions {
	return InterpretCodeOptions{
		WorkingDir: options.WorkingDir,
		DebugMode:  options.DebugMode,
		SkipClean:  options.SkipClean,
		Input:      options.Input,
		Output:     options.Output,
	}
}
