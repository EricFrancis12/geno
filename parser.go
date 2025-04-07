package main

type Parser struct {
	sourceFile SourceFile

	tokens []TokenWithFilePos
	pos    int // Current position (index) in the token list
}

func NewParser(sourceFile SourceFile, tokenizerFunc func(string) []TokenWithFilePos) *Parser {
	return &Parser{
		sourceFile: sourceFile,
		tokens:     tokenizerFunc(sourceFile.content),
	}
}

func NewBasicParser(sourceFile SourceFile) *Parser {
	return &Parser{
		sourceFile: sourceFile,
		tokens:     Tokenize(sourceFile.content),
	}
}

func (p *Parser) getPos() int {
	return p.pos
}

func (p *Parser) setPos(pos int) {
	p.pos = pos
}

func (p *Parser) fileCursorPos() int {
	return p.tokens[p.pos].FilePos
}

func (p *Parser) toFileCursorPos(fcp int) {
	p.setPos(0)
	for p.fileCursorPos() < fcp {
		p.advance() // Advance the parser until we reach the desired file cursor position
	}
}

func (p *Parser) currentToken() Token {
	return p.tokens[p.pos].Token
}

func (p *Parser) advance() Token {
	return p.advanceN(1)
}

func (p *Parser) advanceN(n int) Token {
	tk := p.currentToken()
	newPos := p.pos + n
	if newPos < 0 {
		newPos = 0
	}
	p.pos = newPos
	return tk
}

func (p *Parser) advanceBasicTo(kinds ...TokenKind) (Token, bool) {
	i := 0
	for _, kind := range kinds {
		i++
		if p.advanceBasic().kind != kind {
			return p.currentBasicToken(), false
		}
	}
	return p.currentBasicToken(), true
}

func (p *Parser) currentBasicToken() BasicToken {
	tk, ok := p.currentToken().(BasicToken)
	if !ok {
		return BasicToken{
			kind: UNKNOWN,
		}
	}
	return tk
}

func (p *Parser) advanceBasic() BasicToken {
	tk, ok := p.advance().(BasicToken)
	if !ok {
		return BasicToken{
			kind: UNKNOWN,
		}
	}
	return tk
}

func (p *Parser) advanceBasicN(n int) BasicToken {
	tk, ok := p.advanceN(n).(BasicToken)
	if !ok {
		return BasicToken{
			kind: UNKNOWN,
		}
	}
	return tk
}

func (p *Parser) atEOF() bool {
	return p.pos >= len(p.tokens)
}
