package types

import "fmt"

// Opcode enum parent.
type OpCode int32

// all the possible types opcodes supported by the vm.
const (
	MoveDPtrForward OpCode = iota
	MoveDPtrBackward
	ReadFromStdin
	WriteToStdin
	Increment
	Decrement
	MoveIPtr
)

// impl stringer for OpCode
func (op OpCode) String() string {
	switch op {
	case MoveDPtrForward:
		return "OpCode(MovePtrForward)"
	case MoveDPtrBackward:
		return "OpCode(MovePtrBackward)"
	case ReadFromStdin:
		return "OpCode(ReadFromStdin)"
	case WriteToStdin:
		return "OpCode(WriteToStdin)"
	case Increment:
		return "OpCode(Increment)"
	case Decrement:
		return "OpCode(Decrement)"
	case MoveIPtr:
		return "OpCode(MoveIPtr)"
	}

	return fmt.Sprintf("OpCode(%d)", int(op))
}
