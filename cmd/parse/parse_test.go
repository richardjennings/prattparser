package main

import (
	"bytes"
	"testing"
)

func TestExec_invalid_token(t *testing.T) {
	src := "1 + t"
	expected := "Parse Error: t\n"
	buf := bytes.NewBuffer([]byte{})
	cli := NewCli(buf)
	cli.exec([]string{"", src})
	actual := buf.String()
	if actual != expected {
		t.Errorf("expected %v got %v", expected, actual)
	}
}

func TestExec_missing_rparen(t *testing.T) {
	src := "(1 + 2"
	expected := "Parse Error: expected )\n"
	buf := bytes.NewBuffer([]byte{})
	cli := NewCli(buf)
	cli.exec([]string{"", src})
	actual := buf.String()
	if actual != expected {
		t.Errorf("expected %v got %v", expected, actual)
	}
}
