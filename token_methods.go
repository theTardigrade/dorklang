package main

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
