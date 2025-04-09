package base

import (
	"fmt"

	"github.com/EricFrancis12/geno"
)

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

func (t BaseToken) Parse(tp geno.TokenParser) (geno.Token, error) {
	s := tp.Advance().String()

	tk, took := t.FindString(s)
	if tk == nil {
		return nil, fmt.Errorf("cannot parse BaseToken from string (%s)", s)
	}

	if len(took) == len(s) {
		return tk, nil
	}

	return nil, fmt.Errorf(
		"partial match: expected to consume '%s', but only consumed '%s'",
		s,
		took,
	)
}

func (t BaseToken) String() string {
	return t.Value
}
