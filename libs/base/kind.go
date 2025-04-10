package base

type BaseTokenKind int

const (
	// Unknown (default) kind
	UNKNOWN BaseTokenKind = iota

	// Literals
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

	// Assignment & Comparison Operators
	ASSIGNMENT // =
	EQUALS     // ==
	NOT_EQUALS // !=
	NOT        // !

	// Relational Operators
	LESS           // <
	LESS_EQUALS    // <=
	GREATER        // >
	GREATER_EQUALS // >=

	// Logical Operators
	OR  // ||
	AND // &&

	// Symbols & Punctuation
	DOT        // .
	DOT_DOT    // ..
	SEMI_COLON // ;
	COLON      // :
	QUESTION   // ?
	COMMA      // ,
	HASHTAG    // #

	// Increment/Decrement & Compound Assignment
	PLUS_PLUS          // ++
	MINUS_MINUS        // --
	PLUS_EQUALS        // +=
	MINUS_EQUALS       // -=
	NULLISH_ASSIGNMENT // ??=

	// Arithmetic Operators
	PLUS    // +
	DASH    // -
	SLASH   // /
	STAR    // *
	PERCENT // %

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

	// Advanced/Modern Keywords
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
)
