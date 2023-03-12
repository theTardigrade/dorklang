package main

const (
	treeSaveFileExtension = ".dork.txt"
	treeSaveStackMaxLen   = 1 << 20 // 1_048_576
)

type tree struct {
	rootNode       *parentTreeNode
	saveStackIndex int
	saveStacks     [2]memoryCellCollection
}

type treeNode interface {
	getTree() *tree
	getData() []byte
	value(memoryCell) (memoryCell, error)
}

type defaultTreeNode struct {
	lexeme lexeme
	data   []byte
	tree   *tree
}

type parentTreeNode struct {
	defaultTreeNode
	childNodes []treeNode
}

type terminalTreeNode struct {
	defaultTreeNode
	parentNode *parentTreeNode
}
