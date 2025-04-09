package base

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/EricFrancis12/geno"
)

func TestBaseToken(t *testing.T) {
	type Test struct {
		source        string
		expectedToken geno.Token
		expectedTook  string
	}

	var tests = []Test{
		// Var
		{
			source: "var i = 4;",
			expectedToken: BaseToken{
				Kind:  VAR,
				Value: "var",
			},
			expectedTook: "var",
		},
		// Identifier
		{
			source: "i = 4;",
			expectedToken: BaseToken{
				Kind:  IDENTIFIER,
				Value: "i",
			},
			expectedTook: "i",
		},
		// Number
		{
			source: "4;",
			expectedToken: BaseToken{
				Kind:  NUMBER,
				Value: "4",
			},
			expectedTook: "4",
		},
		// String
		{
			source: `"my first string" + " my second string"`,
			expectedToken: BaseToken{
				Kind:  STRING,
				Value: `"my first string"`,
			},
			expectedTook: `"my first string"`,
		},
		// Comment
		{
			source:        "// This is my comment",
			expectedToken: nil,
			expectedTook:  "// This is my comment",
		},
	}

	for _, test := range tests {
		tk, took := BaseToken{}.FindString(test.source)
		assert.Equal(t, test.expectedToken, tk)
		assert.Equal(t, test.expectedTook, took)

		if tk != nil {
			btk, ok := tk.(BaseToken)
			assert.True(t, ok)
			assert.Equal(t, test.expectedToken, btk)
		}
	}
}
