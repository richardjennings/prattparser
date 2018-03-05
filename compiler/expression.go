package compiler

import (
	"github.com/richardjennings/pratt/ast"
)

// CompileExpr invokes Expr type specific methods to return Expr collapsed as Instructions
func (c *Compiler) CompileExpr(node ast.Expr) []Instruction {
	var insts []Instruction
	switch node.(type) {
	case ast.UnaryExpr:
		node := node.(ast.UnaryExpr)
		return c.CompileUnaryExpr(node)
	case ast.BinaryExpr:
		node := node.(ast.BinaryExpr)
		return c.CompileBinaryExpr(node)
	}
	return insts
}

// CompileUnaryExpr compiles a Unary Expression to Instructions
func (c *Compiler) CompileUnaryExpr(node ast.UnaryExpr) []Instruction {
	inst := NewInstruction(node.Op)
	insts := c.CompileExprOperand(&inst.Arg2, node.X)
	inst.Ret = c.newTValArg()
	return append(insts, inst)
}

// CompileBiniaryExpr compiles a Binary Expression to Instructions
func (c *Compiler) CompileBinaryExpr(node ast.BinaryExpr) []Instruction {
	inst := NewInstruction(node.Op)
	insts := append(c.CompileExprOperand(&inst.Arg1, node.X), c.CompileExprOperand(&inst.Arg2, node.Y)...)
	inst.Ret = c.newTValArg()
	return append(insts, inst)
}

// CompileExprOperand recurses an X or Y operand of an Expr returning Instructions
func (c *Compiler) CompileExprOperand(arg *Argument, node ast.Expr) []Instruction {
	var pinsts []Instruction
	switch node.(type) {
	case ast.ScalarExpr:
		node := node.(ast.ScalarExpr)
		*arg = c.newLValArg(node.Val)
	default:
		pinsts = c.CompileExpr(node)
		*arg = pinsts[len(pinsts)-1].Ret
	}
	return pinsts
}
