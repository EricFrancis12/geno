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
				Directives: []Directive{},
				Value:      `//#[]`,
			},
			expectedTook: `//#[]`,
		},
		{
			source: `// #[]`,
			expectedToken: CommentDirective{
				Directives: []Directive{},
				Value:      `// #[]`,
			},
			expectedTook: `// #[]`,
		},
		{
			source: `// #[foo]`,
			expectedToken: CommentDirective{
				Directives: []Directive{
					{Name: "foo", Params: []string{}},
				},
				Value: `// #[foo]`,
			},
			expectedTook: `// #[foo]`,
		},
		{
			source: `// #[foo()]`,
			expectedToken: CommentDirective{
				Directives: []Directive{
					{Name: "foo", Params: []string{}},
				},
				Value: `// #[foo()]`,
			},
			expectedTook: `// #[foo()]`,
		},
		{
			source: `// #[foo(bar)]`,
			expectedToken: CommentDirective{
				Directives: []Directive{
					{Name: "foo", Params: []string{"bar"}},
				},
				Value: `// #[foo(bar)]`,
			},
			expectedTook: `// #[foo(bar)]`,
		},
		{
			source: `// #[foo(bar), baz]`,
			expectedToken: CommentDirective{
				Directives: []Directive{
					{Name: "foo", Params: []string{"bar"}},
					{Name: "baz", Params: []string{}},
				},
				Value: `// #[foo(bar), baz]`,
			},
			expectedTook: `// #[foo(bar), baz]`,
		},
		{
			source: `// #[baz, foo(bar)]`,
			expectedToken: CommentDirective{
				Directives: []Directive{
					{Name: "baz", Params: []string{}},
					{Name: "foo", Params: []string{"bar"}},
				},
				Value: `// #[baz, foo(bar)]`,
			},
			expectedTook: `// #[baz, foo(bar)]`,
		},
		{
			source: `// #[foo(bar, baz)]`,
			expectedToken: CommentDirective{
				Directives: []Directive{
					{Name: "foo", Params: []string{"bar", "baz"}},
				},
				Value: `// #[foo(bar, baz)]`,
			},
			expectedTook: `// #[foo(bar, baz)]`,
		},
		{
			source: `// #[foo(bar, baz), qux]`,
			expectedToken: CommentDirective{
				Directives: []Directive{
					{Name: "foo", Params: []string{"bar", "baz"}},
					{Name: "qux", Params: []string{}},
				},
				Value: `// #[foo(bar, baz), qux]`,
			},
			expectedTook: `// #[foo(bar, baz), qux]`,
		},
		{
			source: `// #[foo(bar, baz), qux()]`,
			expectedToken: CommentDirective{
				Directives: []Directive{
					{Name: "foo", Params: []string{"bar", "baz"}},
					{Name: "qux", Params: []string{}},
				},
				Value: `// #[foo(bar, baz), qux()]`,
			},
			expectedTook: `// #[foo(bar, baz), qux()]`,
		},
		{
			source: `// #[foo(bar, baz), qux(quux)]`,
			expectedToken: CommentDirective{
				Directives: []Directive{
					{Name: "foo", Params: []string{"bar", "baz"}},
					{Name: "qux", Params: []string{"quux"}},
				},
				Value: `// #[foo(bar, baz), qux(quux)]`,
			},
			expectedTook: `// #[foo(bar, baz), qux(quux)]`,
		},
		{
			source: `// #[hello, foo(bar, baz), qux(quux)]`,
			expectedToken: CommentDirective{
				Directives: []Directive{
					{Name: "hello", Params: []string{}},
					{Name: "foo", Params: []string{"bar", "baz"}},
					{Name: "qux", Params: []string{"quux"}},
				},
				Value: `// #[hello, foo(bar, baz), qux(quux)]`,
			},
			expectedTook: `// #[hello, foo(bar, baz), qux(quux)]`,
		},
	}

	for _, test := range tests {
		tk, took := CommentDirective{}.FindString(test.source)
		assert.Equal(t, test.expectedToken, tk)
		assert.Equal(t, test.expectedTook, took)

		if tk != nil {
			btk, ok := tk.(CommentDirective)
			assert.True(t, ok)
			assert.Equal(t, test.expectedToken, btk)
		}
	}
}
