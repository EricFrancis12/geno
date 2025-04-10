package tests

import (
	"testing"

	"github.com/EricFrancis12/geno"
	"github.com/EricFrancis12/geno/libs/base/directive"
	"github.com/stretchr/testify/assert"
)

func TestDirectiveEngine(t *testing.T) {
	pass := false

	e := directive.NewEngine(
		directive.OnCommentDirective(
			func(ctx *geno.GenContext) {
				cg := geno.CodeGen{
					Code:       "hello",
					OutputPath: "./my/dir/file.txt",
				}
				ctx.WipCodeGen = append(ctx.WipCodeGen, cg)
				pass = true
			},
		),
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

	cgs := e.Gen(sf)

	assert.Equal(
		t,
		[]geno.CodeGen{{Code: "hello", OutputPath: "./my/dir/file.txt"}},
		cgs,
	)
	assert.True(t, pass)
}
