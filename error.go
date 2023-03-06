package main

import "errors"

var (
	ErrNoMatchSectionCharacters = errors.New("the number of starting and ending characters for sections do not match")
	ErrOverlapSectionCharacters = errors.New("the starting and ending characters for sections overlap incorrectly")
	ErrTreeParentNodeUnfound    = errors.New("cannot find parent node for tree")
	ErrTreeUnfound              = errors.New("cannot find tree")
	ErrTreeSaveStackFull        = errors.New("cannot save any more values in the tree stack")
	ErrTreeSaveStackEmpty       = errors.New("cannot load a value from the tree stack")
	ErrLexemeUnrecognized       = errors.New("lexeme is not recognized")
)
