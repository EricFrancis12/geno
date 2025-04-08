package geno

import (
	"regexp"
	"strings"
)

type RegexPattern struct {
	Regex   *regexp.Regexp
	Handler regexHandler
}

type Lexer struct {
	Patterns         []RegexPattern
	PositionedTokens []PositionedToken
	Source           string
	CursorPos        int
}

func Tokenize(source string) []PositionedToken {
	lex := CreateLexer(source)

outerLoop:
	for !lex.AtEOF() {
		for _, pattern := range lex.Patterns {
			loc := pattern.Regex.FindStringIndex(lex.Remainder())
			if len(loc) != 0 && loc[0] == 0 {
				pattern.Handler(lex, pattern.Regex)
				continue outerLoop
			}
		}
		lex.AdvanceN(1)
		lex.Push(NewBaseToken(UNKNOWN, lex.Remainder()[:1]))
	}

	lex.Push(NewBaseToken(EOF, "EOF"))
	return lex.PositionedTokens
}

func (lex *Lexer) AdvanceN(n int) {
	lex.CursorPos += n
}

func (lex *Lexer) Remainder() string {
	return lex.Source[lex.CursorPos:]
}

func (lex *Lexer) Push(bt BaseToken) {
	lex.PositionedTokens = append(lex.PositionedTokens, bt.WithPos(lex.CursorPos))
}

func (lex *Lexer) AtEOF() bool {
	return lex.CursorPos >= len(lex.Source)
}

func CreateLexer(source string) *Lexer {
	return &Lexer{
		CursorPos:        0,
		Source:           source,
		PositionedTokens: make([]PositionedToken, 0),
		Patterns: []RegexPattern{
			{regexp.MustCompile(`\s+`), skipHandler},
			{regexp.MustCompile(`\/\/.*`), commentHandler},
			{regexp.MustCompile(`"[^"]*"`), stringHandler},
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			{regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`), symbolHandler},
			{regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			{regexp.MustCompile(`\]`), defaultHandler(CLOSE_BRACKET, "]")},
			{regexp.MustCompile(`\{`), defaultHandler(OPEN_CURLY, "{")},
			{regexp.MustCompile(`\}`), defaultHandler(CLOSE_CURLY, "}")},
			{regexp.MustCompile(`\(`), defaultHandler(OPEN_PAREN, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(CLOSE_PAREN, ")")},
			{regexp.MustCompile(`==`), defaultHandler(EQUALS, "==")},
			{regexp.MustCompile(`!=`), defaultHandler(NOT_EQUALS, "!=")},
			{regexp.MustCompile(`=`), defaultHandler(ASSIGNMENT, "=")},
			{regexp.MustCompile(`!`), defaultHandler(NOT, "!")},
			{regexp.MustCompile(`<=`), defaultHandler(LESS_EQUALS, "<=")},
			{regexp.MustCompile(`<`), defaultHandler(LESS, "<")},
			{regexp.MustCompile(`>=`), defaultHandler(GREATER_EQUALS, ">=")},
			{regexp.MustCompile(`>`), defaultHandler(GREATER, ">")},
			{regexp.MustCompile(`\|\|`), defaultHandler(OR, "||")},
			{regexp.MustCompile(`&&`), defaultHandler(AND, "&&")},
			{regexp.MustCompile(`\.\.`), defaultHandler(DOT_DOT, "..")},
			{regexp.MustCompile(`\.`), defaultHandler(DOT, ".")},
			{regexp.MustCompile(`;`), defaultHandler(SEMI_COLON, ";")},
			{regexp.MustCompile(`:`), defaultHandler(COLON, ":")},
			{regexp.MustCompile(`\?\?=`), defaultHandler(NULLISH_ASSIGNMENT, "??=")},
			{regexp.MustCompile(`\?`), defaultHandler(QUESTION, "?")},
			{regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
			{regexp.MustCompile(`\+\+`), defaultHandler(PLUS_PLUS, "++")},
			{regexp.MustCompile(`--`), defaultHandler(MINUS_MINUS, "--")},
			{regexp.MustCompile(`\+=`), defaultHandler(PLUS_EQUALS, "+=")},
			{regexp.MustCompile(`-=`), defaultHandler(MINUS_EQUALS, "-=")},
			{regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`-`), defaultHandler(DASH, "-")},
			{regexp.MustCompile(`/`), defaultHandler(SLASH, "/")},
			{regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
			{regexp.MustCompile(`%`), defaultHandler(PERCENT, "%")},
			{regexp.MustCompile(`#`), defaultHandler(HASHTAG, "#")},
		},
	}
}

type regexHandler func(lex *Lexer, regex *regexp.Regexp)

// Default handler which will simply create a token with the matched contents.
// This handler is used with most simple tokens.
func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *Lexer, _ *regexp.Regexp) {
		lex.Push(NewBaseToken(kind, value))
		lex.AdvanceN(len(value))
	}
}

func stringHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.Remainder())
	stringLiteral := lex.Remainder()[match[0]:match[1]]

	lex.Push(NewBaseToken(STRING, stringLiteral))
	lex.AdvanceN(len(stringLiteral))
}

func numberHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.Remainder())
	lex.Push(NewBaseToken(NUMBER, match))
	lex.AdvanceN(len(match))
}

func symbolHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.Remainder())
	t := NewBaseToken(IDENTIFIER, match)
	if kind, found := reservedTokensLookup[match]; found {
		t = NewBaseToken(kind, match)
	}

	lex.Push(t)
	lex.AdvanceN(len(match))
}

func skipHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.Remainder())
	lex.AdvanceN(match[1])
}

func commentHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.Remainder())
	if match != nil {
		commentLiteral := lex.Remainder()[match[0]:match[1]]
		if strings.HasPrefix(commentLiteral, "//") {
			lex.Push(NewBaseToken(COMMENT, commentLiteral))

			// Advance past the entire comment.
			lex.AdvanceN(match[1])
		}
	}
}
