package main

import (
	"fmt"
	"time"
)

const (
	treeSaveStackMaxLen = 256 * 256 * 256 // 16_777_216
)

type tree struct {
	rootNode  *parentTreeNode
	saveStack []uint64
}

type treeNode interface {
	value(uint64) (uint64, error)
}

type parentTreeNode struct {
	lexeme     lexeme
	tree       *tree
	childNodes []treeNode
}

type terminalTreeNode struct {
	lexeme     lexeme
	tree       *tree
	parentNode *parentTreeNode
}

func (tree *tree) Run() (err error) {
	_, err = tree.rootNode.value(0)

	return
}

func (node *parentTreeNode) value(input uint64) (output uint64, err error) {
	output = input

	switch node.lexeme {
	case startJumpSectionLexeme:
		{
			for output > 0 {
				for _, node := range node.childNodes {
					output, err = node.value(output)
					if err != nil {
						return
					}
				}
			}
		}
	case startAdditionSectionLexeme,
		startSubtractionSectionLexeme:
		{
			var localOutput uint64

			for _, node := range node.childNodes {
				localOutput, err = node.value(localOutput)
				if err != nil {
					return
				}
			}

			if node.lexeme == startAdditionSectionLexeme {
				output += localOutput
			} else {
				output -= localOutput
			}
		}
	case startCommentSectionLexeme:
	default:
		err = ErrLexemeUnrecognized
	}

	return
}

func (node *terminalTreeNode) value(input uint64) (output uint64, err error) {
	output = input

	switch node.lexeme {
	case incrementOneLexeme:
		output++
	case incrementEightLexeme:
		output += 8
	case decrementOneLexeme:
		output--
	case decrementEightLexeme:
		output -= 8
	case multiplyTwoLexeme:
		output *= 2
	case multiplyEightLexeme:
		output *= 8
	case divideTwoLexeme:
		output /= 2
	case divideEightLexeme:
		output /= 8
	case squareLexeme:
		output *= output
	case cubeLexeme:
		output *= output * output
	case printCharacterLexeme:
		if _, err = fmt.Printf("%c", output); err != nil {
			return
		}
	case printNumberLexeme:
		if _, err = fmt.Printf("%d", output); err != nil {
			return
		}
	case inputCharacterLexeme:
		if _, err = fmt.Scanf("%c", &output); err != nil {
			return
		}
	case inputNumberLexeme:
		if _, err = fmt.Scanf("%d", &output); err != nil {
			return
		}
	case setZeroLexeme:
		output = 0
	case setOneNybbleLexeme:
		output = 16
	case setEightNybbleLexeme:
		output = 16 * 8 // 128
	case setOneByteLexeme:
		output = 256
	case setEightByteLexeme:
		output = 256 * 8 // 2_048
	case setSecondTimestampLexeme:
		output = uint64(time.Now().Unix())
	case setNanosecondTimestampLexeme:
		output = uint64(time.Now().UnixNano())
	case saveLexeme:
		if node.tree == nil {
			err = ErrTreeUnfound
			return
		}

		if len(node.tree.saveStack) == treeSaveStackMaxLen {
			err = ErrTreeSaveStackFull
			return
		}

		node.tree.saveStack = append(node.tree.saveStack, output)
	case loadLexeme:
		if node.tree == nil {
			err = ErrTreeUnfound
			return
		}

		if len(node.tree.saveStack) == 0 {
			err = ErrTreeSaveStackEmpty
			return
		}

		output = node.tree.saveStack[len(node.tree.saveStack)-1]
		node.tree.saveStack = node.tree.saveStack[:len(node.tree.saveStack)-1]
	default:
		err = ErrLexemeUnrecognized
	}

	return
}

func produceTree(input []lexeme) (output *tree, err error) {
	rootNode := &parentTreeNode{
		lexeme: startAdditionSectionLexeme,
	}
	parentNodeStack := []*parentTreeNode{
		rootNode,
	}

	output = new(tree)
	output.rootNode = rootNode
	output.saveStack = make([]uint64, 0, treeSaveStackMaxLen)

	rootNode.tree = output

	for _, l := range input {
		switch l {
		case startAdditionSectionLexeme,
			startSubtractionSectionLexeme,
			startJumpSectionLexeme,
			startCommentSectionLexeme:
			{
				nextNode := &parentTreeNode{
					lexeme: l,
					tree:   output,
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
			endJumpSectionLexeme,
			endCommentSectionLexeme:
			{
				if len(parentNodeStack) == 0 {
					err = ErrTreeParentNodeUnfound
					return
				}

				parentNodeStack = parentNodeStack[:len(parentNodeStack)-1]
			}
		case incrementOneLexeme,
			incrementEightLexeme,
			decrementOneLexeme,
			decrementEightLexeme,
			multiplyTwoLexeme,
			multiplyEightLexeme,
			divideTwoLexeme,
			divideEightLexeme,
			squareLexeme,
			cubeLexeme,
			setZeroLexeme,
			setOneNybbleLexeme,
			setEightNybbleLexeme,
			setOneByteLexeme,
			setEightByteLexeme,
			setSecondTimestampLexeme,
			setNanosecondTimestampLexeme,
			printCharacterLexeme,
			printNumberLexeme,
			inputCharacterLexeme,
			inputNumberLexeme,
			saveLexeme,
			loadLexeme:
			{
				nextNode := &terminalTreeNode{
					lexeme: l,
					tree:   output,
				}

				if len(parentNodeStack) == 0 {
					err = ErrTreeParentNodeUnfound
					return
				}

				nextParentNode := parentNodeStack[len(parentNodeStack)-1]
				nextParentNode.childNodes = append(nextParentNode.childNodes, nextNode)
				nextNode.parentNode = nextParentNode
			}
		case separatorLexeme:
		default:
			err = ErrLexemeUnrecognized
			return
		}
	}

	return
}
