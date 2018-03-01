// Package token provides a representation of Lexed tokens
package token

// The Token type
type Token int

// Individual Token constants
const (
	ILLEGAL Token = iota
	EOF
	INT    // 12345
	ADD    // +
	SUB    // -
	MUL    // *
	QUO    // /
	REM    // %
	POW    // ^
	LPAREN // (
	RPAREN // )
)

// String representations of Token constants
var tokens = [...]string{
	ILLEGAL: "illegal",
	EOF:     "EOF",
	INT:     "int",
	ADD:     "+",
	SUB:     "-",
	MUL:     "*",
	QUO:     "/",
	REM:     "%",
	POW:     "^",
	LPAREN:  "(",
	RPAREN:  ")",
}

// Representation of highest and lowest precedence
const (
	LowestPrec  = 0
	HighestPrec = 7
)

// Precendence returns a tokens precedence if defined, otherwise LowestPrec
func (op Token) Precedence() int {
	switch op {
	case ADD, SUB:
		return 4
	case MUL, QUO, REM:
		return 5
	case POW:
		return 6
	}
	return LowestPrec
}

// String provides a string representation of a Token
func (tok Token) String() string {
	if 0 <= tok && tok < Token(len(tokens)) {
		return tokens[tok]
	} else {
		return tokens[0]
	}
}
