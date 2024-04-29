package vm

import (
	"bfigo/types"
	"bfigo/utils"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"golang.org/x/term"
)

// Vm struct.
type Vm struct {
	Ip      int32                      // instruction pointer
	Dp      int32                      // data pointer
	State   [utils.VM_STATE_SIZE]uint8 // Vm's internal state of the memory.
	OpCodes []int32                    // instructions to execute.

	Logger *log.Logger
}

func (vm *Vm) Execute() {
	vm.Logger.Info("starting execution ---------------------------------------------")
	for int(vm.Ip) != len(vm.OpCodes) {
		instruction := vm.OpCodes[vm.Ip]
		vm.executeInstruction(types.OpCode(instruction))
	}
}

func (vm *Vm) executeInstruction(instruction types.OpCode) {
	vm.Logger.Debug("executing", "instruction", instruction)
	switch instruction {
	case types.MoveDPtrForward:
		vm.execMoveDPtrFoward()

	case types.MoveDPtrBackward:
		vm.execMoveDPtrBackward()

	case types.ReadFromStdin:
		vm.execReadFromStdin()

	case types.WriteToStdin:
		vm.execWriteToStdin()

	case types.Increment:
		vm.execIncrement()

	case types.Decrement:
		vm.execDecrement()

	case types.MoveIPtr:
		vm.execMoveIPtr()
	}
}

func (vm *Vm) execMoveDPtrFoward() {
	vm.Ip++

	vm.Logger.Debug("move dptr forward")

	vm.Dp++
	if int(vm.Dp) >= utils.VM_STATE_SIZE {
		vm.Dp = 0
	}
}

func (vm *Vm) execMoveDPtrBackward() {
	vm.Ip++

	vm.Logger.Debug("move dptr backward")

	vm.Dp--
	if vm.Dp < 0 {
		vm.Dp = utils.VM_STATE_SIZE - 1
	}
}

func (vm *Vm) execReadFromStdin() {
	// switch stdin into 'raw' mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		utils.BfigoPanic("Couldn't read from stdin in raw mode.")
		return
	}

	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 1)
	_, err = os.Stdin.Read(b)
	if err != nil {
		utils.BfigoPanic("Couldn't read from stdin in raw mode.")
		return
	}

	vm.Logger.Debug("read from stdin", "char", b[0])

	vm.State[vm.Dp] = b[0]
	vm.Ip++
}

func (vm *Vm) execWriteToStdin() {
	c := vm.State[vm.Dp]
	vm.Logger.Debug("writing to stdin", "char", c)
	fmt.Printf("%c", c)
	vm.Ip++
}

func (vm *Vm) execIncrement() {
	vm.Logger.Debug("increment at", "DP", vm.Dp)
	vm.State[vm.Dp]++
	vm.Ip++
}

func (vm *Vm) execDecrement() {
	vm.Logger.Debug("decrement at", "DP", vm.Dp)
	vm.State[vm.Dp]--
	vm.Ip++
}

func (vm *Vm) execMoveIPtr() {
	if vm.State[vm.Dp] == 0 {
		vm.Ip += 2
		return
	}

	jumpLen := vm.OpCodes[vm.Ip+1]
	vm.Logger.Debug("jumping", "ip", jumpLen)
	vm.Ip = int32(jumpLen)
}
