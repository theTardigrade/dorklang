package main

import (
	"fmt"
)

type Tree struct {
	RootNode   *ParentTreeNode
	savedValue uint8
}

type TreeNode interface {
	Value(uint8) (uint8, error)
}

type ParentTreeNode struct {
	Lexeme     Lexeme
	Tree       *Tree
	ChildNodes []TreeNode
}

type TerminalTreeNode struct {
	Lexeme     Lexeme
	Tree       *Tree
	ParentNode *ParentTreeNode
}

func (tree *Tree) Run() (err error) {
	_, err = tree.RootNode.Value(0)

	return
}

func (node *ParentTreeNode) Value(input uint8) (output uint8, err error) {
	output = input

	switch node.Lexeme {
	case StartAdditionSectionLexeme,
		StartSubtractionSectionLexeme:
		{
			var localOutput uint8

			for _, node := range node.ChildNodes {
				localOutput, err = node.Value(localOutput)
				if err != nil {
					return
				}
			}

			if node.Lexeme == StartAdditionSectionLexeme {
				output += localOutput
			} else {
				output -= localOutput
			}
		}
	}

	return
}

func (node *TerminalTreeNode) Value(input uint8) (output uint8, err error) {
	output = input

	switch node.Lexeme {
	case IncrementOneLexeme:
		output++
	case IncrementEightLexeme:
		output += 8
	case DecrementOneLexeme:
		output--
	case DecrementEightLexeme:
		output -= 8
	case MultiplyTwoLexeme:
		output *= 2
	case MultiplyEightLexeme:
		output *= 8
	case DivideTwoLexeme:
		output /= 2
	case DivideEightLexeme:
		output /= 8
	case SquareLexeme:
		output *= output
	case CubeLexeme:
		output *= output * output
	case PrintCharacterLexeme:
		fmt.Printf("%c", input)
	case PrintNumberLexeme:
		fmt.Printf("%d", input)
	case MinimumLexeme:
		output = 0
	case MiddleLexeme:
		output = 128
	case MaximumLexeme:
		output = 255
	case SaveLexeme:
		if node.Tree == nil {
			err = ErrTreeUnfound
			return
		}

		node.Tree.savedValue = output
	case LoadLexeme:
		if node.Tree == nil {
			err = ErrTreeUnfound
			return
		}

		output = node.Tree.savedValue
	default:
		err = ErrLexemeUnrecognized
	}

	return
}

func produceTree(input []Lexeme) (output *Tree, err error) {
	rootNode := &ParentTreeNode{
		Lexeme: StartAdditionSectionLexeme,
	}
	parentNodeStack := []*ParentTreeNode{
		rootNode,
	}

	output = new(Tree)
	output.RootNode = rootNode
	rootNode.Tree = output

	for _, l := range input {
		switch l {
		case StartAdditionSectionLexeme,
			StartSubtractionSectionLexeme:
			{
				nextNode := &ParentTreeNode{
					Lexeme: l,
					Tree:   output,
				}

				if len(parentNodeStack) == 0 {
					err = ErrTreeParentNodeUnfound
					return
				}

				nextParentNode := parentNodeStack[len(parentNodeStack)-1]
				nextParentNode.ChildNodes = append(nextParentNode.ChildNodes, nextNode)

				parentNodeStack = append(parentNodeStack, nextNode)
			}
		case EndAdditionSectionLexeme,
			EndSubtractionSectionLexeme:
			{
				if len(parentNodeStack) == 0 {
					err = ErrTreeParentNodeUnfound
					return
				}

				parentNodeStack = parentNodeStack[:len(parentNodeStack)-1]
			}
		case IncrementOneLexeme,
			IncrementEightLexeme,
			DecrementOneLexeme,
			DecrementEightLexeme,
			MultiplyTwoLexeme,
			MultiplyEightLexeme,
			DivideTwoLexeme,
			DivideEightLexeme,
			SquareLexeme,
			CubeLexeme,
			MinimumLexeme,
			MiddleLexeme,
			MaximumLexeme,
			PrintCharacterLexeme,
			PrintNumberLexeme,
			SaveLexeme,
			LoadLexeme:
			{
				nextNode := &TerminalTreeNode{
					Lexeme: l,
					Tree:   output,
				}

				if len(parentNodeStack) == 0 {
					err = ErrTreeParentNodeUnfound
					return
				}

				nextParentNode := parentNodeStack[len(parentNodeStack)-1]
				nextParentNode.ChildNodes = append(nextParentNode.ChildNodes, nextNode)
				nextNode.ParentNode = nextParentNode
			}
		case SeparatorLexeme:
		default:
			err = ErrLexemeUnrecognized
			return
		}
	}

	return
}
