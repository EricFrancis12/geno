package base

import (
	"regexp"
	"strings"

	"github.com/EricFrancis12/geno"
)

type Lexer struct {
	Patterns         []RegexPattern
	TokensFromSource []geno.TokenFromSource[BaseToken]
	Source           string
	CursorPos        int
}

type RegexPattern struct {
	Regex   *regexp.Regexp
	Handler RegexHandler
}

type RegexHandler func(l *Lexer, regex *regexp.Regexp)

func (l Lexer) AtEOF() bool {
	return l.CursorPos >= len(l.Source)
}

func (l Lexer) Remainder() string {
	return l.Source[l.CursorPos:]
}

func (l *Lexer) AdvanceN(n int) {
	l.CursorPos += n
}

func (l *Lexer) Push(bt BaseToken) {
	l.TokensFromSource = append(l.TokensFromSource, bt.WithPos(l.CursorPos))
}

func (l *Lexer) Match() {
	for _, pattern := range l.Patterns {
		loc := pattern.Regex.FindStringIndex(l.Remainder())
		if len(loc) != 0 && loc[0] == 0 {
			pattern.Handler(l, pattern.Regex)
			return
		}
	}
	l.AdvanceN(1)
	l.Push(NewBaseToken(UNKNOWN, l.Remainder()[:1]))
}

func NewBaseLexer(source string) *Lexer {
	return &Lexer{
		Source:           source,
		TokensFromSource: []geno.TokenFromSource[BaseToken]{},
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

// Default handler which will simply create a token with the matched contents.
// This handler is used with most simple tokens.
func defaultHandler(kind BaseTokenKind, value string) RegexHandler {
	return func(l *Lexer, _ *regexp.Regexp) {
		l.Push(NewBaseToken(kind, value))
		l.AdvanceN(len(value))
	}
}

func stringHandler(l *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(l.Remainder())
	stringLiteral := l.Remainder()[match[0]:match[1]]

	quotesRemoved := stringLiteral[1:(len(stringLiteral) - 1)]
	l.Push(NewBaseToken(STRING, quotesRemoved))
	l.AdvanceN(len(stringLiteral))
}

func numberHandler(l *Lexer, regex *regexp.Regexp) {
	match := regex.FindString(l.Remainder())
	l.Push(NewBaseToken(NUMBER, match))
	l.AdvanceN(len(match))
}

func symbolHandler(l *Lexer, regex *regexp.Regexp) {
	match := regex.FindString(l.Remainder())
	t := NewBaseToken(IDENTIFIER, match)
	if kind, found := reservedTokensLookup[match]; found {
		t = NewBaseToken(kind, match)
	}

	l.Push(t)
	l.AdvanceN(len(match))
}

func skipHandler(l *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(l.Remainder())
	l.AdvanceN(match[1])
}

func commentHandler(l *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(l.Remainder())
	if match != nil {
		commentLiteral := l.Remainder()[match[0]:match[1]]
		if strings.HasPrefix(commentLiteral, "//") {
			// Advance past the entire comment.
			l.AdvanceN(match[1])
		}
	}
}
