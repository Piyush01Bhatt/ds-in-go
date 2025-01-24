package main

import (
	"fmt"

	"github.com/Piyush01Bhatt/ds-in-go/tree/ast"
)

func main() {
	lex := ast.NewLexer("(3 + 5) * 2 - 8 / 4")
	tokens, err := lex.Tokenize()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tokens)
}
