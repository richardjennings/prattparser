// Package ast provides AST representations and stringer methods
package ast

import (
	"fmt"
	"github.com/richardjennings/pratt/token"
)

type (
	// Expr interface for AST Expressions
	Expr interface {}

	// ScalarExpr represents Scalar Expressions such as a Integer
	ScalarExpr struct {
		Expr
		Val string
		Typ token.Token
	}

	// UnaryExpr represents Unary Expressions such as - 1
	UnaryExpr struct {
		Expr
		Op token.Token
		X  Expr
	}

	// BinaryExpr represents Binary Expression such as 1 + 1
	BinaryExpr struct {
		Expr
		X  Expr
		Op token.Token
		Y  Expr
	}
)

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
