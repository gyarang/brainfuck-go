package token

type Type string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	NEXT = ">"
	PREV = "<"

	INCR = "+"
	DECR = "-"

	PRINT = "."
	READ  = ","

	LOOP_START = "["
	LOOP_END   = "]"
)

type Token struct {
	Type Type
}
