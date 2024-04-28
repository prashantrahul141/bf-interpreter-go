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
}

// Top level parser
// Holds a reference to lexer and a array of all opcodes parsed.
type Parser struct {
	OpCodes []int32
	Lexer   *lexer.Lexer
}

}
