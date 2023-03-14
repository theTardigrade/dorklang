package dorklang

type token struct {
	lex             lexeme
	data            []byte          // used only when lex == filePathLexeme
	childCollection tokenCollection // used only when lex == parentLexeme
}

type tokenCollection []token
