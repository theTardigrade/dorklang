package main

import (
	"os"

	"github.com/theTardigrade/dorklang"
)

func main() {
	fileContents, err := os.ReadFile(*flagFile)
	if err != nil {
		panic(err)
	}

	output, err := dorklang.InterpretCode(fileContents, dorklang.InterpretCodeOptions{
		DebugMode: *flagDebug,
		SkipClean: *flagSkipClean,
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
