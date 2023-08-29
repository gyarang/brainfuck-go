package lexer

import "github.com/gyarang/brainfuck-go/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func NewLexer(input string) Lexer {
	l := Lexer{input: input}
	l.ReadChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token
	l.skipWhitespace()

	switch l.ch {
	case '>':
		t.Type = token.NEXT
	case '<':
		t.Type = token.PREV
	case '+':
		t.Type = token.INCR
	case '-':
		t.Type = token.DECR
	case '.':
		t.Type = token.PRINT
	case ',':
		t.Type = token.READ
	case '[':
		t.Type = token.LOOP_START
	case ']':
		t.Type = token.LOOP_END
	case 0:
		t.Type = token.EOF
	}

	l.ReadChar()
	return t
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.ReadChar()
	}
}

func (l *Lexer) ReadChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
