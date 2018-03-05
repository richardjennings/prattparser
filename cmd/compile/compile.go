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
	w io.Writer
}

// NewCli returns a new cli pointer using the supplied io.Writer implementation
func NewCli(writer io.Writer) *cli {
	return &cli{w: writer}
}

// A Command Line Interface to parse expressions
func main() {
	if len(os.Args) != 2 {
		fmt.Println(`example usage: ./compile "1 + 1"`)
		os.Exit(1)
	}
	cli := NewCli(os.Stdout)
	cli.exec(os.Args[1])
}

// exec executes the parser with supplied src, formats the result and writes to cli io.Writer
func (c *cli) exec(src string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(c.w, r)
		}
	}()
	p := parser.NewParser(src)

	expr := p.Parse()
	fmt.Fprintln(c.w, fmt.Sprintf("Expression: %s", expr))
	compiler := compiler.NewCompiler()
	instructions := compiler.Compile(expr)
	fmt.Fprintln(c.w, instructions.String())
}