package directive

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/EricFrancis12/geno"
)

func TestCommentDirectivesToken(t *testing.T) {
	type Test struct {
		source        string
		expectedToken geno.Token
		expectedTook  string
	}

	var tests = []Test{
		{
			source: `//#[]`,
			expectedToken: CommentDirective{
				directives: []Directive{},
				value:      `//#[]`,
			},
			expectedTook: `//#[]`,
		},
		{
			source: `// #[]`,
			expectedToken: CommentDirective{
				directives: []Directive{},
				value:      `// #[]`,
			},
			expectedTook: `// #[]`,
		},
		{
			source: `// #[foo]`,
			expectedToken: CommentDirective{
				directives: []Directive{
					{Name: "foo", Params: []string{}},
				},
				value: `// #[foo]`,
			},
			expectedTook: `// #[foo]`,
		},
		{
			source: `// #[foo()]`,
			expectedToken: CommentDirective{
				directives: []Directive{
					{Name: "foo", Params: []string{}},
				},
				value: `// #[foo()]`,
			},
			expectedTook: `// #[foo()]`,
		},
		{
			source: `// #[foo(bar)]`,
			expectedToken: CommentDirective{
				directives: []Directive{
					{Name: "foo", Params: []string{"bar"}},
				},
				value: `// #[foo(bar)]`,
			},
			expectedTook: `// #[foo(bar)]`,
		},
		{
			source: `// #[foo(bar), baz]`,
			expectedToken: CommentDirective{
				directives: []Directive{
					{Name: "foo", Params: []string{"bar"}},
					{Name: "baz", Params: []string{}},
				},
				value: `// #[foo(bar), baz]`,
			},
			expectedTook: `// #[foo(bar), baz]`,
		},
		{
			source: `// #[baz, foo(bar)]`,
			expectedToken: CommentDirective{
				directives: []Directive{
					{Name: "baz", Params: []string{}},
					{Name: "foo", Params: []string{"bar"}},
				},
				value: `// #[baz, foo(bar)]`,
			},
			expectedTook: `// #[baz, foo(bar)]`,
		},
		{
			source: `// #[foo(bar, baz)]`,
			expectedToken: CommentDirective{
				directives: []Directive{
					{Name: "foo", Params: []string{"bar", "baz"}},
				},
				value: `// #[foo(bar, baz)]`,
			},
			expectedTook: `// #[foo(bar, baz)]`,
		},
		{
			source: `// #[foo(bar, baz), qux]`,
			expectedToken: CommentDirective{
				directives: []Directive{
					{Name: "foo", Params: []string{"bar", "baz"}},
					{Name: "qux", Params: []string{}},
				},
				value: `// #[foo(bar, baz), qux]`,
			},
			expectedTook: `// #[foo(bar, baz), qux]`,
		},
		{
			source: `// #[foo(bar, baz), qux()]`,
			expectedToken: CommentDirective{
				directives: []Directive{
					{Name: "foo", Params: []string{"bar", "baz"}},
					{Name: "qux", Params: []string{}},
				},
				value: `// #[foo(bar, baz), qux()]`,
			},
			expectedTook: `// #[foo(bar, baz), qux()]`,
		},
		{
			source: `// #[foo(bar, baz), qux(quux)]`,
			expectedToken: CommentDirective{
				directives: []Directive{
					{Name: "foo", Params: []string{"bar", "baz"}},
					{Name: "qux", Params: []string{"quux"}},
				},
				value: `// #[foo(bar, baz), qux(quux)]`,
			},
			expectedTook: `// #[foo(bar, baz), qux(quux)]`,
		},
		{
			source: `// #[hello, foo(bar, baz), qux(quux)]`,
			expectedToken: CommentDirective{
				directives: []Directive{
					{Name: "hello", Params: []string{}},
					{Name: "foo", Params: []string{"bar", "baz"}},
					{Name: "qux", Params: []string{"quux"}},
				},
				value: `// #[hello, foo(bar, baz), qux(quux)]`,
			},
			expectedTook: `// #[hello, foo(bar, baz), qux(quux)]`,
		},
	}

	for _, test := range tests {
		tk, took := CommentDirective{}.FindString(test.source)
		assert.Equal(t, test.expectedToken, tk)
		assert.Equal(t, test.expectedTook, took)

		if tk != nil {
			cd, ok := tk.(CommentDirective)
			assert.True(t, ok)
			assert.Equal(t, test.expectedToken, cd)
		}
	}
}
