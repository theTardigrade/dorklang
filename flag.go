package main

import "flag"

var (
	fileFlag        = flag.String("file", "source"+fileExtension, "the path to the source file")
	flagDebug       = flag.Bool("debug", false, "determines whether to print debug information")
	flagIgnoreClean = flag.Bool("ignore-clean", false, "determines whether to skip the cleaning-tokens stage")
)

func init() {
	flag.Parse()
}
