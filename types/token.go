package types

import "fmt"

// tokens enum parent.
type TokenType int

// all the possible types of token.
const (
	TokenRightAngle  TokenType = iota // >
	TokenLeftAngle                    // <
	TokenPlus                         // +
	TokenMinus                        // -
	TokenComma                        // ,
	TokenDot                          // .
	TokenRightSquare                  // ]
	TokenLeftSquare                   // [
	TokenEof                          // end of file.
)

// implement Fmt.Stringer for TokenType
func (ttype TokenType) String() string {
	switch ttype {
	case TokenRightAngle:
		return "TokenType(>)"
	case TokenLeftAngle:
		return "TokenType(<)"
	case TokenPlus:
		return "TokenType(+)"
	case TokenMinus:
		return "TokenType(-)"
	case TokenComma:
		return "TokenType(,)"
	case TokenDot:
		return "TokenType(.)"
	case TokenRightSquare:
		return "TokenType(])"
	case TokenLeftSquare:
		return "TokenType([)"
	case TokenEof:
		return "TokenType(EOF)"
	}
	return "TokenType(unkown)"
}

// Type of a specific token.
type Token struct {
	Line       uint32    // in the source file
	Token_type TokenType // its type.
}

// impl stringer for token.
func (t Token) String() string {
	return fmt.Sprintf("Token(%s, %d)", t.Token_type, t.Line)
}
