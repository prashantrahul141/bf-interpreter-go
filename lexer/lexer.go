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

