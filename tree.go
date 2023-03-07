package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
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
	rootNode       *parentTreeNode
	saveStackIndex uint64
	saveStacks     [2][]uint64
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

func (tree *tree) saveStackPtr() (stackPtr *[]uint64, err error) {
	if tree.saveStackIndex >= uint64(len(tree.saveStacks)) {
		err = ErrTreeSaveStackIndexInvalid
		return
	}

	stackPtr = &tree.saveStacks[tree.saveStackIndex]

	return
}

func (tree *tree) saveStack() (stack []uint64, err error) {
	if tree.saveStackIndex >= uint64(len(tree.saveStacks)) {
		err = ErrTreeSaveStackIndexInvalid
		return
	}

	stack = tree.saveStacks[tree.saveStackIndex]

	return
}

func (node *parentTreeNode) value(input uint64) (output uint64, err error) {
	output = input

	switch node.lexeme {
	case startJumpIfPositiveSectionLexeme:
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
	case startJumpIfZeroSectionLexeme:
		{
			for output == 0 {
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

			var saveStackPtr *[]uint64
			saveStackPtr, err = node.tree.saveStackPtr()
			if err != nil {
				return
			}
			saveStack := *saveStackPtr

			if len(saveStack) < 2 {
				err = ErrTreeSaveStackEmpty
				return
			}

			augend := saveStack[len(saveStack)-1]
			addend := saveStack[len(saveStack)-2]

			*saveStackPtr = saveStack[:len(saveStack)-2]

			output = augend + addend
		}
	case addStackWholeLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStackPtr *[]uint64
			saveStackPtr, err = node.tree.saveStackPtr()
			if err != nil {
				return
			}
			saveStack := *saveStackPtr

			if len(saveStack) == 0 {
				err = ErrTreeSaveStackEmpty
				return
			}

			var sum uint64

			for i := len(saveStack) - 1; i >= 0; i-- {
				sum += saveStack[i]
			}

			*saveStackPtr = saveStack[:0]

			output = sum
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

			var saveStackPtr *[]uint64
			saveStackPtr, err = node.tree.saveStackPtr()
			if err != nil {
				return
			}
			saveStack := *saveStackPtr

			if len(saveStack) < 2 {
				err = ErrTreeSaveStackEmpty
				return
			}

			minuend := saveStack[len(saveStack)-1]
			subtrahend := saveStack[len(saveStack)-2]

			*saveStackPtr = saveStack[:len(saveStack)-2]

			output = minuend - subtrahend
		}
	case subtractStackWholeLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStackPtr *[]uint64
			saveStackPtr, err = node.tree.saveStackPtr()
			if err != nil {
				return
			}
			saveStack := *saveStackPtr

			if len(saveStack) == 0 {
				err = ErrTreeSaveStackEmpty
				return
			}

			subtraction := saveStack[len(saveStack)-1]

			for i := len(saveStack) - 2; i >= 0; i-- {
				subtraction -= saveStack[i]
			}

			*saveStackPtr = saveStack[:0]

			output = subtraction
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

			var saveStackPtr *[]uint64
			saveStackPtr, err = node.tree.saveStackPtr()
			if err != nil {
				return
			}
			saveStack := *saveStackPtr

			if len(saveStack) < 2 {
				err = ErrTreeSaveStackEmpty
				return
			}

			multiplier := saveStack[len(saveStack)-1]
			multiplicand := saveStack[len(saveStack)-2]

			*saveStackPtr = saveStack[:len(saveStack)-2]

			output = multiplier * multiplicand
		}
	case multiplyStackWholeLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStackPtr *[]uint64
			saveStackPtr, err = node.tree.saveStackPtr()
			if err != nil {
				return
			}
			saveStack := *saveStackPtr

			if len(saveStack) == 0 {
				err = ErrTreeSaveStackEmpty
				return
			}

			product := saveStack[len(saveStack)-1]

			for i := len(saveStack) - 2; i >= 0; i-- {
				product *= saveStack[i]
			}

			*saveStackPtr = saveStack[:0]

			output = product
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

			var saveStackPtr *[]uint64
			saveStackPtr, err = node.tree.saveStackPtr()
			if err != nil {
				return
			}
			saveStack := *saveStackPtr

			if len(saveStack) < 2 {
				err = ErrTreeSaveStackEmpty
				return
			}

			dividend := saveStack[len(saveStack)-1]
			divisor := saveStack[len(saveStack)-2]

			*saveStackPtr = saveStack[:len(saveStack)-2]

			output = dividend / divisor
		}
	case divideStackWholeLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStackPtr *[]uint64
			saveStackPtr, err = node.tree.saveStackPtr()
			if err != nil {
				return
			}
			saveStack := *saveStackPtr

			if len(saveStack) == 0 {
				err = ErrTreeSaveStackEmpty
				return
			}

			division := saveStack[len(saveStack)-1]

			for i := len(saveStack) - 2; i >= 0; i-- {
				division %= saveStack[i]
			}

			*saveStackPtr = saveStack[:0]

			output = division
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
	case setRandomByteLexeme:
		{
			b := make([]byte, 1)

			if _, err = rand.Reader.Read(b); err != nil {
				return
			}

			output = uint64(b[0])
		}
	case setRandomMaxLexeme:
		{
			var b, b2 *big.Int

			b = big.NewInt(1)

			b2 = new(big.Int)
			b2.SetUint64(math.MaxUint64)
			b2.Add(b2, b)

			b, err = rand.Int(rand.Reader, b2)
			if err != nil {
				return
			}

			output = b.Uint64()
		}
	case setSecondTimestampLexeme:
		output = uint64(time.Now().Unix())
	case setNanosecondTimestampLexeme:
		output = uint64(time.Now().UnixNano())
	case saveStackUseIndexZeroLexeme:
		if node.tree == nil {
			err = ErrTreeUnfound
			return
		}

		node.tree.saveStackIndex = 0
	case saveStackUseIndexOneLexeme:
		if node.tree == nil {
			err = ErrTreeUnfound
			return
		}

		node.tree.saveStackIndex = 1
	case pushStackLexeme:
		if node.tree == nil {
			err = ErrTreeUnfound
			return
		}

		var saveStackPtr *[]uint64
		saveStackPtr, err = node.tree.saveStackPtr()
		if err != nil {
			return
		}
		saveStack := *saveStackPtr

		if len(saveStack) == treeSaveStackMaxLen {
			err = ErrTreeSaveStackFull
			return
		}

		*saveStackPtr = append(saveStack, output)
	case popStackLexeme:
		if node.tree == nil {
			err = ErrTreeUnfound
			return
		}

		var saveStackPtr *[]uint64
		saveStackPtr, err = node.tree.saveStackPtr()
		if err != nil {
			return
		}
		saveStack := *saveStackPtr

		if len(saveStack) == 0 {
			err = ErrTreeSaveStackEmpty
			return
		}

		output = saveStack[len(saveStack)-1]
		*saveStackPtr = saveStack[:len(saveStack)-1]
	case writeStackToFileLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStack []uint64
			saveStack, err = node.tree.saveStack()
			if err != nil {
				return
			}

			var contentBuilder bytes.Buffer

			for _, value := range saveStack {
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

			var saveStackPtr *[]uint64
			saveStackPtr, err = node.tree.saveStackPtr()
			if err != nil {
				return
			}
			saveStack := *saveStackPtr

			saveStack = saveStack[:0]

			for _, value := range string(content) {
				if len(saveStack) == treeSaveStackMaxLen {
					err = ErrTreeSaveStackFull
					return
				}

				saveStack = append(saveStack, uint64(value))
			}

			*saveStackPtr = saveStack
		}
	case hashStackOneByteLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStack []uint64
			saveStack, err = node.tree.saveStack()
			if err != nil {
				return
			}

			var contentBuilder bytes.Buffer

			for _, value := range saveStack {
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

			var saveStack []uint64
			saveStack, err = node.tree.saveStack()
			if err != nil {
				return
			}

			var contentBuilder bytes.Buffer

			for _, value := range saveStack {
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
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStackPtr *[]uint64
			saveStackPtr, err = node.tree.saveStackPtr()
			if err != nil {
				return
			}
			saveStack := *saveStackPtr

			*saveStackPtr = saveStack[:0]
		}
	case resetStateLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			for i := len(node.tree.saveStacks) - 1; i >= 0; i-- {
				node.tree.saveStacks[i] = node.tree.saveStacks[i][:0]
			}

			output = 0
		}
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

	for i := range output.saveStacks {
		output.saveStacks[i] = make([]uint64, 0, treeSaveStackMaxLen)
	}

	rootNode.tree = output

	for _, l := range input {
		switch l {
		case startAdditionSectionLexeme,
			startSubtractionSectionLexeme,
			startJumpIfPositiveSectionLexeme,
			startJumpIfZeroSectionLexeme,
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
			endJumpIfPositiveSectionLexeme,
			endJumpIfZeroSectionLexeme,
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
			pushStackLexeme,
			popStackLexeme,
			saveStackUseIndexZeroLexeme,
			saveStackUseIndexOneLexeme,
			hashStackOneByteLexeme,
			hashStackEightByteLexeme,
			writeStackToFileLexeme,
			readStackFromFileLexeme,
			deleteFileLexeme,
			clearStackLexeme,
			resetStateLexeme:
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
