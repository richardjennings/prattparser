# Pratt Top-down Operator Precedence Example in Go
[![Go Report Card](https://goreportcard.com/badge/github.com/richardjennings/pratt)](https://goreportcard.com/report/github.com/richardjennings/prattparser) [![codecov](https://codecov.io/gh/richardjennings/prattparser/branch/master/graph/badge.svg)](https://codecov.io/gh/richardjennings/prattparser) [![Build Status](https://travis-ci.org/richardjennings/pratt.svg?branch=master)](https://travis-ci.org/richardjennings/prattparser)


## About

Pratt's Top-down Operator Precedence algorithm addresses a difficulty or complexity inherent in handling operator precedence in a Recursive Decent Parser. The algorithm works by associating semantics alongside tokens instead of / as well as grammar rules.

Please read [this blog post](http://effbot.org/zone/simple-top-down-parsing.htm) and [this blog post](https://dev.to/jrop/pratt-parsing) for an explanation of how the algorithm works.


## Lexer Scanner

The Scanner expects UTF8 input. The Lexer recognises integers and the following operators:

| symbol | token | operation |
|---|---|---|
| + | token.ADD | opcode.ADD |
| - | token.SUB | opcode.SUBTRACT |
| * | token.MUL | opcode.MULTIPLY |
| / | token.QUO | opcode.DIVIDE |
| % | token.REM | opcode.REMAINDER |
| ^ | token.POW | opcode.POW |
| ( | token.LPAREN | |
| ) | token.RPAREN | |


## Parser

The Parser consumes tokens provided by the Lexer and applies the Pratt
Operator Precedence Algorithm to generate an AST respecting operator precedence and associativity.

## Examples
```
$ go run cmd/parse/parse.go "1 + 2 * 3"
1 + (2 * 3)
```
```
$ go run cmd/parse/parse.go "(1 + 2) * 3"
(1 + 2) * 3
```
```
$ go run cmd/parse/parse.go "2 ^ 3 ^ 4"
2 ^ (3 ^ 4)
```

## Tests

To run tests:
```
$ go test -v ./...
```