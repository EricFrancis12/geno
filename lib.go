package geno

type TokenLib[T Token] interface {
	Tokenize(string) []T
	TokenizeWithTrace(string) []TokenFromSource[T]
}
