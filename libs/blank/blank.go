package blank

import "github.com/EricFrancis12/geno"

func NewEngine(triggers ...geno.GenTrigger) *geno.GenEngine[geno.Token] {
	return geno.NewGenEngine(BlankTokenLib{}, triggers...)
}

func NewParser(sourceFile geno.SourceFile) *geno.Parser[geno.Token] {
	return geno.NewParser(sourceFile, BlankTokenLib{})
}
