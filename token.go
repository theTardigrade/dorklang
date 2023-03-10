package main

import (
	"unicode"
)

type token struct {
	lex  lexeme
	data []byte
}

func produceTokens(input []byte) (output []token, err error) {
	output = make([]token, 0, len(input))
	sectionStack := make([]lexeme, 0, len(input)/2+1)

	for i, r := range input {
		l := invalidLexeme
		var d []byte

		sectionStackTopLexeme := invalidLexeme
		if len(sectionStack) > 0 {
			sectionStackTopLexeme = sectionStack[len(sectionStack)-1]
		}

		if sectionStackTopLexeme == startCommentSectionLexeme {
			switch r {
			case '{':
				if len(output) > 0 && output[len(output)-1].lex == startCommentSectionLexeme {
					if len(sectionStack) == 0 {
						err = ErrLexemeSectionStackEmpty
						return
					}
					if sectionStack[len(sectionStack)-1] != startCommentSectionLexeme {
						err = ErrLexemeSectionStackNoMatch
						return
					}
					sectionStack[len(sectionStack)-1] = startReadFileSectionLexeme
					output[len(output)-1].lex = startReadFileSectionLexeme
				} else {
					l = startCommentSectionLexeme
					sectionStack = append(sectionStack, l)
				}
			case '}':
				l = endCommentSectionLexeme
				sectionStack = sectionStack[:len(sectionStack)-1]
			}
		} else if sectionStackTopLexeme == startReadFileSectionLexeme {
			switch r {
			case '}':
				{
					if i+1 >= len(input) || input[i+1] != r {
						err = ErrLexemeSectionStackNoMatch
						return
					}

					input[i+1] = ' '

					if len(sectionStack) == 0 {
						err = ErrNoMatchSectionCharacters
						return
					}

					l = endReadFileSectionLexeme
					sectionStack = sectionStack[:len(sectionStack)-1]
				}
			default:
				{
					var handled bool

					if unicode.IsSpace(rune(r)) {
						if len(output) == 0 || output[len(output)-1].lex != separatorLexeme {
							l = separatorLexeme
						}

						handled = true
					}

					if !handled {
						if len(output) == 0 || output[len(output)-1].lex != plaintextLexeme {
							l = plaintextLexeme
							d = append(d, r)
						}

						if len(output) != 0 && output[len(output)-1].lex == plaintextLexeme {
							output[len(output)-1].data = append(output[len(output)-1].data, r)
						}
					}
				}
			}
		} else {
			switch r {
			case '+':
				{
					outputLen := len(output)
					lastOutput := invalidLexeme

					if outputLen > 0 {
						lastOutput = output[outputLen-1].lex
					}

					switch lastOutput {
					case addStackPairLexeme:
						output[outputLen-1].lex = addStackWholeLexeme
					case modifierLexeme:
						output[outputLen-1].lex = addStackPairLexeme
					case addOneLexeme:
						output[outputLen-1].lex = addEightLexeme
					default:
						l = addOneLexeme
					}
				}
			case '-':
				{
					outputLen := len(output)
					lastOutput := invalidLexeme

					if outputLen > 0 {
						lastOutput = output[outputLen-1].lex
					}

					switch lastOutput {
					case subtractStackPairLexeme:
						output[outputLen-1].lex = subtractStackWholeLexeme
					case modifierLexeme:
						output[outputLen-1].lex = subtractStackPairLexeme
					case subtractOneLexeme:
						output[outputLen-1].lex = subtractEightLexeme
					default:
						l = subtractOneLexeme
					}
				}
			case '*':
				{
					outputLen := len(output)
					lastOutput := invalidLexeme

					if outputLen > 0 {
						lastOutput = output[outputLen-1].lex
					}

					switch lastOutput {
					case multiplyStackPairLexeme:
						output[outputLen-1].lex = multiplyStackWholeLexeme
					case modifierLexeme:
						output[outputLen-1].lex = multiplyStackPairLexeme
					case multiplyTwoLexeme:
						output[outputLen-1].lex = multiplyEightLexeme
					default:
						l = multiplyTwoLexeme
					}
				}
			case '/':
				{
					outputLen := len(output)
					lastOutput := invalidLexeme

					if outputLen > 0 {
						lastOutput = output[outputLen-1].lex
					}

					switch lastOutput {
					case divideStackPairLexeme:
						output[outputLen-1].lex = divideStackWholeLexeme
					case modifierLexeme:
						output[outputLen-1].lex = divideStackPairLexeme
					case divideTwoLexeme:
						output[outputLen-1].lex = divideEightLexeme
					default:
						l = divideTwoLexeme
					}
				}
			case '^':
				if len(output) > 0 && output[len(output)-1].lex == squareLexeme {
					output[len(output)-1].lex = cubeLexeme
				} else {
					l = squareLexeme
				}
			case '!':
				if len(output) > 0 && output[len(output)-1].lex == printCharacterLexeme {
					output[len(output)-1].lex = printNumberLexeme
				} else {
					l = printCharacterLexeme
				}
			case '?':
				if len(output) > 0 && output[len(output)-1].lex == inputCharacterLexeme {
					output[len(output)-1].lex = inputNumberLexeme
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
						lastOutput = output[outputLen-1].lex
					}

					switch lastOutput {
					case setOneMebibyteLexeme:
						output[outputLen-1].lex = setEightMebibyteLexeme
					case modifierLexeme:
						output[outputLen-1].lex = setOneMebibyteLexeme
					case setOneByteLexeme:
						output[outputLen-1].lex = setEightByteLexeme
					default:
						l = setOneByteLexeme
					}
				}
			case '"':
				{
					outputLen := len(output)
					lastOutput := invalidLexeme

					if outputLen > 0 {
						lastOutput = output[outputLen-1].lex
					}

					switch lastOutput {
					case setOneGibibyteLexeme:
						output[outputLen-1].lex = setEightGibibyteLexeme
					case modifierLexeme:
						output[outputLen-1].lex = setOneGibibyteLexeme
					case setOneKibibyteLexeme:
						output[outputLen-1].lex = setEightKibibyteLexeme
					default:
						l = setOneKibibyteLexeme
					}
				}
			case '`':
				if len(output) > 0 && output[len(output)-1].lex == setRandomByteLexeme {
					output[len(output)-1].lex = setRandomMaxLexeme
				} else {
					l = setRandomByteLexeme
				}
			case '@':
				if len(output) > 0 && output[len(output)-1].lex == setSecondTimestampLexeme {
					output[len(output)-1].lex = setNanosecondTimestampLexeme
				} else {
					l = setSecondTimestampLexeme
				}
			case '$':
				if len(output) > 0 && output[len(output)-1].lex == saveStackUseIndexZeroLexeme {
					output[len(output)-1].lex = saveStackUseIndexOneLexeme
				} else {
					l = saveStackUseIndexZeroLexeme
				}
			case ':':
				if len(output) > 0 && output[len(output)-1].lex == modifierLexeme {
					output[len(output)-1].lex = countStackLexeme
				} else {
					l = pushStackLexeme
				}
			case ';':
				if len(output) > 0 && output[len(output)-1].lex == modifierLexeme {
					output[len(output)-1].lex = popStackRandomLexeme
				} else {
					l = popStackLastLexeme
				}
			case '#':
				if len(output) > 0 && output[len(output)-1].lex == hashStackOneByteLexeme {
					output[len(output)-1].lex = hashStackEightByteLexeme
				} else {
					l = hashStackOneByteLexeme
				}
			case '&':
				{
					outputLen := len(output)
					lastOutput := invalidLexeme

					if outputLen > 0 {
						lastOutput = output[outputLen-1].lex
					}

					switch lastOutput {
					case logicalAndStackPairLexeme:
						output[outputLen-1].lex = logicalAndStackWholeLexeme
					case modifierLexeme:
						output[outputLen-1].lex = logicalAndStackPairLexeme
					}
				}
			case '\\':
				l = invertLexeme
			case '.':
				l = writeStackToFileLexeme
			case ',':
				l = readStackFromFileLexeme
			case 's':
				{
					outputLen := len(output)
					lastOutput := invalidLexeme

					if outputLen > 0 {
						lastOutput = output[outputLen-1].lex
					}

					switch lastOutput {
					case modifierLexeme:
						output[outputLen-1].lex = shuffleStackLexeme
					default:
						l = sortStackLexeme
					}
				}
			case 'x':
				l = swapStackTopLexeme
			case 'i':
				if len(output) > 0 && output[len(output)-1].lex == iotaFromZeroLexeme {
					output[len(output)-1].lex = iotaFromOneLexeme
				} else {
					l = iotaFromZeroLexeme
				}
			case '|':
				{
					outputLen := len(output)
					lastOutput := invalidLexeme

					if outputLen > 0 {
						lastOutput = output[outputLen-1].lex
					}

					switch lastOutput {
					case modifierLexeme:
						output[outputLen-1].lex = resetStateLexeme
					case deleteFileLexeme:
						output[outputLen-1].lex = clearStackLexeme
					default:
						l = deleteFileLexeme
					}
				}
			case '(':
				if len(output) > 0 && output[len(output)-1].lex == startAdditionSectionLexeme {
					if len(sectionStack) == 0 {
						err = ErrLexemeSectionStackEmpty
						return
					}
					if sectionStack[len(sectionStack)-1] != startAdditionSectionLexeme {
						err = ErrLexemeSectionStackNoMatch
						return
					}
					sectionStack[len(sectionStack)-1] = startMultiplicationSectionLexeme
					output[len(output)-1].lex = startMultiplicationSectionLexeme
				} else {
					l = startAdditionSectionLexeme
					sectionStack = append(sectionStack, l)
				}
			case ')':
				{
					var localLexeme lexeme

					if len(output) > 0 && output[len(output)-1].lex == endAdditionSectionLexeme {
						localLexeme = endMultiplicationSectionLexeme
						output[len(output)-1].lex = localLexeme
					} else {
						localLexeme = endAdditionSectionLexeme
						l = localLexeme
					}

					if len(sectionStack) == 0 {
						err = ErrNoMatchSectionCharacters
						return
					}

					switch localLexeme {
					case endAdditionSectionLexeme:
						if i+1 < len(input) && input[i+1] == r {
							localLexeme = invalidLexeme
						} else {
							localLexeme = startAdditionSectionLexeme
						}
					case endMultiplicationSectionLexeme:
						localLexeme = startMultiplicationSectionLexeme
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
			case '[':
				if len(output) > 0 && output[len(output)-1].lex == startSubtractionSectionLexeme {
					if len(sectionStack) == 0 {
						err = ErrLexemeSectionStackEmpty
						return
					}
					if sectionStack[len(sectionStack)-1] != startSubtractionSectionLexeme {
						err = ErrLexemeSectionStackNoMatch
						return
					}
					sectionStack[len(sectionStack)-1] = startDivisionSectionLexeme
					output[len(output)-1].lex = startDivisionSectionLexeme
				} else {
					l = startSubtractionSectionLexeme
					sectionStack = append(sectionStack, l)
				}
			case ']':
				{
					var localLexeme lexeme

					if len(output) > 0 && output[len(output)-1].lex == endSubtractionSectionLexeme {
						localLexeme = endDivisionSectionLexeme
						output[len(output)-1].lex = localLexeme
					} else {
						localLexeme = endSubtractionSectionLexeme
						l = localLexeme
					}

					if len(sectionStack) == 0 {
						err = ErrNoMatchSectionCharacters
						return
					}

					switch localLexeme {
					case endSubtractionSectionLexeme:
						if i+1 < len(input) && input[i+1] == r {
							localLexeme = invalidLexeme
						} else {
							localLexeme = startSubtractionSectionLexeme
						}
					case endDivisionSectionLexeme:
						localLexeme = startDivisionSectionLexeme
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
			case '<':
				if len(output) > 0 && output[len(output)-1].lex == startJumpIfPositiveSectionLexeme {
					if len(sectionStack) == 0 {
						err = ErrLexemeSectionStackEmpty
						return
					}
					if sectionStack[len(sectionStack)-1] != startJumpIfPositiveSectionLexeme {
						err = ErrLexemeSectionStackNoMatch
						return
					}
					sectionStack[len(sectionStack)-1] = startJumpIfZeroSectionLexeme
					output[len(output)-1].lex = startJumpIfZeroSectionLexeme
				} else {
					l = startJumpIfPositiveSectionLexeme
					sectionStack = append(sectionStack, l)
				}
			case '>':
				{
					var localLexeme lexeme

					if len(output) > 0 && output[len(output)-1].lex == endJumpIfPositiveSectionLexeme {
						localLexeme = endJumpIfZeroSectionLexeme
						output[len(output)-1].lex = localLexeme
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
			case '%':
				l = modifierLexeme
			default:
				{
					var handled bool

					if unicode.IsSpace(rune(r)) {
						if len(output) == 0 || output[len(output)-1].lex != separatorLexeme {
							l = separatorLexeme
						}

						handled = true
					}

					if !handled {
						err = ErrLexemeUnrecognized
						return
					}
				}
			}
		}

		if l != invalidLexeme {
			output = append(output, token{
				lex:  l,
				data: d,
			})
		}
	}

	if len(sectionStack) != 0 {
		err = ErrNoMatchSectionCharacters
	}

	return
}
