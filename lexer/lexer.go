// Provides a lexer for scanning brainfuck code
package lexer

import (
	"bfigo/types"
	"bfigo/utils"

	"github.com/charmbracelet/log"
)

// Top level lexer
// implements `ILexer`
type Lexer struct {
	Tokens []types.Token
	Source string
	Logger *log.Logger
}

// Scans tokens and stores them in the
// lexer's tokens slice.
func (lexer *Lexer) ParseTokens() {
	lexer.Logger.Info("start parsing tokens.")
	var line uint32 = 1

	for _, char := range lexer.Source {
		if char == utils.EOF_ASCII_CODE {
			lexer.Logger.Info("breaking parsing tokens")
			break
		}
		lexer.Logger.Debug("current", "char", char)

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
		default:
			// do nothing, dont need break because we are using hecking go.
		}
	}

	lexer.Tokens = append(lexer.Tokens, types.Token{Line: line + 1, Token_type: types.TokenEof})

	lexer.Logger.Debug("reversing array.")
	// reverse array because we will be using peek and pop to retrive tokens.
	utils.ReverseArray(lexer.Tokens)

	lexer.Logger.Info("parsed tokens (in reverse order): ")
	for _, v := range lexer.Tokens {
		lexer.Logger.Info(v)
	}
}

// Returns next token in the stream without consuming it.
// Returns TokenEOF if there are no tokens left
// in stream.
func (lexer *Lexer) Peek() types.Token {
	var peekedToken types.Token
	if len(lexer.Tokens) > 0 {
		peekedToken = lexer.Tokens[len(lexer.Tokens)-1]
	} else {
		peekedToken = types.Token{Line: 0, Token_type: types.TokenEof}
	}
	// lexer.Logger.Debug("peeked", "token", peekedToken)
	return peekedToken
}

// Returns next token in the stream AND consumes it
// Returns TokenEOF if there are no tokens left in stream.
func (lexer *Lexer) Pop() types.Token {
	popedToken := types.Token{Line: 0, Token_type: types.TokenEof}
	if len(lexer.Tokens) > 0 {
		popedToken = lexer.Tokens[len(lexer.Tokens)-1]
		lexer.Tokens = lexer.Tokens[:len(lexer.Tokens)-1]
	}

	// lexer.Logger.Debug("poped", "token", popedToken)
	return popedToken
}

// Private function to add token at given line number of given type.
func (lexer *Lexer) addToken(line uint32, token_type types.TokenType) {
	var newToken = types.Token{Line: line, Token_type: token_type}
	lexer.Logger.Debug("add", "token", newToken)
	lexer.Tokens = append(lexer.Tokens, newToken)
}
