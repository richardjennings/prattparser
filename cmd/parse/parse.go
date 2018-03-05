package main

import (
	"fmt"
	"github.com/richardjennings/pratt/parser"
	"io"
	"os"
	"strings"
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
	NewCli(os.Stdout).exec(os.Args)
}

// exec executes the parser with supplied src, formats the result and writes to cli io.Writer
func (c *cli) exec(args []string) {
	if len(args) != 2 {
		fmt.Println(`example usage: ./parse "1 + 1"`)
		os.Exit(1)
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(c.w, r)
		}
	}()
	p := parser.NewParser(args[1])

	s := fmt.Sprintf("%s", p.Parse())
	s = strings.TrimPrefix(s, "(")
	s = strings.TrimSuffix(s, ")")
	fmt.Fprintln(c.w, s)
}
