
# Pratt Top-down Operator Precedence Example in Go
[![Go Report Card](https://goreportcard.com/badge/github.com/richardjennings/pratt)](https://goreportcard.com/report/github.com/richardjennings/pratt) [![Coverage Status](https://coveralls.io/repos/github/richardjennings/pratt/badge.svg?branch=master)](https://coveralls.io/github/richardjennings/pratt?branch=master)

## Examples
```
$ go run cmd/parse.go "1 + 2 * 3"   
1 + (2 * 3)    
```
```
$ go run cmd/parse.go "(1 + 2) * 3"   
(1 + 2) * 3    
```
```
$ go run cmd/parse.go "2 ^ 3 ^ 4"   
2 ^ (3 ^ 4)    
```

## About
Pratt's Top-down Operator Precedence algorithm addresses a difficulty or complexity inherent in handling operator precedence in a Recursive Decent Parser. The algorithm works by associating semantics alongside tokens instead of / as well as grammar rules. I found [this blog post](http://effbot.org/zone/simple-top-down-parsing.htm) helpful in understanding how the algorithm works.

This project implements a simple Scanner / Lexer which recognises integers, '+', '-', '*', '/', '%', '^', '(' & ')'.

The parser implementing Pratt Top-down Operator Precedence consumes the Lexer and produces an AST returned as a formated string.

The aim of this project is to make writing a Recursive Decent Parser for a more featured language easier.
