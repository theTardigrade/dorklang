package main

import (
	"os"
	"path/filepath"

	"github.com/theTardigrade/dorklang"
)

func main() {
	fileAbsPath, err := filepath.Abs(*flagFile)
	if err != nil {
		panic(err)
	}

	fileContents, err := os.ReadFile(fileAbsPath)
	if err != nil {
		panic(err)
	}

	output, err := dorklang.InterpretCode(fileContents, dorklang.InterpretCodeOptions{
		WorkingDir: filepath.Dir(fileAbsPath),
		DebugMode:  *flagDebug,
		SkipClean:  *flagSkipClean,
		Input:      os.Stdin,
		Output:     os.Stdout,
	})
	if err != nil {
		panic(err)
	}
	if !*flagSkipExitStatus {
		if output <= 124 {
			os.Exit(int(output))
		} else {
			os.Exit(125)
		}
	}
}
