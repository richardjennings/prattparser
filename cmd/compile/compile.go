package main

import (
	"fmt"
	"github.com/richardjennings/pratt/compiler"
	"github.com/richardjennings/pratt/parser"
	"io"
	"os"
)

// cli struct provides a configurable io.Writer for exec
type cli struct {
	o      io.Writer
	e      io.Writer
	exitFn func(int int)
}

// NewCli returns a new cli pointer using the supplied io.Writer implementations
func NewCli(out io.Writer, err io.Writer, exit func(int int)) *cli {
	return &cli{o: out, e: err, exitFn: exit}
}

// A Command Line Interface to parse expressions
func main() {
	NewCli(os.Stdout, os.Stderr, os.Exit).exec(os.Args)
}

// exec executes the parser with supplied src, formats the result and writes to cli io.Writer
func (c *cli) exec(args []string) {
	if len(args) != 2 {
		fmt.Fprintln(c.e, `example usage: ./compile "1 + 1"`)
		c.exitFn(1)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(c.e, r)
		}
	}()
	p := parser.NewParser(args[1])
	expr := p.Parse()
	fmt.Fprintln(c.o, fmt.Sprintf("Expression: %s", expr))
	compiler := compiler.NewCompiler()
	instructions := compiler.Compile(expr)
	fmt.Fprintln(c.o, instructions.String())
}
