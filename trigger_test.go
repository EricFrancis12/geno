package geno

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pass = false

type myGenTrigger struct{}

func (myGenTrigger) FindString(string) (Token, string) {
	return myGenTrigger{}, ""
}

func (myGenTrigger) Parse(TokenParser) (Token, error) {
	return myGenTrigger{}, nil
}

func (myGenTrigger) OnParse(*GenContext) {
	pass = true
}

func (myGenTrigger) String() string {
	return ""
}

func TestGenTrigger(t *testing.T) {
	m := myGenTrigger{}

	tk, err := m.Parse(nil)
	assert.Nil(t, err)

	op, ok := tk.(OnParse)
	assert.True(t, ok)

	op.OnParse(nil)
	assert.True(t, pass)
}
