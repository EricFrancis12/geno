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

// This function extracts Directives from comments that use the following format:
// // #[foo, bar(baz)]
//
// The comment above would create 2 directives:
// 1. Directive{Name: "foo", Params: []string{}}
// 2. Directive{Name: "bar", Params: []string{"baz"}}
func (c CommentDirective) FindString(s string) (geno.Token, string) {
	directives := []Directive{}

	content := strings.TrimSpace(s)
	if !strings.HasPrefix(content, "//#") && !strings.HasPrefix(content, "// #") {
		return nil, ""
	}

	// Remove the comment prefix
	content = strings.TrimPrefix(content, "//")

	// Create a new parser to parse the remaining comment content: #[foo, bar(baz)]
	p := base.NewBaseParser(geno.SourceFile{Content: content})

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

	diff := len(s) - len(p.Remainder())
	took := s[:diff+1]

	return CommentDirective{
		Value:      took,
		Directives: directives,
	}, took
}

// TODO: combine this logic with BaseToken{}.Parse() (via tp.Parse() using StringFinder interface?)
func (c CommentDirective) Parse(tp geno.TokenParser) (geno.Token, error) {
	rem := tp.Remainder()

	tk, took := c.FindString(rem)
	if tk == nil {
		return nil, fmt.Errorf("cannot parse CommentDirective")
	}

	wip := ""

	for !tp.AtEOF() {
		wip += tp.Advance().String()

		if wip == took {
			return tk, nil
		} else if !strings.HasPrefix(took, wip) {
			return nil, fmt.Errorf(
				"expected '%s', to have prefix '%s'",
				took,
				wip,
			)
		}
	}

	return nil, fmt.Errorf("eof")
}

func (c CommentDirective) OnParse(ctx *geno.GenContext) {
	fmt.Println("TODO: CommentDirective{}.OnParse()")
}

func (c CommentDirective) String() string {
	return c.Value
}

type Directive struct {
	Name   string
	Params []string
}
