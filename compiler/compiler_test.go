package compiler

import (
	"github.com/richardjennings/pratt/ast"
	"github.com/richardjennings/pratt/token"
	"testing"
)

// compareInstructions
// cmp.Equal could be used but requires adding a dependency
func compareInstrucions(inst1 Instruction, inst2 Instruction) bool {
	return inst1.Op == inst2.Op &&
		inst1.Arg1.TVal == inst2.Arg1.TVal &&
		inst1.Arg1.Val == inst2.Arg1.Val &&
		inst1.Arg1.ValType == inst2.Arg1.ValType &&
		inst1.Arg2.TVal == inst2.Arg2.TVal &&
		inst1.Arg2.Val == inst2.Arg2.Val &&
		inst1.Arg2.ValType == inst2.Arg2.ValType &&
		inst1.Ret.TVal == inst2.Ret.TVal &&
		inst1.Ret.Val == inst2.Ret.Val &&
		inst1.Ret.ValType == inst2.Ret.ValType
}

// Compile a BinaryExpr and assert correct result
func TestCompiler_Compile_binaryExpr(t *testing.T) {
	expr := ast.BinaryExpr{
		X:  ast.ScalarExpr{Val: "1", Typ: token.INT},
		Op: token.ADD,
		Y: ast.BinaryExpr{
			X:  ast.ScalarExpr{Val: "2", Typ: token.INT},
			Op: token.MUL,
			Y:  ast.ScalarExpr{Val: "2", Typ: token.INT},
		},
	}
	c := NewCompiler()
	expected := []Instruction{
		{Op: MULTIPLY, Arg1: Argument{Val:"2", ValType:INTEGER}, Arg2: Argument{Val:"2", ValType:INTEGER}, Ret:Argument{TVal:1}},
		{Op: ADD, Arg1: Argument{Val:"1", ValType:INTEGER}, Arg2: Argument{TVal:1}, Ret:Argument{TVal:2}},
	}
	insts := c.Compile(expr)
	if insts.tvals != 2 {
		t.Errorf("expected 2 got %s", insts.tvals)
	}
	for i, actual := range(insts.instructions) {
		if compareInstrucions(expected[i], actual) == false {
			t.Errorf("expected %s got %s", expected, actual)
		}
	}
}

// Compile a UnaryExpr and assert correct result
func TestCompiler_Compile_unaryExpr(t *testing.T) {
	expr := ast.UnaryExpr{
		X:  ast.ScalarExpr{Val: "1", Typ: token.INT},
		Op: token.SUB,
	}
	c := NewCompiler()
	expected := []Instruction{
		{Op: SUBTRACT, Arg2: Argument{Val:"1", ValType:INTEGER}, Ret:Argument{TVal:1}},
	}
	insts := c.Compile(expr)
	if insts.tvals != 1 {
		t.Errorf("expected 2 got %s", insts.tvals)
	}
	for i, actual := range(insts.instructions) {
		if compareInstrucions(expected[i], actual) == false {
			t.Errorf("expected %s got %s", expected, actual)
		}
	}
}
