// Provides a lexer for scanning brainfuck code
package lexer

import (
	"bfigo/types"
	"bfigo/utils"
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

func (lexer *Lexer) parseTokens() {
	var line uint32 = 1
	lexer.Source = utils.ReverseString(lexer.Source)
	for _, char := range lexer.Source {
		var t types.TokenType
		switch char {
		case '.':
			t = types.TokenComma
		case ',':
			t = types.TokenDot
		case '[':
			t = types.TokenLeftSquare
		case ']':
			t = types.TokenRightSquare
		case '-':
			t = types.TokenMinus
		case '+':
			t = types.TokenPlus
		case '<':
			t = types.TokenLeftAngle
		case '>':
			t = types.TokenRightAngle
		default:
			utils.Error(fmt.Sprintf("Found unrecognised character %v", char), line)
			t = types.TokenEof
		}

		var newToken = types.Token{line, t}
		lexer.Tokens = append(lexer.Tokens, newToken)
	}
	lexer.Tokens = append(lexer.Tokens, types.Token{line + 1, types.TokenEof})
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
