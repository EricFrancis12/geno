package main

type Token interface {
	parse(*Parser) (bool, ParseHandler)
}

type TokenWithFilePos struct {
	Token   Token
	FilePos int // Length of the token in the source file
}

type BasicToken struct {
	kind  TokenKind
	Value string
}

func newBasicToken(kind TokenKind, value string) BasicToken {
	return BasicToken{
		kind:  kind,
		Value: value,
	}
}

func (t BasicToken) parse(p *Parser) (bool, ParseHandler) {
	bt, ok := p.advance().(BasicToken)
	if !ok {
		return false, nil
	}
	return t == bt, nil
}

type TokenKind int

const (
	EOF TokenKind = iota
	NULL
	TRUE
	FALSE
	NUMBER
	STRING
	IDENTIFIER

	// Grouping & Braces
	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	// Equivilance
	ASSIGNMENT
	EQUALS
	NOT_EQUALS
	NOT

	// Conditional
	LESS
	LESS_EQUALS
	GREATER
	GREATER_EQUALS

	// Logical
	OR
	AND

	// Symbols
	DOT
	DOT_DOT
	SEMI_COLON
	COLON
	QUESTION
	COMMA
	HASHTAG

	// Shorthand
	PLUS_PLUS
	MINUS_MINUS
	PLUS_EQUALS
	MINUS_EQUALS
	NULLISH_ASSIGNMENT // ??=

	//Maths
	PLUS
	DASH
	SLASH
	STAR
	PERCENT

	// Reserved Keywords
	LET
	CONST
	CLASS
	IMPORT
	FROM
	FN
	IF
	ELSE
	FOREACH
	WHILE
	FOR
	EXPORT
	TYPEOF
	IN
	ENUM
	TYPE
	IOTA

	// Misc
	COMMENT

	// Unknown
	UNKNOWN
)

var reservedTokensLookup map[string]TokenKind = map[string]TokenKind{
	"true":    TRUE,
	"false":   FALSE,
	"null":    NULL,
	"let":     LET,
	"const":   CONST,
	"class":   CLASS,
	"import":  IMPORT,
	"from":    FROM,
	"fn":      FN,
	"if":      IF,
	"else":    ELSE,
	"foreach": FOREACH,
	"while":   WHILE,
	"for":     FOR,
	"export":  EXPORT,
	"typeof":  TYPEOF,
	"in":      IN,
	"enum":    ENUM,
	"type":    TYPE,
	"iota":    IOTA,
}
