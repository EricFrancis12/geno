package main

type Parser struct {
	sourceFile SourceFile

	positionedTokens []PositionedToken
	pos              int // Current position (index) in positionedTokens
}

func NewParser(sourceFile SourceFile, tokenizerFunc func(string) []PositionedToken) *Parser {
	return &Parser{
		sourceFile:       sourceFile,
		positionedTokens: tokenizerFunc(sourceFile.content),
	}
}

func NewBaseParser(sourceFile SourceFile) *Parser {
	return &Parser{
		sourceFile:       sourceFile,
		positionedTokens: Tokenize(sourceFile.content),
	}
}

func (p *Parser) atEOF() bool {
	return p.pos >= len(p.positionedTokens)
}

func (p *Parser) getPos() int {
	return p.pos
}

func (p *Parser) setPos(pos int) {
	p.pos = pos
}

func (p *Parser) cursorPos() int {
	return p.positionedTokens[p.pos].CursorPos
}

func (p *Parser) seekToNearestToken(cursorPos int) {
	p.setPos(0) // Reset the parser position to the start of the file
	for p.cursorPos() < cursorPos {
		p.advance() // Advance the parser until we reach the desired file cursor position
	}
}

func (p *Parser) currentToken() Token {
	return p.positionedTokens[p.pos].Token
}

func (p *Parser) advance() Token {
	return p.advanceN(1)
}

func (p *Parser) advanceN(n int) Token {
	tk := p.currentToken()
	newPos := p.pos + n
	// If n is negative, we don't want to go past the start of the token list
	if newPos < 0 {
		newPos = 0
	}
	p.pos = newPos
	return tk
}

func (p *Parser) advanceBaseTo(kinds ...TokenKind) (BaseToken, bool) {
	for _, kind := range kinds {
		if p.advanceBase().kind != kind {
			return p.currentBaseToken(), false
		}
	}
	return p.currentBaseToken(), true
}

func (p *Parser) currentBaseToken() BaseToken {
	tk, ok := p.currentToken().(BaseToken)
	if !ok {
		return BaseToken{
			kind: UNKNOWN,
		}
	}
	return tk
}

func (p *Parser) advanceBase() BaseToken {
	return p.advanceBaseN(1)
}

func (p *Parser) advanceBaseN(n int) BaseToken {
	tk, ok := p.advanceN(n).(BaseToken)
	if !ok {
		return BaseToken{
			kind: UNKNOWN,
		}
	}
	return tk
}
