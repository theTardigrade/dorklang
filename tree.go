package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"

	hash "github.com/theTardigrade/golang-hash"
)

const (
	treeSaveFileExtension = ".dork.txt"
	treeSaveStackMaxLen   = 1 << 20 // 1_048_576
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
	case addOneLexeme:
		output++
	case addEightLexeme:
		output += 8
	case addStackPairLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			if len(node.tree.saveStack) < 2 {
				err = ErrTreeSaveStackEmpty
				return
			}

			augend := node.tree.saveStack[len(node.tree.saveStack)-1]
			addend := node.tree.saveStack[len(node.tree.saveStack)-2]

			node.tree.saveStack = node.tree.saveStack[:len(node.tree.saveStack)-2]

			output = augend + addend
		}
	case subtractOneLexeme:
		output--
	case subtractEightLexeme:
		output -= 8
	case subtractStackPairLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			if len(node.tree.saveStack) < 2 {
				err = ErrTreeSaveStackEmpty
				return
			}

			minuend := node.tree.saveStack[len(node.tree.saveStack)-1]
			subtrahend := node.tree.saveStack[len(node.tree.saveStack)-2]

			node.tree.saveStack = node.tree.saveStack[:len(node.tree.saveStack)-2]

			output = minuend - subtrahend
		}
	case multiplyTwoLexeme:
		output *= 2
	case multiplyEightLexeme:
		output *= 8
	case multiplyStackPairLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			if len(node.tree.saveStack) < 2 {
				err = ErrTreeSaveStackEmpty
				return
			}

			multiplier := node.tree.saveStack[len(node.tree.saveStack)-1]
			multiplicand := node.tree.saveStack[len(node.tree.saveStack)-2]

			node.tree.saveStack = node.tree.saveStack[:len(node.tree.saveStack)-2]

			output = multiplier * multiplicand
		}
	case divideTwoLexeme:
		output /= 2
	case divideEightLexeme:
		output /= 8
	case divideStackPairLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			if len(node.tree.saveStack) < 2 {
				err = ErrTreeSaveStackEmpty
				return
			}

			dividend := node.tree.saveStack[len(node.tree.saveStack)-1]
			divisor := node.tree.saveStack[len(node.tree.saveStack)-2]

			node.tree.saveStack = node.tree.saveStack[:len(node.tree.saveStack)-2]

			output = dividend / divisor
		}
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
	case setOneByteLexeme:
		output = 1 << 3 // 8
	case setEightByteLexeme:
		output = 1 << 6 // 64
	case setOneKibibyteLexeme:
		output = 1 << 13 // 8_192
	case setEightKibibyteLexeme:
		output = 1 << 16 // 65_536
	case setOneMebibyteLexeme:
		output = 1 << 23 // 8_388_608
	case setEightMebibyteLexeme:
		output = 1 << 26 // 67_108_864
	case setOneGibibyteLexeme:
		output = 1 << 33 // 8_589_934_592
	case setEightGibibyteLexeme:
		output = 1 << 36 // 68_719_476_736
	case setSecondTimestampLexeme:
		output = uint64(time.Now().Unix())
	case setNanosecondTimestampLexeme:
		output = uint64(time.Now().UnixNano())
	case saveToStackLexeme:
		if node.tree == nil {
			err = ErrTreeUnfound
			return
		}

		if len(node.tree.saveStack) == treeSaveStackMaxLen {
			err = ErrTreeSaveStackFull
			return
		}

		node.tree.saveStack = append(node.tree.saveStack, output)
	case loadFromStackLexeme:
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
	case writeStackToFileLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var contentBuilder bytes.Buffer

			for _, value := range node.tree.saveStack {
				if _, err = contentBuilder.WriteRune(rune(value)); err != nil {
					return
				}
			}

			content := contentBuilder.Bytes()
			fileName := strconv.FormatUint(output, 10) + treeSaveFileExtension

			if err = os.WriteFile(fileName, content, os.ModePerm); err != nil {
				return
			}
		}
	case readStackFromFileLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var content []byte

			fileName := strconv.FormatUint(output, 10) + treeSaveFileExtension
			content, err = os.ReadFile(fileName)
			if err != nil {
				return
			}

			node.tree.saveStack = node.tree.saveStack[:0]

			for _, value := range string(content) {
				node.tree.saveStack = append(node.tree.saveStack, uint64(value))
			}
		}
	case hashStackOneByteLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var contentBuilder bytes.Buffer

			for _, value := range node.tree.saveStack {
				if _, err = contentBuilder.WriteRune(rune(value)); err != nil {
					return
				}
			}

			content := contentBuilder.Bytes()

			output = uint64(hash.Uint8(content))
		}
	case hashStackEightByteLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var contentBuilder bytes.Buffer

			for _, value := range node.tree.saveStack {
				if _, err = contentBuilder.WriteRune(rune(value)); err != nil {
					return
				}
			}

			content := contentBuilder.Bytes()

			output = hash.Uint64(content)
		}
	case deleteFileLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			fileName := strconv.FormatUint(output, 10) + treeSaveFileExtension

			if err = os.Remove(fileName); err != nil {
				return
			}
		}
	case clearStackLexeme:
		node.tree.saveStack = node.tree.saveStack[:0]
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
		case addOneLexeme,
			addEightLexeme,
			addStackPairLexeme,
			subtractOneLexeme,
			subtractEightLexeme,
			subtractStackPairLexeme,
			multiplyTwoLexeme,
			multiplyEightLexeme,
			multiplyStackPairLexeme,
			divideTwoLexeme,
			divideEightLexeme,
			divideStackPairLexeme,
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
			setSecondTimestampLexeme,
			setNanosecondTimestampLexeme,
			printCharacterLexeme,
			printNumberLexeme,
			inputCharacterLexeme,
			inputNumberLexeme,
			saveToStackLexeme,
			loadFromStackLexeme,
			hashStackOneByteLexeme,
			hashStackEightByteLexeme,
			writeStackToFileLexeme,
			readStackFromFileLexeme,
			deleteFileLexeme,
			clearStackLexeme:
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
