package main

type lexeme uint64

const (
	invalidLexeme lexeme = iota
	startAdditionSectionLexeme
	endAdditionSectionLexeme
	startSubtractionSectionLexeme
	endSubtractionSectionLexeme
	startJumpSectionLexeme
	endJumpSectionLexeme
	startCommentSectionLexeme
	endCommentSectionLexeme
	addOneLexeme
	addEightLexeme
	addStackPairLexeme
	subtractOneLexeme
	subtractEightLexeme
	subtractStackPairLexeme
	multiplyTwoLexeme
	multiplyEightLexeme
	multiplyStackPairLexeme
	divideTwoLexeme
	divideEightLexeme
	divideStackPairLexeme
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
	setSecondTimestampLexeme
	setNanosecondTimestampLexeme
	printCharacterLexeme
	printNumberLexeme
	inputCharacterLexeme
	inputNumberLexeme
	writeToFileLexeme
	loadFromFileLexeme
	deleteFileLexeme
	clearStackLexeme
	saveToStackLexeme
	loadFromStackLexeme
	hashStackOneByteLexeme
	hashStackEightByteLexeme
	modifierLexeme
	separatorLexeme
)

func produceLexemes(input []byte) (output []lexeme, err error) {
	output = make([]lexeme, 0, len(input))
	sectionStack := make([]lexeme, 0, len(input)/2)

	for _, r := range input {
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
					case modifierLexeme:
						output[outputLen-1] = subtractStackPairLexeme
					case subtractOneLexeme:
						output[outputLen-1] = subtractEightLexeme
					default:
						l = subtractEightLexeme
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
			case '@':
				if len(output) > 0 && output[len(output)-1] == setSecondTimestampLexeme {
					output[len(output)-1] = setNanosecondTimestampLexeme
				} else {
					l = setSecondTimestampLexeme
				}
			case ':':
				l = saveToStackLexeme
			case ';':
				l = loadFromStackLexeme
			case '#':
				if len(output) > 0 && output[len(output)-1] == hashStackOneByteLexeme {
					output[len(output)-1] = hashStackEightByteLexeme
				} else {
					l = hashStackOneByteLexeme
				}
			case '.':
				l = writeToFileLexeme
			case ',':
				l = loadFromFileLexeme
			case '|':
				if len(output) > 0 && output[len(output)-1] == deleteFileLexeme {
					output[len(output)-1] = clearStackLexeme
				} else {
					l = deleteFileLexeme
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
					err = ErrOverlapSectionCharacters
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
					err = ErrOverlapSectionCharacters
					return
				}
				sectionStack = sectionStack[:len(sectionStack)-1]
			case '<':
				l = startJumpSectionLexeme
				sectionStack = append(sectionStack, l)
			case '>':
				l = endJumpSectionLexeme
				if len(sectionStack) == 0 {
					err = ErrNoMatchSectionCharacters
					return
				}
				if sectionStack[len(sectionStack)-1] != startJumpSectionLexeme {
					err = ErrOverlapSectionCharacters
					return
				}
				sectionStack = sectionStack[:len(sectionStack)-1]
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
