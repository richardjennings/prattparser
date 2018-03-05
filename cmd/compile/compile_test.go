package main

import (
	"bytes"
	"testing"
)

func TestExec_exec(t *testing.T) {
	src := "-4^7*2+1"
	buf := bytes.NewBuffer([]byte{})
	cli := NewCli(buf)
	cli.exec([]string{"", src})
	actual := buf.String()
	expected := `Expression: (((-4 ^ 7) * 2) + 1)
tvals: 4
l0: Op: SUBTRACT, Arg1: {}, Arg2: { Val: 4, INTEGER }, Ret: { TV1 }
l1: Op: POW, Arg1: { TV1 }, Arg2: { Val: 7, INTEGER }, Ret: { TV2 }
l2: Op: MULTIPLY, Arg1: { TV2 }, Arg2: { Val: 2, INTEGER }, Ret: { TV3 }
l3: Op: ADD, Arg1: { TV3 }, Arg2: { Val: 1, INTEGER }, Ret: { TV4 }

`
	if actual != expected {
		t.Errorf("\nexpected\n%q \ngot\n%q", expected, actual)
	}
}
