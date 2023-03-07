package main

import (
	"strconv"
	"strings"
)

type lexeme uint64

const (
	invalidLexeme lexeme = iota
	startAdditionSectionLexeme
	endAdditionSectionLexeme
	startSubtractionSectionLexeme
	endSubtractionSectionLexeme
	startJumpIfPositiveSectionLexeme
	endJumpIfPositiveSectionLexeme
	startJumpIfZeroSectionLexeme
	endJumpIfZeroSectionLexeme
	startCommentSectionLexeme
	endCommentSectionLexeme
	addOneLexeme
	addEightLexeme
	addStackPairLexeme
	addStackWholeLexeme
	subtractOneLexeme
	subtractEightLexeme
	subtractStackPairLexeme
	subtractStackWholeLexeme
	multiplyTwoLexeme
	multiplyEightLexeme
	multiplyStackPairLexeme
	multiplyStackWholeLexeme
	divideTwoLexeme
	divideEightLexeme
	divideStackPairLexeme
	divideStackWholeLexeme
	squareLexeme
	cubeLexeme
	setZeroLexeme
	setOneByteLexeme
	setEightByteLexeme
	setOneKibibyteLexeme
	setEightKibibyteLexeme
	setOneMebibyteLexeme
	setEightMebibyteLexeme
	setOneGibibyteLexeme
	setEightGibibyteLexeme
	setRandomByteLexeme
	setRandomMaxLexeme
	setSecondTimestampLexeme
	setNanosecondTimestampLexeme
	printCharacterLexeme
	printNumberLexeme
	inputCharacterLexeme
	inputNumberLexeme
	writeStackToFileLexeme
	readStackFromFileLexeme
	deleteFileLexeme
	clearStackLexeme
	resetStateLexeme
	pushStackLexeme
	popStackLexeme
	saveStackUseIndexZeroLexeme
	saveStackUseIndexOneLexeme
	hashStackOneByteLexeme
	hashStackEightByteLexeme
	modifierLexeme
	separatorLexeme
)

