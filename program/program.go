package program

import (
	"bufio"
	"fmt"
	"github.com/gyarang/brainfuck-go/lexer"
	"github.com/gyarang/brainfuck-go/token"
	"os"
)

const size int = 65535

type Program struct {
	tok []token.Token
	lex lexer.Lexer
}

func NewProgram(lex lexer.Lexer) *Program {
	return &Program{
		lex: lex,
	}
}

func (p *Program) Compile() {
	for {
		t := p.lex.NextToken()
		p.tok = append(p.tok, t)
		if t.Type == token.EOF {
			break
		}
	}
}

func (p *Program) Execute() error {
	memory := make([]uint8, size)
	ptr := 0
	loopStart := 0

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(p.tok); i++ {
		switch p.tok[i].Type {
		case token.NEXT:
			if len(memory) == ptr+1 {
				memory = append(memory, 0)
			}
			ptr++
		case token.PREV:
			if ptr == 0 {
				return fmt.Errorf("pointer reached out of range")
			}
			ptr--
		case token.INCR:
			memory[ptr]++
		case token.DECR:
			memory[ptr]--
		case token.PRINT:
			fmt.Printf("%c", memory[ptr])
		case token.READ:
			read, _ := reader.ReadByte()
			memory[ptr] = uint8(read)
		case token.LOOP_START:
			if memory[ptr] == 0 {
				loopStart = 0
				for j := i + 1; j < len(memory); j++ {
					if p.tok[j].Type == token.LOOP_END {
						i = j + 1
						break
					}
				}
				return fmt.Errorf("loop end not found")
			} else {
				loopStart = i
			}
		case token.LOOP_END:
			if i == 0 {
				return fmt.Errorf("unexpected loop end")
			}
			if loopStart == 0 {
				return fmt.Errorf("unexpected loop end")
			}

			if memory[ptr] != 0 {
				i = loopStart
			} else {
				loopStart = 0
			}
		}
	}
	return nil
}
