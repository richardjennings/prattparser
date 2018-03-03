package main

import (
	"bytes"
	"testing"
)

func TestMain_invalid_src(t *testing.T) {
	src := "1 + t"
	expected := "Parse Error: t\n"
	buf := bytes.NewBuffer([]byte{})
	cli := NewCli(buf)
	cli.exec(src)
	actual := buf.String()
	if actual != expected {
		t.Errorf("expected %v got %v", expected, actual)
	}
}
