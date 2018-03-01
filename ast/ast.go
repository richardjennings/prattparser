// Package ast provides AST representations and stringer methods
package ast

import (
	"fmt"
	"github.com/richardjennings/pratt/token"
)

// The general Expr interface
type Expr interface {
}

// A Scalar Expression represents Scalars such as a Integer
type ScalarExpr struct {
	Expr
	Val string
	Typ token.Token
}

// A Unary Expression such as - 1
type UnaryExpr struct {
	Expr
	Op token.Token
	X  Expr
}

// A Binary Expression such as 1 + 1
type BinaryExpr struct {
	Expr
	X  Expr
	Op token.Token
	Y  Expr
}

// String representation of Scalar Expression
func (s ScalarExpr) String() string {
	return fmt.Sprintf("%s", s.Val)
}

// String representation of Unary Expression
func (s UnaryExpr) String() string {
	return fmt.Sprintf("%s%s", s.Op, s.X)
}

// String representation of Binary Expression
func (s BinaryExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", s.X, s.Op, s.Y)
}
