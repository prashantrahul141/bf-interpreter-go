package parser

import (
	"bfigo/lexer"
	"bfigo/types"
	"bfigo/utils"
	"fmt"
)

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

// parses one statement at a time.
func (parser *Parser) parseOpCode() {
	// get the current token.
	current := parser.Lexer.Pop()

	switch current.Token_type {
	case types.TokenLeftSquare:
		parser.parseLeftSquare()

		// we called a right square token without a starting '[' block.
	case types.TokenRightSquare:
		utils.Error("Got ']' outside a loop block.", current.Line)

	// parses every other token.
	default:
		parser.parseNormalToken(current.Token_type)
	}

}

// parses '[' square
func (parser *Parser) parseLeftSquare() {
	var jumpLen int32 = 1 // this will keep count of number of instructions we need to jump back.
	// loop untill we reach a right square token.
	for !parser.matchToken(types.TokenRightSquare) && !parser.matchToken(types.TokenEof) {
		jumpLen++
		parser.parseNormalToken(parser.Lexer.Pop().Token_type)
	}

	parser.emitOpCode(types.MoveIPtr)
	parser.emitInt32(-jumpLen)

	fmt.Print(parser.Lexer.Peek())

	if parser.matchToken(types.TokenEof) {
		utils.Error("Non-terminating '['", parser.Lexer.Peek().Line)
	}

	// consume ending right square bracket.
	parser.Lexer.Pop()

}

// parses a norma token ( basically everything but '[' )
func (parser *Parser) parseNormalToken(tt types.TokenType) {
	switch tt {
	case types.TokenRightAngle:
		parser.emitOpCode(types.MovePtrForward)

	case types.TokenLeftAngle:
		parser.emitOpCode(types.MovePtrBackward)

	case types.TokenPlus:
		parser.emitOpCode(types.Increment)

	case types.TokenMinus:
		parser.emitOpCode(types.Decrement)

	case types.TokenComma:
		parser.emitOpCode(types.ReadFromStdin)

	case types.TokenDot:
		parser.emitOpCode(types.WriteToStdin)

	default:
		utils.BfigoPanic(fmt.Sprintf("Recieved a non-normal token in parseNormalToken : '%s'", tt))
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
