package base

import "github.com/EricFrancis12/geno"

type BaseTokenLib struct{}

func (b BaseTokenLib) Tokenize(source string) []BaseToken {
	baseTokens := []BaseToken{}
	for _, twp := range b.TokenizeWithTrace(source) {
		baseTokens = append(baseTokens, twp.Token)
	}
	return baseTokens
}

func (b BaseTokenLib) TokenizeWithTrace(source string) []geno.TokenFromSource[BaseToken] {
	l := NewBaseLexer(source)

	for !l.AtEOF() {
		l.Match()
	}

	return l.TokensFromSource
}
