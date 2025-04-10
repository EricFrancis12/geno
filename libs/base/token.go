package base

import (
	"fmt"
	"strings"

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

// Parse determines if a single BaseToken can be parsed
// by consuming tokens from tp, and if so returns it.
// We don't know what tokens tp uses, so we are using
// FindString() to parse the remainder into BaseTokens.
func (t BaseToken) Parse(tp geno.TokenParser) (geno.Token, error) {
	rem := tp.Remainder()

	tk, took := t.FindString(rem)
	if tk == nil {
		return nil, fmt.Errorf("cannot parse BaseToken")
	}

	wip := ""

	for !tp.AtEOF() {
		wip += tp.Advance().String()

		if wip == took {
			return tk, nil
		} else if !strings.HasPrefix(took, wip) {
			return nil, fmt.Errorf(
				"expected '%s', to have prefix '%s'",
				took,
				wip,
			)
		}
	}

	return nil, fmt.Errorf("eof")
}

func (t BaseToken) String() string {
	return t.Value
}
