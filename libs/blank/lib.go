package blank

import (
	"github.com/EricFrancis12/geno"
	"github.com/EricFrancis12/geno/libs/base"
)

type BlankTokenLib struct {
	Tokens []geno.Token
}

func (b *BlankTokenLib) AddToken(token geno.Token) {
	b.Tokens = append(b.Tokens, token)
}

func (b BlankTokenLib) Tokenize(source string) []geno.Token {
	tokens := []geno.Token{}
	for _, twp := range b.TokenizeWithTrace(source) {
		tokens = append(tokens, twp.Token)
	}
	return tokens
}

func (b BlankTokenLib) TokenizeWithTrace(source string) []geno.TokenFromSource[geno.Token] {
	result := []geno.TokenFromSource[geno.Token]{}

	l := base.NewBaseLexer(source)

eofLoop:
	for !l.AtEOF() {
		for _, tk := range b.Tokens {
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

		l.AdvanceN(1)
	}

	return result
}
