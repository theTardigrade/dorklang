package dorklang

import (
	"log"
	"strings"
)

// func (collection tokenCollection) peekToken(i int) (t token, found bool) {
// 	if i < 0 || i >= len(collection) {
// 		return
// 	}

// 	t = collection[i]
// 	found = true

// 	return
// }

// func (collection tokenCollection) peekNextToken(i int, ignoreLexemes ...lexeme) (t token, index int, found bool) {
// 	for j := i + 1; j < len(collection); j++ {
// 		t = collection[j]
// 		found = true

// 		for _, l := range ignoreLexemes {
// 			if t.lex == l {
// 				found = false
// 				break
// 			}
// 		}

// 		if found {
// 			index = j
// 			break
// 		}
// 	}

// 	return
// }

// func (collection tokenCollection) peekNextUsefulToken(i int) (t token, index int, found bool) {
// 	t, index, found = collection.peekNextToken(i, separatorLexeme, emptyLexeme)

// 	return
// }

func (collection tokenCollection) peekPrevToken(i int, ignoreLexemes ...lexeme) (t token, index int, found bool) {
	for j := i - 1; j >= 0; j-- {
		t = collection[j]
		found = true

		for _, l := range ignoreLexemes {
			if t.lex == l {
				found = false
				break
			}
		}

		if found {
			index = j
			break
		}
	}

	return
}

func (collection tokenCollection) peekPrevUsefulToken(i int) (t token, index int, found bool) {
	t, index, found = collection.peekPrevToken(i, separatorLexeme, emptyLexeme)

	return
}

func (t token) log(indent int) {
	var builder strings.Builder

	for i := 0; i < indent; i++ {
		builder.WriteByte('\t')
	}

	builder.WriteString("lexeme: ")
	builder.WriteString(t.lex.String())

	log.Println(builder.String())

	if t.lex == parentLexeme {
		for _, t2 := range t.childCollection {
			t2.log(indent + 1)
		}
	}
}

func (collection tokenCollection) log() {
	for _, t := range collection {
		t.log(0)
	}
}