func (lexeme lexeme) String() string {
	var builder strings.Builder

	switch lexeme {
	case invalidLexeme:
		builder.WriteString("INVALID")
	case startAdditionSectionLexeme:
		builder.WriteString("START-ADD-SECT")
	case endAdditionSectionLexeme:
		builder.WriteString("END-ADD-SECT")
	case startSubtractionSectionLexeme:
		builder.WriteString("START-SUB-SECT")
	case endSubtractionSectionLexeme:
		builder.WriteString("END-SUB-SECT")
	case startJumpIfPositiveSectionLexeme:
		builder.WriteString("START-JMP-IF-POS-SECT")
	case endJumpIfPositiveSectionLexeme:
		builder.WriteString("END-JMP-IF-POS-SECT")
	case startJumpIfZeroSectionLexeme:
		builder.WriteString("START-JMP-IF-ZERO-SECT")
	case endJumpIfZeroSectionLexeme:
		builder.WriteString("END-JMP-IF-ZERO-SECT")
	case startCommentSectionLexeme:
		builder.WriteString("START-CMNT-SECT")
	case endCommentSectionLexeme:
		builder.WriteString("END-CMNT-SECT")
	case addOneLexeme:
		builder.WriteString("ADD-ONE")
	case addEightLexeme:
		builder.WriteString("ADD-EIGHT")
	case addStackPairLexeme:
		builder.WriteString("ADD-STACK-PAIR")
	case addStackWholeLexeme:
		builder.WriteString("ADD-STACK-WHOLE")
	case subtractOneLexeme:
		builder.WriteString("SUB-ONE")
	case subtractEightLexeme:
		builder.WriteString("SUB-EIGHT")
	case subtractStackPairLexeme:
		builder.WriteString("SUB-STACK-PAIR")
	case subtractStackWholeLexeme:
		builder.WriteString("SUB-STACK-WHOLE")
	case multiplyTwoLexeme:
		builder.WriteString("MULT-TWO")
	case multiplyEightLexeme:
		builder.WriteString("MULT-EIGHT")
	case multiplyStackPairLexeme:
		builder.WriteString("MULT-STACK-PAIR")
	case multiplyStackWholeLexeme:
		builder.WriteString("MULT-STACK-WHOLE")
	case divideTwoLexeme:
		builder.WriteString("DIV-TWO")
	case divideEightLexeme:
		builder.WriteString("DIV-EIGHT")
	case divideStackPairLexeme:
		builder.WriteString("DIV-STACK-PAIR")
	case divideStackWholeLexeme:
		builder.WriteString("DIV-STACK-WHOLE")
	case squareLexeme:
		builder.WriteString("SQUARE")
	case cubeLexeme:
		builder.WriteString("CUBE")
	case setZeroLexeme:
		builder.WriteString("SET-ZERO")
	case setOneByteLexeme:
		builder.WriteString("SET-ONE-BYTE")
	case setEightByteLexeme:
		builder.WriteString("SET-EIGHT-BYTE")
	case setOneKibibyteLexeme:
		builder.WriteString("SET-ONE-KIBI")
	case setEightKibibyteLexeme:
		builder.WriteString("SET-EIGHT-KIBI")
	case setOneMebibyteLexeme:
		builder.WriteString("SET-ONE-MEBI")
	case setEightMebibyteLexeme:
		builder.WriteString("SET-EIGHT-MEBI")
	case setOneGibibyteLexeme:
		builder.WriteString("SET-ONE-GIBI")
	case setEightGibibyteLexeme:
		builder.WriteString("SET-EIGHT-GIBI")
	case setRandomByteLexeme:
		builder.WriteString("SET-RAND-BYTE")
	case setRandomMaxLexeme:
		builder.WriteString("SET-RAND-MAX")
	case setSecondTimestampLexeme:
		builder.WriteString("SET-SEC-TIME")
	case setNanosecondTimestampLexeme:
		builder.WriteString("SET-NANO-TIME")
	case printCharacterLexeme:
		builder.WriteString("PRINT-CHAR")
	case printNumberLexeme:
		builder.WriteString("PRINT-NUM")
	case writeStackToFileLexeme:
		builder.WriteString("WRITE-STACK-FILE")
	case readStackFromFileLexeme:
		builder.WriteString("READ-STACK-FILE")
	case deleteFileLexeme:
		builder.WriteString("DELETE-STACK-FILE")
	case clearStackLexeme:
		builder.WriteString("CLEAR-STACK")
	case pushStackLexeme:
		builder.WriteString("PUSH-STACK")
	case popStackLexeme:
		builder.WriteString("POP-STACK")
	case saveStackUseIndexZeroLexeme:
		builder.WriteString("USE-STACK-ZERO")
	case saveStackUseIndexOneLexeme:
		builder.WriteString("USE-STACK-ONE")
	case hashStackOneByteLexeme:
		builder.WriteString("HASH-STACK-ONE-BYTE")
	case hashStackEightByteLexeme:
		builder.WriteString("HASH-STACK-EIGHT-BYTE")
	case modifierLexeme:
		builder.WriteString("MODIFIER")
	case separatorLexeme:
		builder.WriteString("SEP")
	default:
		builder.WriteString("UNKNOWN")
	}

	builder.WriteByte(' ')
	builder.WriteByte('[')
	builder.WriteString(strconv.FormatUint(uint64(lexeme), 10))
	builder.WriteByte(']')

	return builder.String()
}

