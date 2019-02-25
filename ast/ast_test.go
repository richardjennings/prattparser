package ast

import (
	"fmt"
	"github.com/richardjennings/prattparser/token"
	"testing"
)

func TestAstTreeStringer(t *testing.T) {
	expr := BinaryExpr{
		X: UnaryExpr{
			X:  ScalarExpr{Val: "1", Typ: token.INT},
			Op: token.SUB,
		},
		Op: token.MUL,
		Y: BinaryExpr{
			X: UnaryExpr{
				X:  ScalarExpr{Val: "2", Typ: token.INT},
				Op: token.ADD,
			},
			Op: token.MUL,
			Y: UnaryExpr{
				X:  ScalarExpr{Val: "2", Typ: token.INT},
				Op: token.SUB,
			},
		},
	}

	expected := "(-1 * (+2 * -2))"
	actual := fmt.Sprintf("%s", expr)

	if actual != expected {
		t.Errorf("expected %s got %s", expected, actual)
	}
}
