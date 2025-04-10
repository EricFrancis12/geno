package directive

import (
	"github.com/EricFrancis12/geno"
	"github.com/EricFrancis12/geno/libs/base"
)

func Lib() base.CustomBaseTokenLib {
	lib := base.CustomBaseTokenLib{}
	lib.AddToken(CommentDirective{})
	return lib
}

func NewEngine(triggers ...geno.GenTrigger) *geno.GenEngine[geno.Token] {
	return geno.NewGenEngine(Lib(), triggers...)
}

func NewParser(sourceFile geno.SourceFile) *geno.Parser[geno.Token] {
	return geno.NewParser(sourceFile, Lib())
}
