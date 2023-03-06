package main

import "flag"

var (
	fileFlag = flag.String("file", "source.eso", "the path to the source file")
)

func init() {
	flag.Parse()
}
