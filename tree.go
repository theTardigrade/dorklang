package main

import (
	"fmt"
)

type tree struct {
	rootNode   *parentTreeNode
	savedValue uint8
}

type treeNode interface {
	value(uint8) (uint8, error)
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

func (node *parentTreeNode) value(input uint8) (output uint8, err error) {
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
			var localOutput uint8

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

func (node *terminalTreeNode) value(input uint8) (output uint8, err error) {
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
		fmt.Printf("%c", input)
	case printNumberLexeme:
		fmt.Printf("%d", input)
	case minimumLexeme:
		output = 0
	case middleLexeme:
		output = 128
	case maximumLexeme:
		output = 255
	case saveLexeme:
		if node.tree == nil {
			err = ErrTreeUnfound
			return
		}

		node.tree.savedValue = output
	case loadLexeme:
		if node.tree == nil {
			err = ErrTreeUnfound
			return
		}

		output = node.tree.savedValue
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
			minimumLexeme,
			middleLexeme,
			maximumLexeme,
			printCharacterLexeme,
			printNumberLexeme,
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
