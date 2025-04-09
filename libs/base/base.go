package base

import "github.com/EricFrancis12/geno"

func NewBaseGenEngine(triggers ...geno.GenTrigger[BaseToken]) *geno.GenEngine[BaseToken] {
	return &geno.GenEngine[BaseToken]{
		TokenLib: BaseTokenLib{},
		Triggers: triggers,
	}
}

func NewBaseParser(sourceFile geno.SourceFile) *geno.Parser[BaseToken] {
	return &geno.Parser[BaseToken]{
		SourceFile:       sourceFile,
		PositionedTokens: BaseTokenLib{}.TokenizeWithPos(sourceFile.Content),
	}
}
