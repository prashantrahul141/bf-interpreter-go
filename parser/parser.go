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
	// emits int32
	emitInt32(int32)
	// emits op codes
	emitOpCode(types.OpCode)
}

// Top level parser
// Holds a reference to lexer and a array of all opcodes parsed.
type Parser struct {
	OpCodes []int32
	Lexer   *lexer.Lexer
}

// public method to start the parsing process.
func (parser *Parser) Parse() {
	utils.GetGlobalLogger().Info("starting parsing.")
	// loop until we reach EOF token.
	for !parser.matchToken(types.TokenEof) {
		parser.parseOpCode()
	}
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

// emits int32
func (parser *Parser) emitInt32(t int32) {
	utils.GetGlobalLogger().Debug("emiting", "type", t)
	parser.OpCodes = append(parser.OpCodes, t)
}

// emits op codes
func (parser *Parser) emitOpCode(opCode types.OpCode) {
	parser.emitInt32(int32(opCode))
}