func produceLexemes(input []byte) (output []lexeme, err error) {
	output = make([]lexeme, 0, len(input))
	sectionStack := make([]lexeme, 0, len(input)/2+1)

	for i, r := range input {
		l := invalidLexeme

		if len(sectionStack) > 0 && sectionStack[len(sectionStack)-1] == startCommentSectionLexeme {
			switch r {
			case '{':
				l = startCommentSectionLexeme
				sectionStack = append(sectionStack, l)
			case '}':
				l = endCommentSectionLexeme
				sectionStack = sectionStack[:len(sectionStack)-1]
			}
		} else {
			switch r {
			case '+':
				{
					outputLen := len(output)
					lastOutput := invalidLexeme

					if outputLen > 0 {
						lastOutput = output[outputLen-1]
					}

					switch lastOutput {
					case addStackPairLexeme:
						output[outputLen-1] = addStackWholeLexeme
					case modifierLexeme:
						output[outputLen-1] = addStackPairLexeme
					case addOneLexeme:
						output[outputLen-1] = addEightLexeme
					default:
						l = addOneLexeme
					}
				}
			case '-':
				{
					outputLen := len(output)
					lastOutput := invalidLexeme

					if outputLen > 0 {
						lastOutput = output[outputLen-1]
					}

					switch lastOutput {
					case subtractStackPairLexeme:
						output[outputLen-1] = subtractStackWholeLexeme
					case modifierLexeme:
						output[outputLen-1] = subtractStackPairLexeme
					case subtractOneLexeme:
						output[outputLen-1] = subtractEightLexeme
					default:
						l = subtractOneLexeme
					}
				}
			case '*':
				{
					outputLen := len(output)
					lastOutput := invalidLexeme

					if outputLen > 0 {
						lastOutput = output[outputLen-1]
					}

					switch lastOutput {
					case multiplyStackPairLexeme:
						output[outputLen-1] = multiplyStackWholeLexeme
					case modifierLexeme:
						output[outputLen-1] = multiplyStackPairLexeme
					case multiplyTwoLexeme:
						output[outputLen-1] = multiplyEightLexeme
					default:
						l = multiplyTwoLexeme
					}
				}
			case '/':
				{
					outputLen := len(output)
					lastOutput := invalidLexeme

					if outputLen > 0 {
						lastOutput = output[outputLen-1]
					}

					switch lastOutput {
					case divideStackPairLexeme:
						output[outputLen-1] = divideStackWholeLexeme
					case modifierLexeme:
						output[outputLen-1] = divideStackPairLexeme
					case divideTwoLexeme:
						output[outputLen-1] = divideEightLexeme
					default:
						l = divideTwoLexeme
					}
				}
			case '^':
				if len(output) > 0 && output[len(output)-1] == squareLexeme {
					output[len(output)-1] = cubeLexeme
				} else {
					l = squareLexeme
				}
			case '!':
				if len(output) > 0 && output[len(output)-1] == printCharacterLexeme {
					output[len(output)-1] = printNumberLexeme
				} else {
					l = printCharacterLexeme
				}
			case '?':
				if len(output) > 0 && output[len(output)-1] == inputCharacterLexeme {
					output[len(output)-1] = inputNumberLexeme
				} else {
					l = inputCharacterLexeme
				}
			case '~':
				l = setZeroLexeme
			case '\'':
				{
					outputLen := len(output)
					lastOutput := invalidLexeme

					if outputLen > 0 {
						lastOutput = output[outputLen-1]
					}

					switch lastOutput {
					case setOneMebibyteLexeme:
						output[outputLen-1] = setEightMebibyteLexeme
					case modifierLexeme:
						output[outputLen-1] = setOneMebibyteLexeme
					case setOneByteLexeme:
						output[outputLen-1] = setEightByteLexeme
					default:
						l = setOneByteLexeme
					}
				}
			case '"':
				{
					outputLen := len(output)
					lastOutput := invalidLexeme

					if outputLen > 0 {
						lastOutput = output[outputLen-1]
					}

					switch lastOutput {
					case setOneGibibyteLexeme:
						output[outputLen-1] = setEightGibibyteLexeme
					case modifierLexeme:
						output[outputLen-1] = setOneGibibyteLexeme
					case setOneKibibyteLexeme:
						output[outputLen-1] = setEightKibibyteLexeme
					default:
						l = setOneKibibyteLexeme
					}
				}
			case '`':
				if len(output) > 0 && output[len(output)-1] == setRandomByteLexeme {
					output[len(output)-1] = setRandomMaxLexeme
				} else {
					l = setRandomByteLexeme
				}
			case '@':
				if len(output) > 0 && output[len(output)-1] == setSecondTimestampLexeme {
					output[len(output)-1] = setNanosecondTimestampLexeme
				} else {
					l = setSecondTimestampLexeme
				}
			case '$':
				if len(output) > 0 && output[len(output)-1] == saveStackUseIndexZeroLexeme {
					output[len(output)-1] = saveStackUseIndexOneLexeme
				} else {
					l = saveStackUseIndexZeroLexeme
				}
			case ':':
				l = pushStackLexeme
			case ';':
				l = popStackLexeme
			case '#':
				if len(output) > 0 && output[len(output)-1] == hashStackOneByteLexeme {
					output[len(output)-1] = hashStackEightByteLexeme
				} else {
					l = hashStackOneByteLexeme
				}
			case '.':
				l = writeStackToFileLexeme
			case ',':
				l = readStackFromFileLexeme
			case '|':
				{
					outputLen := len(output)
					lastOutput := invalidLexeme

					if outputLen > 0 {
						lastOutput = output[outputLen-1]
					}

					switch lastOutput {
					case modifierLexeme:
						output[outputLen-1] = resetStateLexeme
					case deleteFileLexeme:
						output[outputLen-1] = clearStackLexeme
					default:
						l = deleteFileLexeme
					}
				}
			case '(':
				l = startAdditionSectionLexeme
				sectionStack = append(sectionStack, l)
			case ')':
				l = endAdditionSectionLexeme
				if len(sectionStack) == 0 {
					err = ErrNoMatchSectionCharacters
					return
				}
				if sectionStack[len(sectionStack)-1] != startAdditionSectionLexeme {
					err = ErrLexemeSectionStackNoMatch
					return
				}
				sectionStack = sectionStack[:len(sectionStack)-1]
			case '[':
				l = startSubtractionSectionLexeme
				sectionStack = append(sectionStack, l)
			case ']':
				l = endSubtractionSectionLexeme
				if len(sectionStack) == 0 {
					err = ErrNoMatchSectionCharacters
					return
				}
				if sectionStack[len(sectionStack)-1] != startSubtractionSectionLexeme {
					err = ErrLexemeSectionStackNoMatch
					return
				}
				sectionStack = sectionStack[:len(sectionStack)-1]
			case '<':
				if len(output) > 0 && output[len(output)-1] == startJumpIfPositiveSectionLexeme {
					if len(sectionStack) == 0 {
						err = ErrLexemeSectionStackEmpty
						return
					}
					if sectionStack[len(sectionStack)-1] != startJumpIfPositiveSectionLexeme {
						err = ErrLexemeSectionStackNoMatch
						return
					}
					sectionStack[len(sectionStack)-1] = startJumpIfZeroSectionLexeme
					output[len(output)-1] = startJumpIfZeroSectionLexeme
				} else {
					l = startJumpIfPositiveSectionLexeme
					sectionStack = append(sectionStack, l)
				}
			case '>':
				{
					var localLexeme lexeme

					if len(output) > 0 && output[len(output)-1] == endJumpIfPositiveSectionLexeme {
						localLexeme = endJumpIfZeroSectionLexeme
						output[len(output)-1] = localLexeme
					} else {
						localLexeme = endJumpIfPositiveSectionLexeme
						l = localLexeme
					}

					if len(sectionStack) == 0 {
						err = ErrNoMatchSectionCharacters
						return
					}

					switch localLexeme {
					case endJumpIfPositiveSectionLexeme:
						if i+1 < len(input) && input[i+1] == r {
							localLexeme = invalidLexeme
						} else {
							localLexeme = startJumpIfPositiveSectionLexeme
						}
					case endJumpIfZeroSectionLexeme:
						localLexeme = startJumpIfZeroSectionLexeme
					default:
						err = ErrLexemeSectionStackNoMatch
						return
					}

					if localLexeme != invalidLexeme {
						if sectionStack[len(sectionStack)-1] != localLexeme {
							err = ErrLexemeSectionStackNoMatch
							return
						}

						sectionStack = sectionStack[:len(sectionStack)-1]
					}
				}
			case '{':
				l = startCommentSectionLexeme
				sectionStack = append(sectionStack, l)
			case ' ', '\n', '\r', '\t', '\v', '\f', 0x85, 0xa0:
				if len(output) == 0 || output[len(output)-1] != separatorLexeme {
					l = separatorLexeme
				}
			case '%':
				l = modifierLexeme
			default:
				err = ErrLexemeUnrecognized
				return
			}
		}

		if l != invalidLexeme {
			output = append(output, l)
		}
	}

	if len(sectionStack) != 0 {
		err = ErrNoMatchSectionCharacters
	}

	return
}
