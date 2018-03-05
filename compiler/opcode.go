package compiler

type Opcode int

// Opcodes
const (
	NOP Opcode = iota
	ASSIGN
	ADD
	SUBTRACT
	MULTIPLY
	DIVIDE
	POW
	REMAINDER
	ECHO
)

// Map of Opcodes to string representation
var Opcodes = [...]string{
	NOP:       "NOP",
	ASSIGN:    "ASSIGN",
	ADD:       "ADD",
	SUBTRACT:  "SUBTRACT",
	MULTIPLY:  "MULTIPLY",
	DIVIDE:    "DIVIDE",
	POW:       "POW",
	REMAINDER: "REMAINDER",
}

// Return an Opcode as a string
func (o Opcode) String() string {
	return Opcodes[o]
}
