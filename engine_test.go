package geno

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pass = false

type MyGenTrigger struct{}

func (m MyGenTrigger) FindString(string) (Token, string) {
	return MyGenTrigger{}, ""
}

func (m MyGenTrigger) Parse(TokenParser) (Token, error) {
	return MyGenTrigger{}, nil
}

func (m MyGenTrigger) OnParse(*GenContext[Token]) {
	pass = true
}

func TestEngine(t *testing.T) {
	m := MyGenTrigger{}

	tk, _ := m.Parse(nil)

	op, ok := tk.(OnParse[Token])
	assert.True(t, ok)

	op.OnParse(nil)
	assert.True(t, pass)
}
