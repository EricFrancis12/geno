package directive

import (
	"strings"

	"github.com/EricFrancis12/geno"
	"github.com/EricFrancis12/geno/libs/base"
)

type CommentDirective struct {
	Directives []Directive
}

// TODO: ...
// func (c *CommentDirective) FindString(remainder string) string {}

// TODO: ...
// func (c *CommentDirective) Parse(tp geno.TokenParser) error {}

type Directive struct {
	Name   string
	Params []string
}

// This function extracts Directives from comments that use the following format:
// // #[foo, bar(baz)]
//
// This comment would create 2 directives:
// 1. Directive{Name: "foo", Params: []string{}}
// 2. Directive{Name: "bar", Params: []string{"baz"}}
func ParseCommentDirectives(s string) ([]Directive, bool) {
	directives := []Directive{}

	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "//#") && !strings.HasPrefix(s, "// #") {
		return directives, false
	}

	// Remove the comment prefix
	s = strings.TrimPrefix(s, "//")

	// Create a new parser to parse the remaining comment content: #[foo, bar(baz)]
	p := base.NewBaseParser(geno.SourceFile{Content: s})

	if p.Advance().Kind != base.HASHTAG {
		return directives, false
	}
	if p.Advance().Kind != base.OPEN_BRACKET {
		return directives, false
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

	return directives, true
}
