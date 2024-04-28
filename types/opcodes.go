package types

// Opcode enum parent.
type OpCode int32

// all the possible types opcodes supported by the vm.
const (
	MovePtrForward OpCode = iota
	MovePtrBackward
	ReadFromStdin
	WriteToStdin
	Increment
	Decrement
	MoveIPtr
)

// impl stringer for OpCode
func (op OpCode) String() string {
	switch op {
	case MovePtrForward:
		return "OpCode(MovePtrForward)"
	case MovePtrBackward:
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

	return "OpCode(unkown)"
}
