package dorklang

type tree struct {
	rootNode             *parentTreeNode
	interpretCodeOptions InterpretCodeOptions
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
