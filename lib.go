package geno

type TokenLib[T Token] interface {
	Tokenize(string) []T
	TokenizeWithPos(string) []TokenWithCursorPos[T]
}
