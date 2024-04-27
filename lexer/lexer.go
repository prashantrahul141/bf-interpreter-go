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
	ParseTokens()

	// Returns next token in the stream AND consumes it
	// Returns TokenEOF if there are no tokens left in stream.
	Pop()

	// Returns next token in the stream without consuming it.
	// Returns TokenEOF if there are no tokens left
	// in stream.
	Peek()

	// Private function to add token at given line number of given type.
	addToken(line uint32, token_type types.TokenType)
}

// Top level lexer
// implements `ILexer`
type Lexer struct {
	Tokens []types.Token
	Source string
}

func (lexer *Lexer) ParseTokens() {
	var line uint32 = 1

	for _, char := range lexer.Source {
		if char == 10 {
			break
		}

		switch char {
		case '.':
			lexer.addToken(line, types.TokenDot)
		case ',':
			lexer.addToken(line, types.TokenComma)
		case '[':
			lexer.addToken(line, types.TokenLeftSquare)
		case ']':
			lexer.addToken(line, types.TokenRightSquare)
		case '-':
			lexer.addToken(line, types.TokenMinus)
		case '+':
			lexer.addToken(line, types.TokenPlus)
		case '<':
			lexer.addToken(line, types.TokenLeftAngle)
		case '>':
			lexer.addToken(line, types.TokenRightAngle)
		case '\n':
			line++
		case ' ' | '\t':
		// do nothing, dont need break because we are using hecking go.

		default:
			utils.Error(fmt.Sprintf("Found unrecognised character: '%v'", char), line)
		}
	}

	lexer.Tokens = append(lexer.Tokens, types.Token{line + 1, types.TokenEof})

	// reverse array because we will be using peek and pop to retrive tokens.
	utils.ReverseArray(lexer.Tokens)
}

func (lexer *Lexer) Peek() types.Token {
	if len(lexer.Tokens) > 0 {
		return lexer.Tokens[len(lexer.Tokens)-1]
	}
	return types.Token{0, types.TokenEof}
}

func (lexer *Lexer) Pop() types.Token {
	if len(lexer.Tokens) > 0 {
		last := lexer.Tokens[len(lexer.Tokens)-1]
		lexer.Tokens = lexer.Tokens[:len(lexer.Tokens)-1]
		return last
	}
	return types.Token{0, types.TokenEof}
}

func (lexer *Lexer) addToken(line uint32, token_type types.TokenType) {
	var newToken = types.Token{line, token_type}
	lexer.Tokens = append(lexer.Tokens, newToken)
}
