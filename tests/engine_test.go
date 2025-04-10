package tests

import (
	"testing"

	"github.com/EricFrancis12/geno"
	"github.com/EricFrancis12/geno/libs/base/custom"
	"github.com/EricFrancis12/geno/libs/directive"
)

func TestEngine(t *testing.T) {
	lib := custom.BaseTokenLib{}
	lib.AddToken(directive.CommentDirective{})

	e := geno.NewGenEngine(
		lib,
		directive.CommentDirective{},
	)

	sf := geno.SourceFile{
		Content: `
			// #[foo(bar), baz]
			enum Thing {
				ONE,
				TWO,
				THREE,
			}
		`,
	}

	e.Gen(sf)
}
