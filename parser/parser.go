package parser

import (
	"bfigo/lexer"
	"bfigo/types"
	"bfigo/utils"
	"fmt"
)

// interface to implement parser.
type IParser interface {
	// public method to start the parsing process.
	Parse()
	// Matches current token with the given token.
	matchToken(types.Token) bool
	// Returns if there are no more tokens to parse.
	isEmpty() bool
}

// Top level parser
// Holds a reference to lexer and a array of all opcodes parsed.
type Parser struct {
	OpCodes []int32
	Lexer   *lexer.Lexer
}








// Matches current token with the given token.
func (parser *Parser) matchToken(ttype types.TokenType) bool {
	if parser.isEmpty() {
		return false
	}
	return parser.Lexer.Peek().Token_type == ttype
}

// Returns if there are no more tokens to parse.
func (parser *Parser) isEmpty() bool {
	return len(parser.Lexer.Tokens) <= 0
}

