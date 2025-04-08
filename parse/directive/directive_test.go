package directive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommentDirectives(t *testing.T) {
	type Test struct {
		source   string
		expected []Directive
	}

	var tests = []Test{
		{
			source:   `//#[]`,
			expected: []Directive{},
		},
		{
			source:   `// #[]`,
			expected: []Directive{},
		},
		{
			source: `// #[foo]`,
			expected: []Directive{
				{Name: "foo", Params: []string{}},
			},
		},
		{
			source: `// #[foo()]`,
			expected: []Directive{
				{Name: "foo", Params: []string{}},
			},
		},
		{
			source: `// #[foo(bar)]`,
			expected: []Directive{
				{Name: "foo", Params: []string{"bar"}},
			},
		},
		// TODO: ...
		// {
		// 	source: `// #[foo(bar), baz]`,
		// 	expected: []Directive{
		// 		{Name: "foo", Params: []string{"bar"}},
		// 		{Name: "baz", Params: []string{}},
		// 	},
		// },
		// {
		// 	source: `// #[baz, foo(bar)]`,
		// 	expected: []Directive{
		// 		{Name: "baz", Params: []string{}},
		// 		{Name: "foo", Params: []string{"bar"}},
		// 	},
		// },
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, ParseCommentDirectives(test.source))
	}
}
