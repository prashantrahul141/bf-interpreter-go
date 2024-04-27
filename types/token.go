package types

type TokenType int

const (
	TokenRightAngle  = iota // >
	TokenLeftAngle          // <
	TokenPlus               // +
	TokenMinus              // -
	TokenRightSquare        // ]
	TokenLeftSquare         // [
	TokenComma              // ,
	TokenDot                // .
	TokenEof                // end of file.
)

// Type of a specific token.
type Token struct {
	Line       uint32    // in the source file
	Token_type TokenType // its type.
}
