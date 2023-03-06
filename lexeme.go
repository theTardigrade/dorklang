package main

type Lexeme uint64

const (
	InvalidLexeme Lexeme = iota
	StartAdditionSectionLexeme
	EndAdditionSectionLexeme
	StartSubtractionSectionLexeme
	EndSubtractionSectionLexeme
	StartJumpSectionLexeme
	EndJumpSectionLexeme
	StartCommentSectionLexeme
	EndCommentSectionLexeme
	IncrementOneLexeme
	IncrementEightLexeme
	DecrementOneLexeme
	DecrementEightLexeme
	MultiplyTwoLexeme
	MultiplyEightLexeme
	DivideTwoLexeme
	DivideEightLexeme
	SquareLexeme
	CubeLexeme
	MinimumLexeme
	MiddleLexeme
	MaximumLexeme
	PrintCharacterLexeme
	PrintNumberLexeme
	SaveLexeme
	LoadLexeme
	SeparatorLexeme
)

func produceLexemes(input []byte) (output []Lexeme, err error) {
	output = make([]Lexeme, 0, len(input))

	var sectionStack []Lexeme

	for _, r := range input {
		l := InvalidLexeme

		if len(sectionStack) > 0 && sectionStack[len(sectionStack)-1] == StartCommentSectionLexeme {
			switch r {
			case '{':
				l = StartCommentSectionLexeme
				sectionStack = append(sectionStack, l)
			case '}':
				l = EndCommentSectionLexeme
				sectionStack = sectionStack[:len(sectionStack)-1]
			}
		} else {
			switch r {
			case '+':
				if len(output) > 0 && output[len(output)-1] == IncrementOneLexeme {
					output[len(output)-1] = IncrementEightLexeme
				} else {
					l = IncrementOneLexeme
				}
			case '-':
				if len(output) > 0 && output[len(output)-1] == DecrementOneLexeme {
					output[len(output)-1] = DecrementEightLexeme
				} else {
					l = DecrementOneLexeme
				}
			case '*':
				if len(output) > 0 && output[len(output)-1] == MultiplyTwoLexeme {
					output[len(output)-1] = MultiplyEightLexeme
				} else {
					l = MultiplyTwoLexeme
				}
			case '/':
				if len(output) > 0 && output[len(output)-1] == DivideTwoLexeme {
					output[len(output)-1] = DivideEightLexeme
				} else {
					l = DivideTwoLexeme
				}
			case '^':
				if len(output) > 0 && output[len(output)-1] == SquareLexeme {
					output[len(output)-1] = CubeLexeme
				} else {
					l = SquareLexeme
				}
			case '!':
				if len(output) > 0 && output[len(output)-1] == PrintCharacterLexeme {
					output[len(output)-1] = PrintNumberLexeme
				} else {
					l = PrintCharacterLexeme
				}
			case '\'':
				l = MinimumLexeme
			case '~':
				l = MiddleLexeme
			case '"':
				l = MaximumLexeme
			case ':':
				l = SaveLexeme
			case ';':
				l = LoadLexeme
			case '(':
				l = StartAdditionSectionLexeme
				sectionStack = append(sectionStack, l)
			case ')':
				l = EndAdditionSectionLexeme
				if len(sectionStack) == 0 {
					err = ErrNoMatchSectionCharacters
					return
				}
				if sectionStack[len(sectionStack)-1] != StartAdditionSectionLexeme {
					err = ErrOverlapSectionCharacters
					return
				}
				sectionStack = sectionStack[:len(sectionStack)-1]
			case '[':
				l = StartSubtractionSectionLexeme
				sectionStack = append(sectionStack, l)
			case ']':
				l = EndSubtractionSectionLexeme
				if len(sectionStack) == 0 {
					err = ErrNoMatchSectionCharacters
					return
				}
				if sectionStack[len(sectionStack)-1] != StartSubtractionSectionLexeme {
					err = ErrOverlapSectionCharacters
					return
				}
				sectionStack = sectionStack[:len(sectionStack)-1]
			case '<':
				l = StartJumpSectionLexeme
				sectionStack = append(sectionStack, l)
			case '>':
				l = EndJumpSectionLexeme
				if len(sectionStack) == 0 {
					err = ErrNoMatchSectionCharacters
					return
				}
				if sectionStack[len(sectionStack)-1] != StartJumpSectionLexeme {
					err = ErrOverlapSectionCharacters
					return
				}
				sectionStack = sectionStack[:len(sectionStack)-1]
			case '{':
				l = StartCommentSectionLexeme
				sectionStack = append(sectionStack, l)
			case ' ', '\n', '\r', '\t', '\v', '\f', 0x85, 0xa0:
				if len(output) == 0 || output[len(output)-1] != SeparatorLexeme {
					l = SeparatorLexeme
				}
			default:
				err = ErrLexemeUnrecognized
				return
			}
		}

		if l != InvalidLexeme {
			output = append(output, l)
		}
	}

	if len(sectionStack) != 0 {
		err = ErrNoMatchSectionCharacters
	}

	return
}
