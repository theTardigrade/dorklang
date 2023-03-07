package main

import "flag"

const (
	fileExtension = ".dork"
)

var (
	fileFlag  = flag.String("file", "source"+fileExtension, "the path to the source file")
	flagDebug = flag.Bool("debug", false, "determines whether to print debug information")
)

func init() {
	flag.Parse()
}
