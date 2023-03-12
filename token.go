package main

type token struct {
	lex  lexeme
	data []byte
}

type tokenCollection []token
