package app

import (
	"bfigo/lexer"
	"bfigo/parser"
	"bfigo/types"
	"bfigo/utils"
	"bfigo/vm"
)

type App struct {
}

func (app *App) Run() {
	sourceContent := utils.GetFileContent()
	utils.GetGlobalLogger().Info("File read", "content", sourceContent)
	logger := utils.GetGlobalLogger()

	var tokens []types.Token
	m_lexer := lexer.Lexer{Tokens: tokens, Source: sourceContent, Logger: logger}
	m_lexer.ParseTokens()

	var opCodes []int32
	m_parser := parser.Parser{OpCodes: opCodes, Lexer: &m_lexer, Logger: logger}
	m_parser.Parse()
	utils.GetGlobalLogger().Info("done parsing", "opcode", m_parser.OpCodes)

	var vm_initial_state [utils.VM_STATE_SIZE]uint8
	m_vm := vm.Vm{Ip: 0, Dp: 0, State: vm_initial_state, OpCodes: m_parser.OpCodes, Logger: logger}
	m_vm.Execute()

}
