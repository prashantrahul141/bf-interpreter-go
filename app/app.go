package app

import (
	"bfigo/lexer"
	"bfigo/types"
	"bfigo/utils"
)

type App struct {
}

func (app *App) Run() {
	sourceContent := utils.GetFileContent()
	utils.GetGlobalLogger().Debug("File read", "content", sourceContent)

	var tokens []types.Token
	lexer := lexer.Lexer{tokens, sourceContent}
	lexer.ParseTokens()

}
