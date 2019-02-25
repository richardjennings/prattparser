# Pratt Top-down Operator Precedence Example in Go
[![Go Report Card](https://goreportcard.com/badge/github.com/richardjennings/pratt)](https://goreportcard.com/report/github.com/richardjennings/prattparser) [![codecov](https://codecov.io/gh/richardjennings/prattparser/branch/master/graph/badge.svg)](https://codecov.io/gh/richardjennings/prattparser) [![Build Status](https://travis-ci.org/richardjennings/pratt.svg?branch=master)](https://travis-ci.org/richardjennings/prattparser)


## About

The Pratt Top-down Operator Precedence algorithm addresses a difficulty or complexity inherent in handling operator precedence in a Recursive Decent Parser. 
The algorithm works by associating semantics alongside tokens instead of / as well as grammar rules that may be implemented in the parser.
The algorithm enables the handling of:
* infix and unary operators, e.g -10 (unary) 10 - 2 (infix) 
* right associativity, e.g. 1 ^ 2 ^ 3 => 1 ^ (2 ^ 3) not (1 ^ 2) ^ 3 
* operator precedence
* parenthesis to override precedence


## How does it work?
The scanner and token packages read a string such as "1 + 2 * 3" into a sequence of Tokens over which the parser can operate.
E.g. Token INT with literal value 1, Token ADD, Token INT literal value 2 and so on.
Each Token is associated with 2 behaviours and a weight. These are:
* Null Denotation (nud)
* Left Denotation (led)
* Left Binding Power value (lbp)

Null Denotation handles prefix contexts such as - 1. In this implementation, NUD creates either a scalar or unary ast node.  
Left Denotation handles infix contexts such as 2 - 1. In this implementation, LED creates a binary ast node.    
Left Binding Power is a weight which is used by the algorithm to determine the order by which operators in an expression should be evaluated.      

The core of the algorithm is:    
```
func expr(rbp = 0) {
	left = NUD(next())
	while rbp < LBP(peek()) {
        left = LED(left, next())
    }
    return left
}
```
The algorithm uses recursive calls to build a tree from the input stream of tokens. The right binding power is used to stop a
recursive call from parsing past tokens with a lower binding power than the operator currently being handled.

To illustrate; some examples:
```
//2 + 3 * 2
expr(0)   
NUD(2) => ast.scalar(INT, '2')
LBP(+) = 5 // greater than 0    
LED(left, next()) => LED(left, +) => ast.binary(ast.scalar(INT, '2'), +, expr(5))   
....expr(5)
....NUD(3) => ast.scalar(INT, '3')
....LBP(*) = 7 // greater than 5
....LED(left, next()) => LED(left, *) => ast.binary(ast.scalar(INT, '3'), *, expr(7))
........expr(7)
........NUD(2) => ast.scalar(INT, '2')
........LBP() = 0 // not greater than 7
........return ast.scalar(INT, '2')
....return ast.binary(ast.scalar(INT, '3'), *, ast.scalar(INT, '2'))
return ast.binary(ast.scalar(INT, '2'), +, ast.binary(ast.scalar(INT, '3'), *, ast.scalar(INT, '2')))

  + 
 / \
2   *
   / \
  3   2 
  
  
  
// 2 * 2 + 1
expr(0)
NUD(2) => ast.scalar(INT, '2')
LBP(*) = 7 // greater than 0 
LED(left, next()) => LED(left, *) => ast.binary(ast.scalar(INT, '2'), *, expr(7))
....expr(7)
....NUD(2) => ast.scalar(INT, '2')
....LBP(+) = 5 // less than 7
....return ast.scalar(INT, '2')
LBP(+) = 5 // more than 0
LED(left, next()) => LED(left, '+') => ast.binary(ast.binary(ast.scalar(INT, '2'), *, ast.scalar(INT, '2')), +, expr(5))
....expr(5)
....NUD(1) => ast.scalar(INT, '1')
....LBP() = 0 // less than 5
....return ast.scalar(INT, '1')
return ast.binary(ast.binary(ast.scalar(INT, '2'), *, ast.scalar(INT, '2')), +, ast.scalar(INT, '1'))
  
    +
   / \
  *   1
 / \
2   2  
```

Further reference articles: 
* [simple-top-down-parsing](http://effbot.org/zone/simple-top-down-parsing.htm)
* [pratt-parsing](https://dev.to/jrop/pratt-parsing) 


## Lexer Scanner

The Scanner expects UTF8 input. The Lexer recognises integers and the following operators:

| symbol | token |
|---|---|
| + | token.ADD |
| - | token.SUB |
| * | token.MUL |
| / | token.QUO |
| % | token.REM |
| ^ | token.POW |
| ( | token.LPAREN |
| ) | token.RPAREN |


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