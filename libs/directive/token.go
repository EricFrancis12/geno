package directive

import (
	"fmt"
	"strings"

	"github.com/EricFrancis12/geno"
	"github.com/EricFrancis12/geno/libs/base"
)

// CommentDirective satisfies the Token interface
type CommentDirective struct {
	Directives []Directive
	Value      string // The original string that the Directive was parsed from
}

func (c CommentDirective) FindString(s string) (geno.Token, string) {
	directives := []Directive{}
	value := s

	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "//#") && !strings.HasPrefix(s, "// #") {
		return nil, ""
	}

	// Remove the comment prefix
	s = strings.TrimPrefix(s, "//")

	// Create a new parser to parse the remaining comment content: #[foo, bar(baz)]
	p := base.NewBaseParser(geno.SourceFile{Content: s})

	if p.Advance().Kind != base.HASHTAG {
		return nil, ""
	}
	if p.Advance().Kind != base.OPEN_BRACKET {
		return nil, ""
	}

	for p.CurrentToken().Kind != base.CLOSE_BRACKET {
		d := Directive{
			Name:   p.Advance().Value,
			Params: []string{},
		}

		if p.CurrentToken().Kind == base.OPEN_PAREN {
			p.Advance()

			for p.CurrentToken().Kind != base.CLOSE_PAREN {
				d.Params = append(d.Params, p.Advance().Value)

				if p.CurrentToken().Kind == base.COMMA {
					p.Advance()
				}
			}

			// Advance past the closing parenthesis
			p.Advance()
		}

		directives = append(directives, d)

		if p.CurrentToken().Kind == base.COMMA {
			p.Advance()
		}
	}

	return CommentDirective{
		Value:      value,
		Directives: directives,
	}, value
}

// This function extracts Directives from comments that use the following format:
// // #[foo, bar(baz)]
//
// This comment would create 2 directives:
// 1. Directive{Name: "foo", Params: []string{}}
// 2. Directive{Name: "bar", Params: []string{"baz"}}
func (c CommentDirective) Parse(tp geno.TokenParser) (geno.Token, error) {
	s := tp.Advance().String()

	tk, took := c.FindString(s)
	if tk == nil {
		return nil, fmt.Errorf("cannot parse CommentDirective from string (%s)", s)
	}

	if len(took) != len(s) {
		return nil, fmt.Errorf(
			"partial match: expected to consume '%s', but only consumed '%s'",
			s,
			took,
		)
	}

	return tk, nil
}

func (c CommentDirective) OnParse(ctx *geno.GenContext) {

}

func (c CommentDirective) String() string {
	return c.Value
}

type Directive struct {
	Name   string
	Params []string
}
