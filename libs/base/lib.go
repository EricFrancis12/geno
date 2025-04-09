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
	for _, twp := range b.TokenizeWithPos(source) {
		baseTokens = append(baseTokens, twp.Token)
	}
	return baseTokens
}

func (b BaseTokenLib) TokenizeWithPos(source string) []geno.TokenWithCursorPos[BaseToken] {
	l := NewBaseLexer(source)

	for !l.AtEOF() {
		l.Match()
	}

	return l.PositionedTokens
}

func (b BaseTokenLib) TokenizeWithPosAddl(source string) []geno.TokenWithCursorPos[geno.Token] {
	result := []geno.TokenWithCursorPos[geno.Token]{}

	l := NewBaseLexer(source)

eofLoop:
	for !l.AtEOF() {
		for _, pt := range l.PositionedTokens {
			twcp, ok := any(pt).(geno.TokenWithCursorPos[geno.Token])
			if !ok {
				panic("TODO")
			}
			result = append(result, twcp)
		}

		// Reset slice
		l.PositionedTokens = []geno.TokenWithCursorPos[BaseToken]{}

		for _, tk := range b.addlTokens {
			_tk, took := tk.FindString(l.Remainder())
			if _tk != nil && len(took) > 0 {
				result = append(result, geno.TokenWithCursorPos[geno.Token]{
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
