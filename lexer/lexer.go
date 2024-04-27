// Provides a lexer for scanning brainfuck code
package lexer

import (
	"fmt"
)
// interface to implement lexer.
type ILexer interface {
	// Scans tokens and stores them in the
	// lexer's tokens slice.
	parseTokens()

	// Returns next token in the stream AND consumes it
	// Returns TokenEOF if there are no tokens left in stream.
	pop()

	// Returns next token in the stream without consuming it.
	// Returns TokenEOF if there are no tokens left
	// in stream.
	peek()
}

// Top level lexer
// implements `ILexer`
type Lexer struct {
	Tokens []types.Token
	Source string
}

func (lexer *Lexer) peek() types.Token {
	if len(lexer.Tokens) > 0 {
		return lexer.Tokens[len(lexer.Tokens)-1]
	}
	return types.Token{0, types.TokenEof}
}

func (lexer *Lexer) pop() types.Token {
	if len(lexer.Tokens) > 0 {
		last := lexer.Tokens[len(lexer.Tokens)-1]
		lexer.Tokens = lexer.Tokens[:len(lexer.Tokens)-1]
		return last
	}
	return types.Token{0, types.TokenEof}
}
