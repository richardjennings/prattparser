# Pratt Top-down Operator Precedence Example in Go
[![Go Report Card](https://goreportcard.com/badge/github.com/richardjennings/pratt)](https://goreportcard.com/report/github.com/richardjennings/pratt) [![Coverage Status](https://coveralls.io/repos/github/richardjennings/pratt/badge.svg?branch=master)](https://coveralls.io/github/richardjennings/pratt?branch=master) [![Build Status](https://travis-ci.org/richardjennings/pratt.svg?branch=master)](https://travis-ci.org/richardjennings/pratt)


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

## Compiler

A compiler is included which reduces an ast.Expr to a set of Instructions that could be executed by a simple VM runtime.

Example:
```
$ go run cmd/compile/compile.go "4^7*2+1"
Expression: (((4 ^ 7) * 2) + 1)
tvals: 3
l0: Op: POW, Arg1: { Val: 4, INTEGER }, Arg2: { Val: 7, INTEGER }, Ret: { TV1 }
l1: Op: MULTIPLY, Arg1: { TV1 }, Arg2: { Val: 2, INTEGER }, Ret: { TV2 }
l2: Op: ADD, Arg1: { TV2 }, Arg2: { Val: 1, INTEGER }, Ret: { TV3 }
```

## Contributing

Please feel free to create Issues or Pull Requests identifying simplifications, conceptual inaccuracies, non idiomatic go usage, bugs, performance improvements, ...


## Tests

To run tests:
```
$ go test -v ./...
```