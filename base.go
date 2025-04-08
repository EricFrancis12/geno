package geno

type BaseToken struct {
	Kind  BaseTokenKind
	Value string
}

func NewBaseToken(kind BaseTokenKind, value string) BaseToken {
	return BaseToken{
		Kind:  kind,
		Value: value,
	}
}

func (t BaseToken) WithPos(cursorPos int) PositionedToken {
	return NewPositionedToken(t, cursorPos)
}

func (t BaseToken) Parse(p *Parser) (bool, ParseHandler) {
	bt, ok := p.Advance().(BaseToken)
	if !ok {
		return false, nil
	}
	return t == bt, nil
}

type BaseTokenKind int

// TODO: categorize these into groups better
const (
	EOF BaseTokenKind = iota
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

	// Equivalence
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

	// Maths
	PLUS
	DASH
	SLASH
	STAR
	PERCENT

	// Reserved Keywords
	NULL
	TRUE
	FALSE
	VAR
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
	PACKAGE
	DEFER
	GO
	SELECT
	INTERFACE
	CHAN
	MAP
	STRUCT
	FALLTHROUGH
	BREAK
	CONTINUE
	RANGE
	RETURN
	SWITCH
	CASE
	DEFAULT
	ABSTRACT
	ASYNC
	AWAIT
	IMPLEMENTS
	NAMESPACE
	MODULE
	DECLARE
	PRIVATE
	PROTECTED
	PUBLIC
	READONLY
	STATIC
	SUPER
	YIELD
	AS
	ANY
	NEVER
	VOID

	// Misc
	COMMENT_DIRECTIVE
	UNKNOWN
)

var reservedTokensLookup map[string]BaseTokenKind = map[string]BaseTokenKind{
	"true":        TRUE,
	"false":       FALSE,
	"null":        NULL,
	"let":         LET,
	"const":       CONST,
	"class":       CLASS,
	"import":      IMPORT,
	"from":        FROM,
	"fn":          FN,
	"if":          IF,
	"else":        ELSE,
	"foreach":     FOREACH,
	"while":       WHILE,
	"for":         FOR,
	"export":      EXPORT,
	"typeof":      TYPEOF,
	"in":          IN,
	"enum":        ENUM,
	"type":        TYPE,
	"iota":        IOTA,
	"package":     PACKAGE,
	"defer":       DEFER,
	"go":          GO,
	"select":      SELECT,
	"interface":   INTERFACE,
	"chan":        CHAN,
	"map":         MAP,
	"struct":      STRUCT,
	"fallthrough": FALLTHROUGH,
	"break":       BREAK,
	"continue":    CONTINUE,
	"range":       RANGE,
	"return":      RETURN,
	"switch":      SWITCH,
	"case":        CASE,
	"default":     DEFAULT,
	"var":         VAR,
	"abstract":    ABSTRACT,
	"async":       ASYNC,
	"await":       AWAIT,
	"implements":  IMPLEMENTS,
	"namespace":   NAMESPACE,
	"module":      MODULE,
	"declare":     DECLARE,
	"private":     PRIVATE,
	"protected":   PROTECTED,
	"public":      PUBLIC,
	"readonly":    READONLY,
	"static":      STATIC,
	"super":       SUPER,
	"yield":       YIELD,
	"as":          AS,
	"any":         ANY,
	"never":       NEVER,
	"void":        VOID,
}
