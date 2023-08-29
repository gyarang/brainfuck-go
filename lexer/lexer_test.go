package lexer

import (
	"github.com/gyarang/brainfuck-go/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `><+-.,[]`

	tests := []struct {
		expectedType token.Type
	}{
		{token.NEXT},
		{token.PREV},
		{token.INCR},
		{token.DECR},
		{token.PRINT},
		{token.READ},
		{token.LOOP_START},
		{token.LOOP_END},
		{token.EOF},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type mismatch. Expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
	}
}
