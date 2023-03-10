package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"os"
	"sort"
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

func (node defaultTreeNode) getTree() *tree {
	return node.tree
}

func (node defaultTreeNode) getData() []byte {
	return node.data
}

type parentTreeNode struct {
	defaultTreeNode
	childNodes []treeNode
}

type terminalTreeNode struct {
	defaultTreeNode
	parentNode *parentTreeNode
}

func (tree *tree) Run() (err error) {
	_, err = tree.rootNode.value(0)

	return
}

func (tree *tree) saveStackPtr() (stackPtr *memoryCellCollection, err error) {
	if tree.saveStackIndex >= len(tree.saveStacks) {
		err = ErrTreeSaveStackIndexInvalid
		return
	}

	stackPtr = &tree.saveStacks[tree.saveStackIndex]

	return
}

func (tree *tree) saveStack() (stack memoryCellCollection, err error) {
	stackPtr, err := tree.saveStackPtr()
	if err != nil {
		return
	}

	stack = *stackPtr

	return
}

func (node *parentTreeNode) value(input memoryCell) (output memoryCell, err error) {
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
		startSubtractionSectionLexeme,
		startMultiplicationSectionLexeme,
		startDivisionSectionLexeme:
		{
			var localOutput memoryCell

			for _, node := range node.childNodes {
				localOutput, err = node.value(localOutput)
				if err != nil {
					return
				}
			}

			switch node.lexeme {
			case startAdditionSectionLexeme:
				output += localOutput
			case startSubtractionSectionLexeme:
				output -= localOutput
			case startMultiplicationSectionLexeme:
				output *= localOutput
			case startDivisionSectionLexeme:
				output /= localOutput
			default:
				err = ErrLexemeUnrecognized
				return
			}
		}
	case startReadFileSectionLexeme:
		{
			for _, node := range node.childNodes {
				data := node.getData()

				if len(data) == 0 {
					continue
				}

				tree := node.getTree()

				if tree == nil {
					err = ErrTreeUnfound
					return
				}

				fileName := string(data)

				var content []byte
				content, err = os.ReadFile(fileName)
				if err != nil {
					return
				}

				var saveStackPtr *memoryCellCollection
				saveStackPtr, err = tree.saveStackPtr()
				if err != nil {
					return
				}
				saveStack := *saveStackPtr

				contentRunes := []rune(string(content))

				for i := len(contentRunes) - 1; i >= 0; i-- {
					if len(saveStack) == treeSaveStackMaxLen {
						err = ErrTreeSaveStackFull
						return
					}

					saveStack = append(saveStack, memoryCell(contentRunes[i]))
				}

				*saveStackPtr = saveStack
			}
		}
	case startCommentSectionLexeme:
	default:
		err = ErrLexemeUnrecognized
	}

	return
}

