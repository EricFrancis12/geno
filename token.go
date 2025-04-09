package geno

type Token interface {
	FindString(string) string
	Parse(TokenParser) error
}

// This is a work-around type so Token can Accept a Parser[T]
// as an arg to .Parse(), otherwise it will be an
// invalid recursive loop
type TokenParser interface {
	AtEOF() bool
	Pos() int
	SetPos(int)
	CursorPos() int
	SeekTokenAt(cursorPos int)
	CurrentToken() Token
	Advance() Token
	AdvanceN(int) Token
}

type TokenWithCursorPos[T Token] struct {
	Token     T
	CursorPos int
}

func NewTokenWithCursorPos[T Token](t T, cursorPos int) TokenWithCursorPos[T] {
	return TokenWithCursorPos[T]{
		Token:     t,
		CursorPos: cursorPos,
	}
}
