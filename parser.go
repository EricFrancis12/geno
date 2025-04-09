package geno

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
	newIndex := p.Index + n

	// If n is negative, we don't want to go past the start of the token list
	if newIndex < 0 {
		newIndex = 0
	}
	p.Index = newIndex

	return tk
}

func (p *Parser[T]) GetSourceFile() SourceFile {
	return p.SourceFile
}

func (p *Parser[T]) Remainder() string {
	return p.SourceFile.Content[p.CursorPos():]
}

func (p *Parser[T]) Generalize() TokenParser {
	tp, ok := any(p).(TokenParser)
	if !ok {
		panic("expected *Parser[T] to be convertable to TokenParser")
	}
	return tp
}
