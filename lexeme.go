package main

import "unicode"

type Lexeme uint64

const (
	InvalidLexeme Lexeme = iota
	StartAdditionSectionLexeme
	EndAdditionSectionLexeme
	StartSubtractionSectionLexeme
	EndSubtractionSectionLexeme
	IncrementLexexme
	DecrementLexeme
	DoubleLexeme
	HalfLexeme
	SquareLexeme
	CubeLexeme
	MinLexeme
	MaxLexeme
	PrintCharacterLexeme
	PrintNumberLexeme
	SeparatorLexeme
)

func produceLexemes(input []byte) (output []Lexeme, err error) {
	output = make([]Lexeme, 0, len(input))

	var sectionStack []Lexeme

	for _, r := range input {
		l := InvalidLexeme

		if unicode.IsSpace(rune(r)) {
			if len(output) == 0 || output[len(output)-1] != SeparatorLexeme {
				l = SeparatorLexeme
			}
		} else {
			switch r {
			case '+':
				l = IncrementLexexme
			case '-':
				l = DecrementLexeme
			case '*':
				l = DoubleLexeme
			case '/':
				l = HalfLexeme
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
				l = MinLexeme
			case '"':
				l = MaxLexeme
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
