package app

import (
	"bfigo/lexer"
	"bfigo/parser"
	"bfigo/types"
	"bfigo/utils"
)

type App struct {
}

func (app *App) Run() {
	sourceContent := utils.GetFileContent()
	utils.GetGlobalLogger().Info("File read", "content", sourceContent)

	var tokens []types.Token
	m_lexer := lexer.Lexer{Tokens: tokens, Source: sourceContent}
	m_lexer.ParseTokens()

	var opCodes []int32
	m_parser := parser.Parser{OpCodes: opCodes, Lexer: &m_lexer}
	m_parser.Parse()
	utils.GetGlobalLogger().Info("done parsing", "opcode", m_parser.OpCodes)

}
