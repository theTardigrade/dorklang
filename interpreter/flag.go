package main

import (
	"flag"

	"github.com/theTardigrade/dorklang"
)

var (
	flagFile           = flag.String("file", "source"+dorklang.FileExtensionForCode, "the path to the source file")
	flagDebug          = flag.Bool("debug", false, "determines whether to print debug information")
	flagSkipClean      = flag.Bool("skip-clean", false, "determines whether to skip the cleaning-tokens stage")
	flagSkipExitStatus = flag.Bool("skip-exit-status", false, "determines whether to skip basing the program's exit code on its final current value")
)

func init() {
	flag.Parse()
}
