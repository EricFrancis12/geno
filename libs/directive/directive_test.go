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
		{
			source: `// #[foo(bar), baz]`,
			expected: []Directive{
				{Name: "foo", Params: []string{"bar"}},
				{Name: "baz", Params: []string{}},
			},
		},
		{
			source: `// #[baz, foo(bar)]`,
			expected: []Directive{
				{Name: "baz", Params: []string{}},
				{Name: "foo", Params: []string{"bar"}},
			},
		},
		{
			source: `// #[foo(bar, baz)]`,
			expected: []Directive{
				{Name: "foo", Params: []string{"bar", "baz"}},
			},
		},
		{
			source: `// #[foo(bar, baz), qux]`,
			expected: []Directive{
				{Name: "foo", Params: []string{"bar", "baz"}},
				{Name: "qux", Params: []string{}},
			},
		},
		{
			source: `// #[foo(bar, baz), qux()]`,
			expected: []Directive{
				{Name: "foo", Params: []string{"bar", "baz"}},
				{Name: "qux", Params: []string{}},
			},
		},
		{
			source: `// #[foo(bar, baz), qux(quux)]`,
			expected: []Directive{
				{Name: "foo", Params: []string{"bar", "baz"}},
				{Name: "qux", Params: []string{"quux"}},
			},
		},
	}

	for _, test := range tests {
		ds, ok := ParseCommentDirectives(test.source)
		assert.True(t, ok)
		assert.Equal(t, test.expected, ds)
	}
}
