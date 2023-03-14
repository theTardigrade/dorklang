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
		defaultNode := defaultTreeNode{
			lexeme: t.lex,
			data:   t.data,
			tree:   output,
		}

		switch t.lex {
		case startAdditionSectionLexeme,
			startSubtractionSectionLexeme,
			startMultiplicationSectionLexeme,
			startDivisionSectionLexeme,
			startJumpIfPositiveSectionLexeme,
			startJumpIfZeroSectionLexeme,
			startReadFileSectionLexeme,
			startCommentSectionLexeme:
			{
				nextNode := &parentTreeNode{
					defaultTreeNode: defaultNode,
				}

				if len(parentNodeStack) == 0 {
					err = ErrTreeParentNodeUnfound
					return
				}

				nextParentNode := parentNodeStack[len(parentNodeStack)-1]
				nextParentNode.childNodes = append(nextParentNode.childNodes, nextNode)

				parentNodeStack = append(parentNodeStack, nextNode)
			}
		case endAdditionSectionLexeme,
			endSubtractionSectionLexeme,
			endMultiplicationSectionLexeme,
			endDivisionSectionLexeme,
			endJumpIfPositiveSectionLexeme,
			endJumpIfZeroSectionLexeme,
			endReadFileSectionLexeme,
			endCommentSectionLexeme:
			{
				if len(parentNodeStack) == 0 {
					err = ErrTreeParentNodeUnfound
					return
				}

				parentNode := parentNodeStack[len(parentNodeStack)-1]
				parentNode.data = t.data

				parentNodeStack = parentNodeStack[:len(parentNodeStack)-1]
			}
		case addOneLexeme,
			addEightLexeme,
			addStackPairLexeme,
			addStackWholeLexeme,
			subtractOneLexeme,
			subtractEightLexeme,
			subtractStackPairLexeme,
			subtractStackWholeLexeme,
			multiplyTwoLexeme,
			multiplyEightLexeme,
			multiplyStackPairLexeme,
			multiplyStackWholeLexeme,
			divideTwoLexeme,
			divideEightLexeme,
			divideStackPairLexeme,
			divideStackWholeLexeme,
			squareLexeme,
			cubeLexeme,
			setZeroLexeme,
			setOneByteLexeme,
			setEightByteLexeme,
			setOneKibibyteLexeme,
			setEightKibibyteLexeme,
			setOneMebibyteLexeme,
			setEightMebibyteLexeme,
			setOneGibibyteLexeme,
			setEightGibibyteLexeme,
			setRandomByteLexeme,
			setRandomMaxLexeme,
			setSecondTimestampLexeme,
			setNanosecondTimestampLexeme,
			printCharacterLexeme,
			printNumberLexeme,
			inputCharacterLexeme,
			inputNumberLexeme,
			logicalAndStackPairLexeme,
			logicalAndStackWholeLexeme,
			pushStackLexeme,
			countStackLexeme,
			popStackLastLexeme,
			popStackRandomLexeme,
			useStackIndexZeroLexeme,
			useStackIndexOneLexeme,
			useStackIndexSwappedLexeme,
			hashStackOneByteLexeme,
			hashStackEightByteLexeme,
			sortStackAscendingLexeme,
			sortStackDescendingLexeme,
			shuffleStackLexeme,
			swapStackTopLexeme,
			reverseStackLexeme,
			invertLexeme,
			iotaFromZeroLexeme,
			iotaFromOneLexeme,
			plaintextLexeme,
			writeStackToFileLexeme,
			readStackFromFileLexeme,
			deleteFileLexeme,
			clearStackLexeme,
			resetStateLexeme:
			{
				nextNode := &terminalTreeNode{
					defaultTreeNode: defaultNode,
				}

				if len(parentNodeStack) == 0 {
					err = ErrTreeParentNodeUnfound
					return
				}

				nextParentNode := parentNodeStack[len(parentNodeStack)-1]
				nextParentNode.childNodes = append(nextParentNode.childNodes, nextNode)
				nextNode.parentNode = nextParentNode
			}
		case startProgramLexeme,
			endProgramLexeme,
			separatorLexeme,
			emptyLexeme:
		default:
			err = ErrLexemeUnrecognized
			return
		}
	}

	return
}
