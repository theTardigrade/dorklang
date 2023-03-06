package main

import (
	"fmt"
	"math"
)

type Tree struct {
	RootNode *ParentTreeNode
}

type TreeNode interface {
	Value(uint8) (uint8, error)
}

type ParentTreeNode struct {
	Lexeme     Lexeme
	ChildNodes []TreeNode
}

type TerminalTreeNode struct {
	Lexeme     Lexeme
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
	case IncrementLexexme:
		output++
	case DecrementLexeme:
		output -= 1
	case DoubleLexeme:
		output *= 2
	case HalfLexeme:
		output /= 2
	case SquareLexeme:
		output *= output
	case PrintCharacterLexeme:
		fmt.Printf("%c", input)
	case PrintNumberLexeme:
		fmt.Printf("%d", input)
	case MinLexeme:
		output = 0
	case MaxLexeme:
		output = math.MaxUint8
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

	for _, l := range input {
		switch l {
		case StartAdditionSectionLexeme,
			StartSubtractionSectionLexeme:
			{
				nextNode := &ParentTreeNode{
					Lexeme: l,
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
		case IncrementLexexme,
			DecrementLexeme,
			DoubleLexeme,
			HalfLexeme,
			SquareLexeme,
			MinLexeme,
			MaxLexeme,
			PrintCharacterLexeme,
			PrintNumberLexeme:
			{
				nextNode := &TerminalTreeNode{
					Lexeme: l,
				}

				if len(parentNodeStack) == 0 {
					err = ErrTreeParentNodeUnfound
					return
				}

				nextParentNode := parentNodeStack[len(parentNodeStack)-1]
				nextParentNode.ChildNodes = append(nextParentNode.ChildNodes, nextNode)
				nextNode.ParentNode = nextParentNode
			}
		}
	}

	return
}
