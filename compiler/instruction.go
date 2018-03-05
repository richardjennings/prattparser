package compiler

import (
	"fmt"
	"github.com/richardjennings/pratt/token"
)

type (
	// ValType represents valid Value Types
	ValType int

	// Instructions struct contains a collection of Instructions and the count of Temporary Variables used.
	Instructions struct {
		tvals        int
		instructions []Instruction
	}

	// Instruction struct describes an individual Instruction
	Instruction struct {
		Op   Opcode
		Arg1 Argument
		Arg2 Argument
		Ret  Argument
	}

	// Argument struct describes an argument component of an Instruction
	Argument struct {
		Val     string
		ValType ValType
		TVal    int
	}
)

// Valid data types for argument values
const (
	NIL ValType = iota
	INTEGER
)

// ValTypes maps ValType constants to string representations
var ValTypes = [...]string{
	NIL:     "NIL",
	INTEGER: "INTEGER",
}

// NewInstruction maps token.Tokens to Opcodes
func NewInstruction(tok token.Token) Instruction {
	var Op Opcode
	switch tok {
	case token.ADD:
		Op = ADD
	case token.SUB:
		Op = SUBTRACT
	case token.MUL:
		Op = MULTIPLY
	case token.QUO:
		Op = DIVIDE
	case token.REM:
		Op = REMAINDER
	case token.POW:
		Op = POW
	default:
		panic("invalid token used as operator")
	}
	inst := Instruction{Op: Op}
	return inst
}

// Returns a string representation of Instructions
func (i Instructions) String() string {
	var str string
	str += fmt.Sprintf("tvals: %d\n", i.tvals)
	for i, v := range i.instructions {
		str += fmt.Sprintf("l%d: %s", i, v)
	}
	return str
}

// Returns a string representation of an Instruction
func (i Instruction) String() string {
	return fmt.Sprintf("Op: %s, Arg1: %s, Arg2: %s, Ret: %s\n", i.Op, i.Arg1, i.Arg2, i.Ret)
}

// Returns a string representation of an Argument
func (a Argument) String() string {
	if a.ValType == 0 && a.TVal == 0{
		return "{}"
	}
	if a.TVal > 0 {
		return fmt.Sprintf("{ TV%d }", a.TVal)
	}
	return fmt.Sprintf("{ Val: %s, %s }", a.Val, ValTypes[a.ValType])
}
