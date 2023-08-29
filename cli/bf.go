package main

import (
	"fmt"
	"github.com/gyarang/brainfuck-go/lexer"
	"github.com/gyarang/brainfuck-go/program"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s brainfuck filename\n", args[0])
		return
	}
	filename := args[1]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", filename, err)
		return
	}
	lex := lexer.NewLexer(string(content))
	prog := program.NewProgram(lex)
	prog.Compile()
	err = prog.Execute()
	if err != nil {
		fmt.Printf("Error occurred while executing %s: %v\n", filename, err)
		return
	}
	fmt.Println()
}
