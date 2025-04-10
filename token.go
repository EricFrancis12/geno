package geno

type Token interface {
	FindString(string) (Token, string)
	String() string
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
