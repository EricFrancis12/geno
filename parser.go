package geno

type Parser[T Token] struct {
	SourceFile SourceFile

	PositionedTokens []TokenWithCursorPos[T]
	Index            int
}

func NewParser[T Token](sourceFile SourceFile, tokenLib TokenLib[T]) *Parser[T] {
	return &Parser[T]{
		SourceFile:       sourceFile,
		PositionedTokens: tokenLib.TokenizeWithPos(sourceFile.Content),
	}
}

func (p *Parser[T]) AtEOF() bool {
	return p.Index >= len(p.PositionedTokens)
}

func (p *Parser[T]) Pos() int {
	return p.Index
}

func (p *Parser[T]) SetPos(pos int) {
	p.Index = pos
}

func (p *Parser[T]) CursorPos() int {
	return p.PositionedTokens[p.Index].CursorPos
}

func (p *Parser[T]) SeekTokenAt(cursorPos int) {
	p.SetPos(0) // Reset the parser position to the start of the file
	for p.CursorPos() < cursorPos {
		p.Advance() // Advance the parser until we reach the desired file cursor position
	}
}

func (p *Parser[T]) CurrentToken() T {
	return p.PositionedTokens[p.Index].Token
}

func (p *Parser[T]) Advance() T {
	return p.AdvanceN(1)
}

func (p *Parser[T]) AdvanceN(n int) T {
	tk := p.CurrentToken()
	newIndex := p.Index + n

	// If n is negative, we don't want to go past the start of the token list
	if newIndex < 0 {
		newIndex = 0
	}
	p.Index = newIndex

	return tk
}
