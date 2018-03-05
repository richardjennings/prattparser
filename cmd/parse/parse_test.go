package main

import (
	"bytes"
	"testing"
)


func TestExec_missing_argument(t *testing.T) {
	expected := `example usage: ./parse "1 + 1"
`
	bufo := bytes.NewBuffer([]byte{})
	bufe := bytes.NewBuffer([]byte{})
	cli := NewCli(bufo, bufe, func(int int){})
	cli.exec([]string{})
	actual := bufe.String()
	if actual != expected {
		t.Errorf("expected %v got %v", expected, actual)
	}
}


func TestExec_invalid_token(t *testing.T) {
	src := "1 + t"
	expected := "Parse Error: t\n"
	bufo := bytes.NewBuffer([]byte{})
	bufe := bytes.NewBuffer([]byte{})
	cli := NewCli(bufo, bufe, func(int int){})
	cli.exec([]string{"", src})
	actual := bufe.String()
	if actual != expected {
		t.Errorf("expected %v got %v", expected, actual)
	}
}

func TestExec_missing_rparen(t *testing.T) {
	src := "(1 + 2"
	expected := "Parse Error: expected )\n"
	bufo := bytes.NewBuffer([]byte{})
	bufe := bytes.NewBuffer([]byte{})
	cli := NewCli(bufo, bufe, func(int int){})
	cli.exec([]string{"", src})
	actual := bufe.String()
	if actual != expected {
		t.Errorf("expected %v got %v", expected, actual)
	}
}

func TestExec(t *testing.T) {
	src := "3 * 2 ^ 4 + (7 / 2) - 1"
	expected := "((3 * (2 ^ 4)) + (7 / 2)) - 1\n"
	bufo := bytes.NewBuffer([]byte{})
	bufe := bytes.NewBuffer([]byte{})
	cli := NewCli(bufo, bufe, func(int int){})
	cli.exec([]string{"", src})
	actual := bufo.String()
	if actual != expected {
		t.Errorf("expected %v got %v", expected, actual)
	}
}