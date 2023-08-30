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
	cursor := 0
	loopStack := LoopStack{}

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < len(p.tok); i++ {
		switch p.tok[i].Type {
		case token.NEXT:
			if len(memory) == cursor+1 {
				memory = append(memory, 0)
			}
			cursor++
		case token.PREV:
			if cursor == 0 {
				return fmt.Errorf("pointer reached out of range")
			}
			cursor--
		case token.INCR:
			memory[cursor]++
		case token.DECR:
			memory[cursor]--
		case token.PRINT:
			fmt.Printf("%c", memory[cursor])
		case token.READ:
			read, _ := reader.ReadByte()
			memory[cursor] = uint8(read)
		case token.LOOP_START:
			if memory[cursor] != 0 {
				loopStack.Push(i)
				continue
			}

			// skip loop
			loopStartCnt := 0
			for j := i + 1; j < len(memory); j++ {
				if p.tok[j].Type == token.LOOP_START {
					loopStartCnt++
				}
				if p.tok[j].Type == token.LOOP_END {
					if loopStartCnt == 0 {
						i = j
						break
					}
					loopStartCnt--
				}
			}
		case token.LOOP_END:
			loopStart, err := loopStack.Pop()
			if err != nil {
				return fmt.Errorf("unexpected loop end")
			}

			if memory[cursor] != 0 {
				i = loopStart - 1
			}
		}
	}
	return nil
}
