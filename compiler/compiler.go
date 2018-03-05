package compiler

import (
	"github.com/richardjennings/pratt/ast"
)

// Compiler struct
type Compiler struct {
	// number of tvals
	cTval int
}

// NewCompiler creates a new Compiler returned by reference
func NewCompiler() *Compiler {
	return &Compiler{cTval: 0}
}

// newTValArg returns a new argument referring to the next Temporary Variable index
func (c *Compiler) newTValArg() Argument {
	c.cTval++
	return Argument{TVal: c.cTval}
}

// newLValArg returns a new argument with a Literal INTEGER value
func (c *Compiler) newLValArg(v string) Argument {
	return Argument{Val: v, ValType: INTEGER}
}

// Compile reduces a tree of Expressions to Instructions
func (c *Compiler) Compile(node ast.Expr) Instructions {
	inst := Instructions{}
	switch node.(type) {
	case ast.Expr:
		inst.instructions = c.CompileExpr(node)
		inst.tvals = c.cTval
	}
	return inst
}
