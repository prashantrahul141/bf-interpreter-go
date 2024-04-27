package app

import (
	"bfigo/lexer"
	"bfigo/types"
	"bfigo/utils"
	"fmt"
	"os"
)

type App struct {
}

func (app *App) Run() {
	sourceContent := getFileContent()
	utils.GetGlobalLogger().Debug("File read", "content", sourceContent)

	var tokens []types.Token
	lexer := lexer.Lexer{tokens, sourceContent}
	lexer.ParseTokens()

}

func getFileContent() string {
	argPassed := os.Args[1:]
	if len(argPassed) <= 0 {
		utils.BfigoPanic(fmt.Sprint("Not enough arguments were passed.\n", utils.USAGE))
	}

	utils.GetGlobalLogger().Debug("", "filename", argPassed[0])

	fileContent, err := os.ReadFile(argPassed[0])
	if err != nil {
		utils.BfigoPanic(fmt.Sprintf("File '%s' does not exists.", argPassed[0]))
	}

	return string(fileContent)
}
