package geno

import (
	"fmt"
	"strings"

	"github.com/EricFrancis12/geno/utils"
)

type Parser[T Token] struct {
	SourceFile SourceFile

	TokensFromSource []TokenFromSource[T]
	Index            int
}

func NewParser[T Token](sourceFile SourceFile, tokenLib TokenLib[T]) *Parser[T] {
	return &Parser[T]{
		SourceFile:       sourceFile,
		TokensFromSource: tokenLib.TokenizeWithTrace(sourceFile.Content),
	}
}

func (p *Parser[T]) AtEOF() bool {
	return p.Index >= len(p.TokensFromSource)
}

func (p *Parser[T]) Pos() int {
	return p.Index
}

func (p *Parser[T]) SetPos(pos int) {
	p.Index = pos
}

func (p *Parser[T]) CursorPos() int {
	return p.TokensFromSource[p.Index].CursorPos
}

func (p *Parser[T]) SeekTokenAt(cursorPos int) {
	p.SetPos(0) // Reset the parser position to the start of the file
	for p.CursorPos() < cursorPos {
		p.Advance() // Advance the parser until we reach the desired file cursor position
	}
}

func (p *Parser[T]) CurrentToken() T {
	return p.TokensFromSource[p.Index].Token
}

func (p *Parser[T]) Advance() T {
	return p.AdvanceN(1)
}

func (p *Parser[T]) AdvanceN(n int) T {
	tk := p.CurrentToken()

	p.Index = utils.Clamp(
		p.Index+n,
		0,
		len(p.TokensFromSource)-1,
	)

	return tk
}

func (p *Parser[T]) GetSourceFile() SourceFile {
	return p.SourceFile
}

func (p *Parser[T]) Remainder() string {
	return p.SourceFile.Content[p.CursorPos():]
}

func (p *Parser[T]) Parse(token Token) (Token, error) {
	rem := p.Remainder()

	tk, took := token.FindString(rem)
	if tk == nil {
		return nil, fmt.Errorf(
			"token '%s' did not match near '%s'",
			token.String(),
			utils.AddSuffixIfLength(rem, 20, "..."),
		)
	}

	wip := ""

	for !p.AtEOF() {
		wip += p.Advance().String()

		if wip == took {
			return tk, nil
		} else if !strings.HasPrefix(took, wip) {
			return nil, fmt.Errorf(
				"expected '%s', to have prefix '%s'",
				took,
				wip,
			)
		}
	}

	return nil, fmt.Errorf("eof")
}
