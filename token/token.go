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

// Precedence returns a tokens precedence if defined, otherwise LowestPrec
func (tok Token) Precedence() int {
	switch tok {
	case ADD, SUB:
		return 4
	case MUL, QUO, REM:
		return 5
	case POW:
		return 6
	}
	return LowestPrec
}

// IsScalar returns true if the token is a scalar
func (tok Token) IsScalar() bool {
	switch tok {
	case INT:
		return true
	}
	return false
}

// IsUnary returns true if the token can be in a unary expression
func (tok Token) IsUnary() bool {
	switch tok {
	case INT, ADD, SUB:
		return true
	}
	return false
}

// IsLeftAssoc returns true if the token is left associative
func (tok Token) IsLeftAssoc() bool {
	switch tok {
	case INT, ADD, SUB, MUL, QUO, REM:
		return true
	}
	return false
}

// IsRightAssoc returns true if the token is right associative
func (tok Token) IsRightAssoc() bool {
	switch tok {
	case POW:
		return true
	}
	return false
}

// String provides a string representation of a Token
func (tok Token) String() string {
	if 0 <= tok && tok < Token(len(tokens)) {
		return tokens[tok]
	}
	return tokens[0]
}
