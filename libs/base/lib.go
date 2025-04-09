package base

import "github.com/EricFrancis12/geno"

type BaseTokenLib struct {
	addlTokens []geno.Token
}

func (b *BaseTokenLib) AddToken(token geno.Token) {
	b.addlTokens = append(b.addlTokens, token)
}

func (b BaseTokenLib) Tokenize(source string) []*BaseToken {
	baseTokens := []*BaseToken{}
	for _, twp := range b.TokenizeWithPos(source) {
		baseTokens = append(baseTokens, twp.Token)
	}
	return baseTokens
}

func (b BaseTokenLib) TokenizeWithPos(source string) []geno.TokenWithCursorPos[*BaseToken] {
	l := NewBaseLexer(source)

outerLoop:
	for !l.AtEOF() {
		for _, pattern := range l.Patterns {
			loc := pattern.Regex.FindStringIndex(l.Remainder())
			if len(loc) != 0 && loc[0] == 0 {
				pattern.Handler(l, pattern.Regex)
				continue outerLoop
			}
		}
		l.AdvanceN(1)
		l.Push(NewBaseToken(UNKNOWN, l.Remainder()[:1]))
	}

	l.Push(NewBaseToken(EOF, "EOF"))
	return l.PositionedTokens
}
