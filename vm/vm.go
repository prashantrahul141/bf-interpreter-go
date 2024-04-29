package vm

import "bfigo/utils"
import (
	"bfigo/types"
	"bfigo/utils"
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

// Vm struct.
type Vm struct {
	Ip      int32                      // instruction pointer
	Dp      int32                      // data pointer
	State   [utils.VM_STATE_SIZE]uint8 // Vm's internal state of the memory.
	OpCodes []int32                    // instructions to execute.
}


func (vm *Vm) execMoveDPtrFoward() {
	vm.Ip++

	vm.Dp++
	if int(vm.Dp) >= utils.VM_STATE_SIZE {
		vm.Dp = 0
	}
}

func (vm *Vm) execMoveDPtrBackward() {
	vm.Ip++

	vm.Dp--
	if vm.Dp < 0 {
		vm.Dp = utils.VM_STATE_SIZE - 1
	}
}

func (vm *Vm) execReadFromStdin() {
	// switch stdin into 'raw' mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 1)
	_, err = os.Stdin.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	vm.Ip++
}

func (vm *Vm) execWriteToStdin() {
	data := vm.State[vm.Dp]
	fmt.Printf("%c", data)
	vm.Ip++
}

func (vm *Vm) execIncrement() {
	vm.State[vm.Dp]++
	vm.Ip++
}

func (vm *Vm) execDecrement() {
	vm.State[vm.Dp]--
	vm.Ip++
}

func (vm *Vm) execMoveIPtr() {
	utils.GetGlobalLogger().Debug(vm.State)
	if vm.State[vm.Dp] == 0 {
		vm.Ip += 2
		return
	}

	jumpLen := vm.OpCodes[vm.Ip+1]
	utils.GetGlobalLogger().Info("jumping", "len", jumpLen)
	vm.Ip = int32(jumpLen)
}
