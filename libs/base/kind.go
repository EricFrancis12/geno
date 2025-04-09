package base

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
