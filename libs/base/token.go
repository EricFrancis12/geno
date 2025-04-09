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

func (t BaseToken) WithPos(cursorPos int) geno.TokenWithCursorPos[BaseToken] {
	return geno.NewTokenWithCursorPos(t, cursorPos)
}

func (t BaseToken) FindString(s string) (geno.Token, string) {
	l := NewBaseLexer(s)
	startLen := len(l.Remainder())

	l.Match() // match once

	endLen := len(l.Remainder())
	diff := startLen - endLen

	var tk geno.Token
	// Check if the lexer matched a token
	if len(l.PositionedTokens) > 0 {
		tk = l.PositionedTokens[0].Token
	}

	return tk, l.Source[:diff]
}

// TODO: is this correct?
func (t BaseToken) Parse(tp geno.TokenParser) (geno.Token, error) {
	p, ok := any(tp).(geno.Parser[BaseToken])
	if !ok {
		return nil, fmt.Errorf("cannot convert parser type")
	}

	return p.Advance(), nil
}
