package dorklang

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"time"

	hash "github.com/theTardigrade/golang-hash"
)

func (tr *tree) addTerminalNode(node *terminalTreeNode, parentNodeStack *[]*parentTreeNode) (err error) {
	if len(*parentNodeStack) == 0 {
		err = ErrTreeParentNodeUnfound
		return
	}

	nextParentNode := (*parentNodeStack)[len(*parentNodeStack)-1]
	nextParentNode.childNodes = append(nextParentNode.childNodes, node)

	node.parentNode = nextParentNode
	node.tree = tr

	return
}

func (tr *tree) addNode(t token, parentNodeStack *[]*parentTreeNode) (err error) {
	handled := true

	switch t.lex {
	case endAdditionSectionLexeme,
		endSubtractionSectionLexeme,
		endMultiplicationSectionLexeme,
		endDivisionSectionLexeme,
		endJumpIfPositiveSectionLexeme,
		endJumpIfZeroSectionLexeme,
		endReadFileSectionLexeme,
		endCommentSectionLexeme:
		{
			if len(*parentNodeStack) == 0 {
				err = ErrTreeParentNodeUnfound
				return
			}

			parentNode := (*parentNodeStack)[len(*parentNodeStack)-1]
			parentNode.data = t.data

			*parentNodeStack = (*parentNodeStack)[:len(*parentNodeStack)-1]
		}
	case parentLexeme:
		{
			dirs := bytes.Split(t.data, []byte{0})
			nextDir := dirs[0]
			initialDir := dirs[1]

			nextNode := &terminalTreeNode{
				defaultTreeNode: defaultTreeNode{
					lexeme: changeDirLexeme,
					data:   nextDir,
				},
			}

			if err = tr.addTerminalNode(nextNode, parentNodeStack); err != nil {
				return
			}

			for _, t2 := range t.childCollection {
				if err = tr.addNode(t2, parentNodeStack); err != nil {
					return
				}
			}

			nextNode = &terminalTreeNode{
				defaultTreeNode: defaultTreeNode{
					lexeme: changeDirLexeme,
					data:   initialDir,
				},
			}

			if err = tr.addTerminalNode(nextNode, parentNodeStack); err != nil {
				return
			}
		}
	case startProgramLexeme,
		endProgramLexeme,
		separatorLexeme,
		emptyLexeme:
	default:
		handled = false
	}

	if handled {
		return
	}

	defaultNode := defaultTreeNode{
		lexeme: t.lex,
		data:   t.data,
		tree:   tr,
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

			if len(*parentNodeStack) == 0 {
				err = ErrTreeParentNodeUnfound
				return
			}

			nextParentNode := (*parentNodeStack)[len(*parentNodeStack)-1]
			nextParentNode.childNodes = append(nextParentNode.childNodes, nextNode)

			*parentNodeStack = append(*parentNodeStack, nextNode)
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
		filePathLexeme,
		writeStackToFileLexeme,
		readStackFromFileLexeme,
		deleteFileLexeme,
		clearStackLexeme,
		resetStateLexeme:
		{
			nextNode := &terminalTreeNode{
				defaultTreeNode: defaultNode,
			}

			if err = tr.addTerminalNode(nextNode, parentNodeStack); err != nil {
				return
			}
		}
	default:
		err = ErrLexemeUnrecognized
		return
	}

	return
}

func (tree *tree) Run() (output memoryCell, err error) {
	initialDir, err := os.Getwd()
	if err != nil {
		return
	}

	err = os.Chdir(tree.interpretCodeOptions.WorkingDir)
	if err != nil {
		return
	}

	output, err = tree.rootNode.value(tree.interpretCodeOptions.initialCurrentValue)
	if err != nil {
		return
	}

	err = os.Chdir(initialDir)
	if err != nil {
		return
	}

	return
}

