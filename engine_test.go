package geno

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pass = false

type MyGenTrigger struct{}

func (MyGenTrigger) FindString(string) (Token, string) {
	return MyGenTrigger{}, ""
}

func (MyGenTrigger) Parse(TokenParser) (Token, error) {
	return MyGenTrigger{}, nil
}

func (MyGenTrigger) OnParse(*GenContext[Token]) {
	pass = true
}

func (MyGenTrigger) String() string {
	return ""
}

func TestEngine(t *testing.T) {
	m := MyGenTrigger{}

	tk, err := m.Parse(nil)
	assert.Nil(t, err)

	op, ok := tk.(OnParse[Token])
	assert.True(t, ok)

	op.OnParse(nil)
	assert.True(t, pass)
}
