package scanner

import (
	"bytes"
	"github.com/richardjennings/prattparser/token"
	"testing"
)

type testCase struct {
	Tok token.Token
	Lit string
}

func errorMsg(t *testing.T, expected testCase, actual Lexed) {
	t.Errorf("expected %v, got %v", expected, actual)
}

func TestInterpLexer_Lex_eof(t *testing.T) {
	s := Scanner{Src: bytes.NewBuffer([]byte(``))}
	if lexed := s.Lex(); lexed.Tok != token.EOF {
		t.Error(lexed)
	}
}

func TestInterpLexer_Lex_illegal(t *testing.T) {
	s := Scanner{Src: bytes.NewBuffer([]byte(`a`))}
	if lexed := s.Lex(); lexed.Tok != token.ILLEGAL || lexed.Lit != "a" {
		t.Error(lexed)
	}
}

func TestInterpLexer_Lex_int(t *testing.T) {
	s := Scanner{Src: bytes.NewBuffer([]byte(`0 10 1009 +1`))}
	expected := [...]testCase{
		{token.INT, "0"},
		{token.INT, "10"},
		{token.INT, "1009"},
		{token.ADD, ""},
		{token.INT, "1"},
	}
	for _, expect := range expected {
		lexed := s.Lex()
		if lexed.Tok != expect.Tok || lexed.Lit != expect.Lit {
			errorMsg(t, expect, lexed)
		}
	}
}

func TestInterpLexer_Lex_operator(t *testing.T) {
	s := Scanner{Src: bytes.NewBuffer([]byte(`+-*/%^()`))}
	expected := [...]testCase{
		{token.ADD, ""},
		{token.SUB, ""},
		{token.MUL, ""},
		{token.QUO, ""},
		{token.REM, ""},
		{token.POW, ""},
		{token.LPAREN, ""},
		{token.RPAREN, ""},
	}
	for _, expect := range expected {
		lexed := s.Lex()
		if lexed.Tok != expect.Tok || lexed.Lit != expect.Lit {
			errorMsg(t, expect, lexed)
		}
	}
}
