package main

import (
	"regexp"
	"strings"
)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	patterns []regexPattern
	Tokens   []TokenWithFilePos
	source   string
	pos      int
}

func Tokenize(source string) []TokenWithFilePos {
	lex := createLexer(source)

outerLoop:
	for !lex.atEOF() {
		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainder())
			if len(loc) != 0 && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				continue outerLoop
			}
		}
		lex.advanceN(1)

		t := TokenWithFilePos{
			Token:   newBasicToken(UNKNOWN, lex.remainder()[:1]),
			FilePos: lex.pos + 1,
		}
		lex.push(t)
	}

	t := TokenWithFilePos{
		Token:   newBasicToken(EOF, "EOF"),
		FilePos: lex.pos,
	}
	lex.push(t)
	return lex.Tokens
}

func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

func (lex *lexer) push(token TokenWithFilePos) {
	lex.Tokens = append(lex.Tokens, token)
}

func (lex *lexer) atEOF() bool {
	return lex.pos >= len(lex.source)
}

func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]TokenWithFilePos, 0),
		patterns: []regexPattern{
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

type regexHandler func(lex *lexer, regex *regexp.Regexp)

// Default handler which will simply create a token with the matched contents.
// This handler is used with most simple tokens.
func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *lexer, _ *regexp.Regexp) {
		t := TokenWithFilePos{
			Token:   newBasicToken(kind, value),
			FilePos: lex.pos,
		}
		lex.push(t)

		lex.advanceN(len(value))
	}
}

func stringHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	stringLiteral := lex.remainder()[match[0]:match[1]]

	t := TokenWithFilePos{
		Token:   newBasicToken(STRING, stringLiteral),
		FilePos: lex.pos,
	}
	lex.push(t)

	lex.advanceN(len(stringLiteral))
}

func numberHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())

	t := TokenWithFilePos{
		Token:   newBasicToken(NUMBER, match),
		FilePos: lex.pos,
	}
	lex.push(t)

	lex.advanceN(len(match))
}

func symbolHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	t := TokenWithFilePos{
		Token:   newBasicToken(IDENTIFIER, match),
		FilePos: lex.pos,
	}
	if kind, found := reservedTokensLookup[match]; found {
		t.Token = newBasicToken(kind, match)
	}
	lex.push(t)

	lex.advanceN(len(match))
}

func skipHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advanceN(match[1])
}

func commentHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	if match != nil {
		commentLiteral := lex.remainder()[match[0]:match[1]]
		if strings.HasPrefix(commentLiteral, "//") {
			t := TokenWithFilePos{
				Token:   newBasicToken(COMMENT, commentLiteral),
				FilePos: lex.pos,
			}
			lex.push(t)

			// Advance past the entire comment.
			lex.advanceN(match[1])
		}
	}
}
