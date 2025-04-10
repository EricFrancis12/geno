package custom

import (
	"github.com/EricFrancis12/geno"
	"github.com/EricFrancis12/geno/libs/base"
)

type CustomBaseTokenLib struct {
	addlTokens []geno.Token
}

func (b *CustomBaseTokenLib) AddToken(token geno.Token) {
	b.addlTokens = append(b.addlTokens, token)
}

func (b CustomBaseTokenLib) Tokenize(source string) []geno.Token {
	tokens := []geno.Token{}
	for _, twp := range b.TokenizeWithTrace(source) {
		tokens = append(tokens, twp.Token)
	}
	return tokens
}

func (b CustomBaseTokenLib) TokenizeWithTrace(source string) []geno.TokenFromSource[geno.Token] {
	result := []geno.TokenFromSource[geno.Token]{}

	l := base.NewBaseLexer(source)

eofLoop:
	for !l.AtEOF() {
		for _, tfs := range l.TokensFromSource {
			result = append(result, tfs.Generalize())
		}

		// Reset slice
		l.TokensFromSource = []geno.TokenFromSource[base.BaseToken]{}

		for _, tk := range b.addlTokens {
			_tk, took := tk.FindString(l.Remainder())
			if _tk != nil && len(took) > 0 {
				result = append(result, geno.TokenFromSource[geno.Token]{
					Token:     _tk,
					CursorPos: l.CursorPos,
				})
				l.AdvanceN(len(took))
				continue eofLoop
			}
		}

		l.Match()
	}

	return result
}
