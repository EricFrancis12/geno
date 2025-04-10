package geno

import "github.com/EricFrancis12/geno/utils"

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

type tokenUsableParser[T Token] struct {
	*Parser[T]
}

func (t tokenUsableParser[T]) CurrentToken() Token {
	return t.Parser.CurrentToken()
}

func (t tokenUsableParser[T]) Advance() Token {
	return t.Parser.Advance()
}

func (t tokenUsableParser[T]) AdvanceN(n int) Token {
	return t.Parser.AdvanceN(n)
}

func (p *Parser[T]) ToTokenUsable() TokenParser {
	return tokenUsableParser[T]{
		Parser: p,
	}
}
