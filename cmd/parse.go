package main

import (
	"fmt"
	"os"
	"github.com/richardjennings/pratt/parser"
	"strings"
)

// A Command Line Interface to parse expressions
func main() {
	if len(os.Args) != 2 {
		fmt.Println(`example usage: ./parse "1 + 1"`)
		os.Exit(1)
	}
	p := parser.NewParser(os.Args[1])
	s := fmt.Sprintf("%s", p.Parse())
	s = strings.TrimPrefix(s,"(")
	s = strings.TrimSuffix(s,")")
	fmt.Println(s)
}
