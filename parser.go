package geno

type Parser struct {
	SourceFile SourceFile

	PositionedTokens []PositionedToken
	Pos              int // Current position (index) in positionedTokens
}

func NewParser(sourceFile SourceFile, tokenizerFunc func(string) []PositionedToken) *Parser {
	return &Parser{
		SourceFile:       sourceFile,
		PositionedTokens: tokenizerFunc(sourceFile.Content),
	}
}

func NewBaseParser(sourceFile SourceFile) *Parser {
	return &Parser{
		SourceFile:       sourceFile,
		PositionedTokens: Tokenize(sourceFile.Content),
	}
}

func (p *Parser) AtEOF() bool {
	return p.Pos >= len(p.PositionedTokens)
}

func (p *Parser) GetPos() int {
	return p.Pos
}

func (p *Parser) SetPos(pos int) {
	p.Pos = pos
}

func (p *Parser) CursorPos() int {
	return p.PositionedTokens[p.Pos].CursorPos
}

func (p *Parser) SeekTokenAt(cursorPos int) {
	p.SetPos(0) // Reset the parser position to the start of the file
	for p.CursorPos() < cursorPos {
		p.Advance() // Advance the parser until we reach the desired file cursor position
	}
}

func (p *Parser) CurrentToken() Token {
	return p.PositionedTokens[p.Pos].Token
}

func (p *Parser) Advance() Token {
	return p.AdvanceN(1)
}

func (p *Parser) AdvanceN(n int) Token {
	tk := p.CurrentToken()
	newPos := p.Pos + n
	// If n is negative, we don't want to go past the start of the token list
	if newPos < 0 {
		newPos = 0
	}
	p.Pos = newPos
	return tk
}

func (p *Parser) AdvanceBaseTo(kinds ...TokenKind) (BaseToken, bool) {
	for _, kind := range kinds {
		if p.AdvanceBase().Kind != kind {
			return p.CurrentBaseToken(), false
		}
	}
	return p.CurrentBaseToken(), true
}

func (p *Parser) CurrentBaseToken() BaseToken {
	tk, ok := p.CurrentToken().(BaseToken)
	if !ok {
		return BaseToken{
			Kind: UNKNOWN,
		}
	}
	return tk
}

func (p *Parser) AdvanceBase() BaseToken {
	return p.AdvanceBaseN(1)
}

func (p *Parser) AdvanceBaseN(n int) BaseToken {
	tk, ok := p.AdvanceN(n).(BaseToken)
	if !ok {
		return BaseToken{
			Kind: UNKNOWN,
		}
	}
	return tk
}
