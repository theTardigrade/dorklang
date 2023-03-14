package dorklang

func produceTree(input tokenCollection, interpretCodeOptions InterpretCodeOptions) (output *tree, err error) {
	rootNode := &parentTreeNode{}
	parentNodeStack := []*parentTreeNode{
		rootNode,
	}

	rootNode.lexeme = startProgramLexeme
	rootNode.tree = output

	output = new(tree)
	output.rootNode = rootNode
	output.interpretCodeOptions = interpretCodeOptions

	for i := range interpretCodeOptions.saveStacks {
		if interpretCodeOptions.saveStacks[i] == nil {
			interpretCodeOptions.saveStacks[i] = make(memoryCellCollection, 0, interpretCodeOptionsSaveStackMaxLen)
		}
	}

	for _, t := range input {
		if err = output.addNode(t, &parentNodeStack); err != nil {
			return
		}
	}

	return
}
