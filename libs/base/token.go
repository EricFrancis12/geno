package base

import "github.com/EricFrancis12/geno"

type BaseToken struct {
	Kind  BaseTokenKind
	Value string
}

func NewBaseToken(kind BaseTokenKind, value string) BaseToken {
	return BaseToken{
		Kind:  kind,
		Value: value,
	}
}

func (t BaseToken) WithPos(cursorPos int) geno.TokenFromSource[BaseToken] {
	return geno.NewTokenFromSource(t, cursorPos)
}

func (t BaseToken) FindString(s string) (geno.Token, string) {
	l := NewBaseLexer(s)

	startLen := len(l.Remainder())
	l.Match() // match once
	endLen := len(l.Remainder())

	diff := startLen - endLen
	took := l.Source[:diff]

	var tk geno.Token
	// Check if the lexer matched a token
	if len(l.TokensFromSource) > 0 {
		tk = l.TokensFromSource[0].Token
	}

	return tk, took
}

func (t BaseToken) String() string {
	return t.Value
}
