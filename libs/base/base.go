package base

import "github.com/EricFrancis12/geno"

func NewBaseGenEngine(triggers ...geno.GenTrigger) *geno.GenEngine[BaseToken] {
	return &geno.GenEngine[BaseToken]{
		TokenLib: BaseTokenLib{},
		Triggers: triggers,
	}
}

func NewBaseParser(sourceFile geno.SourceFile) *geno.Parser[BaseToken] {
	return &geno.Parser[BaseToken]{
		SourceFile:       sourceFile,
		TokensFromSource: BaseTokenLib{}.TokenizeWithTrace(sourceFile.Content),
	}
}
