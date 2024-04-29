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

