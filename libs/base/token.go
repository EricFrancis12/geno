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

// TODO: is this correct?
func (t BaseToken) FindString(s string) (geno.Token, string) {
	l := NewBaseLexer(s)
	l.Match() // match once
	tk := l.PositionedTokens[0].Token
	return tk, tk.Value
}

// TODO: is this correct?
func (t BaseToken) Parse(tp geno.TokenParser) (geno.Token, error) {
	p, ok := any(tp).(geno.Parser[BaseToken])
	if !ok {
		return nil, fmt.Errorf("cannot converting parser type")
	}

	return p.Advance(), nil
}
