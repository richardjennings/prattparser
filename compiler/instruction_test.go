package compiler

import (
	"testing"
	"github.com/richardjennings/pratt/token"
)

func TestNewInstruction(t *testing.T) {
	expected := []Opcode{
		ADD,SUBTRACT,MULTIPLY,DIVIDE,REMAINDER,POW,
	}
	for i, tok := range[]token.Token{
		token.ADD, token.SUB, token.MUL, token.QUO, token.REM, token.POW,
	}{
		inst := NewInstruction(tok)
		if inst.Op != expected[i] {
			t.Errorf("expected %s got %s", expected[i], inst.Op)
		}
	}
}

func TestInstructions_String(t *testing.T) {
	insts := Instructions{instructions:[]Instruction{{}}}
	expected := "tvals: 0\nl0: Op: NOP, Arg1: {}, Arg2: {}, Ret: {}\n"
	actual := insts.String()
	if actual != expected {
		t.Errorf("expected %s got %s", expected, actual)
	}
}

func TestArgument_String(t *testing.T) {
	expected := []string{"{}", "{ Val: test, INTEGER }", "{ TV1 }"}
	for i, tc := range []Argument{
		{},
		{Val: "test", ValType:INTEGER},
		{TVal:1},
	} {
		actual := tc.String()
		if actual != expected[i] {
			t.Errorf("expected %s got %s", expected[i], actual)
		}
	}


}