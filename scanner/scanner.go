// Package scanner provides a UTF8 scanner and Lexer
package scanner

import (
	"bytes"
	"github.com/richardjennings/prattparser/token"
)

// The Scanner struct containing a bytes buffer and the
// next rune if not consumed by pass
type Scanner struct {
	Src *bytes.Buffer
	nch rune
}

// A Lexed struct contains a Lexed Token and Literal
type Lexed struct {
	Tok token.Token
	Lit string
}

// Scan returns the next rune or 0 on EOF
func (s *Scanner) Scan() rune {
	var ch rune
	var err error
	if s.nch != 0 {
		ch = s.nch
		s.nch = 0
	} else if ch, _, err = s.Src.ReadRune(); err != nil {
		return 0
	}

	return ch
}

// Lex skips whitespace and returns a Lexed struct containing a token.Token, string pair
func (s *Scanner) Lex() (lexed Lexed) {
	ch := s.Scan()
	// skip whitespace
	for ch == ' ' || ch == '\n' || ch == '\t' {
		ch = s.Scan()
	}
	switch true {
	case '0' <= ch && ch <= '9':
		// int
		buf := []rune{ch}
		ch = s.Scan()
		for ch >= '0' && ch <= '9' {
			buf = append(buf, ch)
			ch = s.Scan()
		}
		s.nch = ch
		lexed.Tok = token.INT
		lexed.Lit = string(buf)
	default:
		switch ch {
		case 0:
			lexed.Tok = token.EOF
		case '+':
			lexed.Tok = token.ADD
		case '-':
			lexed.Tok = token.SUB
		case '*':
			lexed.Tok = token.MUL
		case '/':
			lexed.Tok = token.QUO
		case '%':
			lexed.Tok = token.REM
		case '^':
			lexed.Tok = token.POW
		case '(':
			lexed.Tok = token.LPAREN
		case ')':
			lexed.Tok = token.RPAREN
		default:
			lexed.Tok = token.ILLEGAL
			lexed.Lit = string(ch)
		}
	}

	return
}
