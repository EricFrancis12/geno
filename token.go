package geno

type Token interface {
	FindString(string) (Token, string)
	Parse(TokenParser) (Token, error)
	String() string
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
	GetSourceFile() SourceFile
	Remainder() string
}

type TokenFromSource[T Token] struct {
	Token     T
	CursorPos int
}

func NewTokenFromSource[T Token](t T, cursorPos int) TokenFromSource[T] {
	return TokenFromSource[T]{
		Token:     t,
		CursorPos: cursorPos,
	}
}

func (tfs TokenFromSource[T]) Generalize() TokenFromSource[Token] {
	tk, ok := any(tfs.Token).(Token)
	if !ok {
		panic("failed to assert type Token")
	}
	return TokenFromSource[Token]{
		Token:     tk,
		CursorPos: tfs.CursorPos,
	}
}