func (tree *tree) saveStackPtr() (stackPtr *memoryCellCollection, err error) {
	if tree.interpretCodeOptions.saveStackIndex >= len(tree.interpretCodeOptions.saveStacks) {
		err = ErrTreeSaveStackIndexInvalid
		return
	}

	stackPtr = &tree.interpretCodeOptions.saveStacks[tree.interpretCodeOptions.saveStackIndex]

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

func (node defaultTreeNode) getLexeme() lexeme {
	return node.lexeme
}

func (node defaultTreeNode) getTree() *tree {
	return node.tree
}

func (node defaultTreeNode) getData() []byte {
	return node.data
}

func (node *parentTreeNode) value(input memoryCell) (output memoryCell, err error) {
	output = input

	switch node.lexeme {
	case startJumpIfPositiveSectionLexeme:
		{
			for output > 0 {
				for _, node2 := range node.childNodes {
					output, err = node2.value(output)
					if err != nil {
						return
					}
				}
			}
		}
	case startJumpIfZeroSectionLexeme:
		{
			for output == 0 {
				for _, node2 := range node.childNodes {
					output, err = node2.value(output)
					if err != nil {
						return
					}
				}
			}
		}
	case startProgramLexeme,
		startReadFileSectionLexeme:
		{
			for _, node2 := range node.childNodes {
				output, err = node2.value(output)
				if err != nil {
					return
				}
			}
		}
	case startAdditionSectionLexeme,
		startSubtractionSectionLexeme,
		startMultiplicationSectionLexeme,
		startDivisionSectionLexeme:
		{
			var localOutput memoryCell

			for _, node2 := range node.childNodes {
				localOutput, err = node2.value(localOutput)
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
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			if _, err = fmt.Fprintf(node.tree.interpretCodeOptions.Output, "%c", output); err != nil {
				return
			}
		}
	case printNumberLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			if _, err = fmt.Fprintf(node.tree.interpretCodeOptions.Output, "%d", output); err != nil {
				return
			}
		}
	case inputCharacterLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			if _, err = fmt.Fscanf(node.tree.interpretCodeOptions.Input, "%c", output); err != nil {
				return
			}
		}
	case inputNumberLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			if _, err = fmt.Fscanf(node.tree.interpretCodeOptions.Input, "%d", output); err != nil {
				return
			}
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

			output = memoryCellFromIntegerConstraint(b[0])
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

			output, err = memoryCellFromBigInt(b)
			if err != nil {
				return
			}
		}
	case setSecondTimestampLexeme:
		output = memoryCellFromIntegerConstraint(time.Now().Unix())
	case setNanosecondTimestampLexeme:
		output = memoryCellFromIntegerConstraint(time.Now().UnixNano())
	case useStackIndexZeroLexeme:
		if node.tree == nil {
			err = ErrTreeUnfound
			return
		}

		node.tree.interpretCodeOptions.saveStackIndex = 0
	case useStackIndexOneLexeme:
		if node.tree == nil {
			err = ErrTreeUnfound
			return
		}

		node.tree.interpretCodeOptions.saveStackIndex = 1
	case useStackIndexSwappedLexeme:
		{
			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			if node.tree.interpretCodeOptions.saveStackIndex == 0 {
				node.tree.interpretCodeOptions.saveStackIndex = 1
			} else {
				node.tree.interpretCodeOptions.saveStackIndex = 0
			}
		}
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

			if len(saveStack) >= interpretCodeOptionsSaveStackMaxLen {
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

			output = memoryCellFromIntegerConstraint(len(saveStack))
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
			fileName := output.String() + FileExtensionForSaveStack

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

			fileName := output.String() + FileExtensionForSaveStack
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
				if len(saveStack) >= interpretCodeOptionsSaveStackMaxLen {
					err = ErrTreeSaveStackFull
					return
				}

				saveStack = append(saveStack, memoryCellFromIntegerConstraint(value))
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

			output = memoryCellFromIntegerConstraint(hash.Uint8(content))
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

			output = memoryCellFromIntegerConstraint(hash.Uint64(content))
		}
	case sortStackAscendingLexeme:
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

			saveStack.SortAscending()
		}
	case sortStackDescendingLexeme:
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

			saveStack.SortDescending()
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
	case reverseStackLexeme:
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

			saveStack.Reverse()
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

			for i := memoryCellFromIntegerConstraint(0); i < output; i++ {
				if len(saveStack) >= interpretCodeOptionsSaveStackMaxLen {
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

			for i := memoryCellFromIntegerConstraint(1); i < output; i++ {
				if len(saveStack) >= interpretCodeOptionsSaveStackMaxLen {
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

			fileName := output.String() + FileExtensionForSaveStack

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

			for i := len(node.tree.interpretCodeOptions.saveStacks) - 1; i >= 0; i-- {
				node.tree.interpretCodeOptions.saveStacks[i] =
					node.tree.interpretCodeOptions.saveStacks[i][:0]
			}

			output = 0
		}
	case filePathLexeme:
		{
			if len(node.data) == 0 {
				return
			}

			if node.tree == nil {
				err = ErrTreeUnfound
				return
			}

			tree := node.tree
			filePath := string(node.data)

			var fileAbsPath string
			fileAbsPath, err = filepath.Abs(filePath)
			if err != nil {
				return
			}

			var content []byte
			content, err = os.ReadFile(fileAbsPath)
			if err != nil {
				return
			}

			fileExt := filepath.Ext(filePath)

			switch fileExt {
			case FileExtensionForCode:
				{
					interpretCodeOptionsCloned := tree.interpretCodeOptions.Clone()

					interpretCodeOptionsCloned.WorkingDir = filepath.Dir(fileAbsPath)
					interpretCodeOptionsCloned.initialCurrentValue = output

					var outputUint64 uint64

					outputUint64, err = InterpretCode(content, interpretCodeOptionsCloned)
					if err != nil {
						return
					}

					output = memoryCellFromIntegerConstraint(outputUint64)
				}
			default:
				{
					var saveStackPtr *memoryCellCollection
					saveStackPtr, err = tree.saveStackPtr()
					if err != nil {
						return
					}
					saveStack := *saveStackPtr

					contentRunes := []rune(string(content))

					for i := len(contentRunes) - 1; i >= 0; i-- {
						if len(saveStack) >= interpretCodeOptionsSaveStackMaxLen {
							err = ErrTreeSaveStackFull
							return
						}

						saveStack = append(saveStack, memoryCellFromIntegerConstraint(contentRunes[i]))
					}

					*saveStackPtr = saveStack
				}
			}
		}
	case changeDirLexeme:
		{
			if err = os.Chdir(string(node.data)); err != nil {
				return
			}
		}
	default:
		err = ErrLexemeUnrecognized
	}

	return
}
