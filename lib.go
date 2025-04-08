package geno

type TokenLib interface {
	Tokenize(string) []PositionedToken
}
