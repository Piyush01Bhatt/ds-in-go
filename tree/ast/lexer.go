package ast

import (
	"errors"
	"unicode"
)

// Token types
const (
	NUMBER = "NUMBER"
	PLUS   = "PLUS"
	MINUS  = "MINUS"
	MUL    = "MUL"
	DIV    = "DIV"
	LPAREN = "LPAREN"
	RPAREN = "RPAREN"
)

type Token struct {
	Type  string
	Value string
}

type Lexer struct {
	input   string
	idx     int
	readIdx int
	ch      byte
}

func NewLexer(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar() // initialize first character
	return lex
}

func (lex *Lexer) readChar() {
	if lex.readIdx >= len(lex.input) {
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readIdx]
	}
	lex.idx = lex.readIdx
	lex.readIdx += 1
}

// func (lex *Lexer) skipWhitespace() {
// 	for unicode.IsSpace(rune(lex.ch)) {
// 		lex.readChar()
// 	}
// }

func (lex *Lexer) Tokenize() ([]Token, error) {
	var tokens []Token
	for lex.ch != 0 {
		switch {
		case unicode.IsSpace(rune(lex.ch)):
			lex.readChar()
		case unicode.IsDigit(rune(lex.ch)):
			start := lex.idx
			for unicode.IsDigit(rune(lex.ch)) {
				lex.readChar()
			}
			token := Token{Type: NUMBER, Value: lex.input[start:lex.idx]}
			tokens = append(tokens, token)
		case lex.ch == '+':
			tokens = append(tokens, Token{Type: PLUS, Value: "+"})
			lex.readChar()
		case lex.ch == '-':
			tokens = append(tokens, Token{Type: MINUS, Value: "-"})
			lex.readChar()
		case lex.ch == '/':
			tokens = append(tokens, Token{Type: DIV, Value: "/"})
			lex.readChar()
		case lex.ch == '*':
			tokens = append(tokens, Token{Type: MUL, Value: "*"})
			lex.readChar()
		case lex.ch == '(':
			tokens = append(tokens, Token{Type: LPAREN, Value: "("})
			lex.readChar()
		case lex.ch == ')':
			tokens = append(tokens, Token{Type: RPAREN, Value: ")"})
			lex.readChar()
		default:
			return nil, errors.New("invalid token")
		}
	}
	return tokens, nil
}
