package base

import "github.com/EricFrancis12/geno"

type BaseTokenLib struct {
	addlTokens []geno.Token
}

func (b *BaseTokenLib) AddToken(token geno.Token) {
	b.addlTokens = append(b.addlTokens, token)
}

func (b BaseTokenLib) Tokenize(source string) []BaseToken {
	baseTokens := []BaseToken{}
	for _, twp := range b.TokenizeWithTrace(source) {
		baseTokens = append(baseTokens, twp.Token)
	}
	return baseTokens
}

func (b BaseTokenLib) TokenizeWithTrace(source string) []geno.TokenFromSource[BaseToken] {
	l := NewBaseLexer(source)

	for !l.AtEOF() {
		l.Match()
	}

	return l.TokensFromSource
}

func (b BaseTokenLib) TokenizeWithTraceAddl(source string) []geno.TokenFromSource[geno.Token] {
	result := []geno.TokenFromSource[geno.Token]{}

	l := NewBaseLexer(source)

eofLoop:
	for !l.AtEOF() {
		for _, tfs := range l.TokensFromSource {
			result = append(result, tfs.Generalize())
		}

		// Reset slice
		l.TokensFromSource = []geno.TokenFromSource[BaseToken]{}

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