func (node *terminalTreeNode) value(input memoryCell) (output memoryCell, err error) {
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

			var saveStackPtr *memoryCellCollection
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

			var saveStackPtr *memoryCellCollection
			saveStackPtr, err = node.tree.saveStackPtr()
			if err != nil {
				return
			}
			saveStack := *saveStackPtr

			if len(saveStack) == 0 {
				err = ErrTreeSaveStackEmpty
				return
			}

			var sum memoryCell

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

			var saveStackPtr *memoryCellCollection
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

			var saveStackPtr *memoryCellCollection
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

			var saveStackPtr *memoryCellCollection
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

			var saveStackPtr *memoryCellCollection
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

			var saveStackPtr *memoryCellCollection
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

			var saveStackPtr *memoryCellCollection
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
	case logicalAndStackPairLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStack memoryCellCollection
			saveStack, err = node.tree.saveStack()
			if err != nil {
				return
			}

			if len(saveStack) < 2 {
				err = ErrTreeSaveStackEmpty
				return
			}

			value1 := saveStack[len(saveStack)-1]
			value2 := saveStack[len(saveStack)-2]

			result := value1 > 0 && value2 > 0

			if result {
				output = 1
			} else {
				output = 0
			}
		}
	case logicalAndStackWholeLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStack memoryCellCollection
			saveStack, err = node.tree.saveStack()
			if err != nil {
				return
			}

			if len(saveStack) == 0 {
				err = ErrTreeSaveStackEmpty
				return
			}

			result := saveStack[len(saveStack)-1] > 0

			for i := len(saveStack) - 2; i >= 0; i-- {
				if !result {
					break
				}

				result = saveStack[i] > 0
			}

			if result {
				output = 1
			} else {
				output = 0
			}
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

			output = memoryCell(b[0])
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

			output = memoryCell(b.Uint64())
		}
	case setSecondTimestampLexeme:
		output = memoryCell(time.Now().Unix())
	case setNanosecondTimestampLexeme:
		output = memoryCell(time.Now().UnixNano())
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
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStackPtr *memoryCellCollection
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
		}
	case countStackLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStack memoryCellCollection
			saveStack, err = node.tree.saveStack()
			if err != nil {
				return
			}

			output = memoryCell(len(saveStack))
		}
	case popStackLastLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStackPtr *memoryCellCollection
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
		}
	case popStackRandomLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStackPtr *memoryCellCollection
			saveStackPtr, err = node.tree.saveStackPtr()
			if err != nil {
				return
			}
			saveStack := *saveStackPtr

			if len(saveStack) == 0 {
				err = ErrTreeSaveStackEmpty
				return
			}

			b := big.NewInt(int64(len(saveStack)))
			b, err = rand.Int(rand.Reader, b)
			if err != nil {
				return
			}

			index := b.Uint64()

			output = saveStack[index]
			*saveStackPtr = append(saveStack[:index], saveStack[index+1:]...)
		}
	case writeStackToFileLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStack memoryCellCollection
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
			fileName := strconv.FormatUint(uint64(output), 10) + treeSaveFileExtension

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

			fileName := strconv.FormatUint(uint64(output), 10) + treeSaveFileExtension
			content, err = os.ReadFile(fileName)
			if err != nil {
				return
			}

			var saveStackPtr *memoryCellCollection
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

				saveStack = append(saveStack, memoryCell(value))
			}

			*saveStackPtr = saveStack
		}
	case hashStackOneByteLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStack memoryCellCollection
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

			output = memoryCell(hash.Uint8(content))
		}
	case hashStackEightByteLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStack memoryCellCollection
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

			output = memoryCell(hash.Uint64(content))
		}
	case sortStackLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStack memoryCellCollection
			saveStack, err = node.tree.saveStack()
			if err != nil {
				return
			}

			sort.Sort(saveStack)
		}
	case shuffleStackLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStack memoryCellCollection
			saveStack, err = node.tree.saveStack()
			if err != nil {
				return
			}

			indicesToSwap := make([]uint64, len(saveStack))

			var b, b2 *big.Int

			b = big.NewInt(int64(len(saveStack)))

			for i := 0; i < len(indicesToSwap); i++ {
				b2, err = rand.Int(rand.Reader, b)
				if err != nil {
					return
				}

				indicesToSwap[i] = b2.Uint64()
			}

			for i, j := range indicesToSwap {
				saveStack[i], saveStack[j] = saveStack[j], saveStack[i]
			}
		}
	case swapStackTopLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStack memoryCellCollection
			saveStack, err = node.tree.saveStack()
			if err != nil {
				return
			}

			saveStackLen := len(saveStack)

			if saveStackLen < 2 {
				err = ErrTreeSaveStackEmpty
				return
			}

			saveStack[saveStackLen-1], saveStack[saveStackLen-2] = saveStack[saveStackLen-2], saveStack[saveStackLen-1]
		}
	case iotaFromZeroLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStackPtr *memoryCellCollection
			saveStackPtr, err = node.tree.saveStackPtr()
			if err != nil {
				return
			}
			saveStack := *saveStackPtr

			for i := memoryCell(0); i < output; i++ {
				if len(saveStack) == treeSaveStackMaxLen {
					err = ErrTreeSaveStackFull
					return
				}

				saveStack = append(saveStack, i)
			}

			*saveStackPtr = saveStack
		}
	case iotaFromOneLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			var saveStackPtr *memoryCellCollection
			saveStackPtr, err = node.tree.saveStackPtr()
			if err != nil {
				return
			}
			saveStack := *saveStackPtr

			for i := memoryCell(1); i < output; i++ {
				if len(saveStack) == treeSaveStackMaxLen {
					err = ErrTreeSaveStackFull
					return
				}

				saveStack = append(saveStack, i)
			}

			*saveStackPtr = saveStack
		}
	case invertLexeme:
		{
			if output == 0 {
				output = 1
			} else {
				output = 0
			}
		}
	case deleteFileLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			fileName := strconv.FormatUint(uint64(output), 10) + treeSaveFileExtension

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

			var saveStackPtr *memoryCellCollection
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
	case plaintextLexeme:
	default:
		err = ErrLexemeUnrecognized
	}

	return
}

func produceTree(input []token) (output *tree, err error) {
	rootNode := &parentTreeNode{}
	parentNodeStack := []*parentTreeNode{
		rootNode,
	}

	rootNode.lexeme = startAdditionSectionLexeme
	rootNode.tree = output

	output = new(tree)
	output.rootNode = rootNode

	for i := range output.saveStacks {
		output.saveStacks[i] = make(memoryCellCollection, 0, treeSaveStackMaxLen)
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
			saveStackUseIndexZeroLexeme,
			saveStackUseIndexOneLexeme,
			hashStackOneByteLexeme,
			hashStackEightByteLexeme,
			sortStackLexeme,
			shuffleStackLexeme,
			swapStackTopLexeme,
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
		case separatorLexeme:
		default:
			err = ErrLexemeUnrecognized
			return
		}
	}

	return
}
