package vm

import "bfigo/utils"

// Vm struct.
type Vm struct {
	Ip      uint32                    // instruction pointer
	Dp      uint32                    // data pointer
	State   [utils.VM_DATA_SIZE]uint8 // Vm's internal state of the memory.
	OpCodes []int32                   // instructions to execute.
}
