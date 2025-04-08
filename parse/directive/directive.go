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

	for p.CurrentBaseToken().Kind != geno.CLOSE_BRACKET && p.CurrentBaseToken().Kind != geno.EOF {
		btk := p.AdvanceBase()
		if btk.Kind != geno.IDENTIFIER {
			return []Directive{}, false
		}

		d := Directive{
			Name:   btk.Value,
			Params: []string{},
		}

		if p.CurrentBaseToken().Kind == geno.OPEN_PAREN {
			for p.AdvanceBase().Kind != geno.CLOSE_PAREN {
				if p.CurrentBaseToken().Kind == geno.IDENTIFIER {
					d.Params = append(d.Params, p.CurrentBaseToken().Value)
				} else if p.CurrentBaseToken().Kind == geno.COMMA {
					p.Advance()
				}
			}
		}

		directives = append(directives, d)

		if p.CurrentBaseToken().Kind == geno.COMMA {
			p.Advance()
		}
	}

	return directives, true
}
