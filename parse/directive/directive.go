package directive

import (
	"strings"

	"github.com/EricFrancis12/geno"
)

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
	p := geno.NewBaseParser(geno.SourceFile{Content: s})

	if _, ok := p.AdvanceBaseTo(geno.HASHTAG, geno.OPEN_BRACKET); !ok {
		return directives, false
	}

	for p.CurrentBaseToken().Kind != geno.CLOSE_BRACKET {
		d := Directive{
			Name:   p.AdvanceBase().Value,
			Params: []string{},
		}

		if p.CurrentBaseToken().Kind == geno.OPEN_PAREN {
			p.Advance()

			for p.CurrentBaseToken().Kind != geno.CLOSE_PAREN {
				d.Params = append(d.Params, p.AdvanceBase().Value)

				if p.CurrentBaseToken().Kind == geno.COMMA {
					p.Advance()
				}
			}

			// Advance past the closing parenthesis
			p.Advance()
		}

		directives = append(directives, d)

		if p.CurrentBaseToken().Kind == geno.COMMA {
			p.Advance()
		}
	}

	return directives, true
}
