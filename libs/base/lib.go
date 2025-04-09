package base

import "github.com/EricFrancis12/geno"

type BaseTokenLib struct{}

func (b BaseTokenLib) Tokenize(source string) []*BaseToken {
	baseTokens := []*BaseToken{}
	for _, twp := range b.TokenizeWithPos(source) {
		baseTokens = append(baseTokens, twp.Token)
	}
	return baseTokens
}

func (b BaseTokenLib) TokenizeWithPos(source string) []geno.TokenWithCursorPos[*BaseToken] {
	lex := CreateBaseLexer(source)

outerLoop:
	for !lex.AtEOF() {
		for _, pattern := range lex.Patterns {
			loc := pattern.Regex.FindStringIndex(lex.Remainder())
			if len(loc) != 0 && loc[0] == 0 {
				pattern.Handler(lex, pattern.Regex)
				continue outerLoop
			}
		}
		lex.AdvanceN(1)
		lex.Push(NewBaseToken(UNKNOWN, lex.Remainder()[:1]))
	}

	lex.Push(NewBaseToken(EOF, "EOF"))
	return lex.PositionedTokens
}
