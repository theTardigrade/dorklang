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
	incrementOneLexeme
	incrementEightLexeme
	decrementOneLexeme
	decrementEightLexeme
	multiplyTwoLexeme
	multiplyEightLexeme
	divideTwoLexeme
	divideEightLexeme
	squareLexeme
	cubeLexeme
	minimumLexeme
	middleLexeme
	maximumLexeme
	printCharacterLexeme
	printNumberLexeme
	inputCharacterLexeme
	inputNumberLexeme
	saveLexeme
	loadLexeme
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
				if len(output) > 0 && output[len(output)-1] == incrementOneLexeme {
					output[len(output)-1] = incrementEightLexeme
				} else {
					l = incrementOneLexeme
				}
			case '-':
				if len(output) > 0 && output[len(output)-1] == decrementOneLexeme {
					output[len(output)-1] = decrementEightLexeme
				} else {
					l = decrementOneLexeme
				}
			case '*':
				if len(output) > 0 && output[len(output)-1] == multiplyTwoLexeme {
					output[len(output)-1] = multiplyEightLexeme
				} else {
					l = multiplyTwoLexeme
				}
			case '/':
				if len(output) > 0 && output[len(output)-1] == divideTwoLexeme {
					output[len(output)-1] = divideEightLexeme
				} else {
					l = divideTwoLexeme
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
			case '\'':
				l = minimumLexeme
			case '~':
				l = middleLexeme
			case '"':
				l = maximumLexeme
			case ':':
				l = saveLexeme
			case ';':
				l = loadLexeme
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
