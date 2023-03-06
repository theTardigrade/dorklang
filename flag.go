package main

import "flag"

const (
	fileExtension = ".dork"
)

var (
	fileFlag = flag.String("file", "source"+fileExtension, "the path to the source file")
)

func init() {
	flag.Parse()
}
