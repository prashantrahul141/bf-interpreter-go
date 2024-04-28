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
	m_lexer := lexer.Lexer{tokens, sourceContent}
	m_lexer.ParseTokens()

	var opCodes []int32
	parser := parser.Parser{opCodes, &m_lexer}
	parser.Parse()
	utils.GetGlobalLogger().Info("done parsing", "opcode", parser.OpCodes)

}
