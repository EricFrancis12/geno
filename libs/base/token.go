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

func (t BaseToken) WithPos(cursorPos int) geno.TokenWithCursorPos[*BaseToken] {
	return geno.NewTokenWithCursorPos(&t, cursorPos)
}

func (t *BaseToken) Parse(tp geno.TokenParser) error {
	p, ok := any(tp).(geno.Parser[*BaseToken])
	if !ok {
		return fmt.Errorf("cannot converting parser type")
	}

	// Advance the parser and absorb the token
	bt := p.Advance()
	t.Kind = bt.Kind
	t.Value = bt.Value

	return nil
}
