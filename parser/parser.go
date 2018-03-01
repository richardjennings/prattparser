// Package parser provides an implementation of Pratt Top-down Operator Precedence parsing
package parser

import (
	"bytes"
	"github.com/richardjennings/pratt/ast"
	"github.com/richardjennings/pratt/scanner"
	"github.com/richardjennings/pratt/token"
)

// The Parser struct
type Parser struct {
	Scanner scanner.Scanner
	lexed   scanner.Lexed
}

// NewParser creates a new parser
func NewParser(src string) *Parser {
	return &Parser{Scanner: scanner.Scanner{Src: bytes.NewBuffer([]byte(src))}}
}

// Parse Lexed tokens returning an Abstract Syntax Treee
func (p *Parser) Parse() ast.Expr {
	p.lexed = p.Scanner.Lex()
	return p.expr(token.LowestPrec)
}

// Implementation of Pratt Precedence
func (p *Parser) expr(rbp int) ast.Expr {
	t := p.lexed
	p.lexed = p.Scanner.Lex()
	//null denotation
	var left interface{}
	switch t.Tok {
	case token.LPAREN:
		left = p.expr(0)
		if p.lexed.Tok != token.RPAREN {
			panic("expected )")
		}
		p.lexed = p.Scanner.Lex()
	case token.INT:
		left = ast.ScalarExpr{Val: t.Lit, Typ: t.Tok}
	case token.ADD, token.SUB:
		left = ast.UnaryExpr{X: p.expr(token.HighestPrec), Op: t.Tok}
	}
	// left binding power
	for rbp < p.lexed.Tok.Precedence() {
		t = p.lexed
		p.lexed = p.Scanner.Lex()
		//left denotation
		switch t.Tok {
		case token.POW:
			left = ast.BinaryExpr{X: left, Op: t.Tok, Y: p.expr(t.Tok.Precedence() - 1)}
		case token.ADD, token.SUB, token.MUL, token.QUO, token.REM:
			left = ast.BinaryExpr{X: left, Op: t.Tok, Y: p.expr(t.Tok.Precedence())}
		}
	}
	return left
}
