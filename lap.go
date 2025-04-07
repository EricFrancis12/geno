package main

import "fmt"

type SourceFile struct {
	name    string
	content string
}

type LAP struct {
	triggers []GenTrigger
}

func (l *LAP) addTrigger(trigger GenTrigger) {
	l.triggers = append(l.triggers, trigger)
}

func (l *LAP) genAll(sourceFiles ...SourceFile) []CodeGen {
	ctx := &GenContext{
		wipCodeGen:  []CodeGen{},
		sourceFiles: sourceFiles,
	}

	for i, sf := range sourceFiles {
		ctx.sourceFilePos = i
		for _, gt := range l.triggers {
			p := NewBasicParser(sf)

			ctx.tokens = p.tokens

			for !p.atEOF() {
				pos := p.getPos()

				if ok, h := gt.parse(p); ok && h != nil {
					ctx.fileCursorPos = p.fileCursorPos()
					ctx.tokensPos = p.getPos()

					h.handle(ctx)
				} else {
					// Reset the parser to the last position
					p.setPos(pos)
				}
			}
		}
	}

	return ctx.wipCodeGen
}

type ParseHandler interface {
	handle(*GenContext)
}

type GenTrigger interface {
	Token
	ParseHandler
}

type GenContext struct {
	wipCodeGen []CodeGen

	sourceFiles   []SourceFile
	sourceFilePos int // The index of the source file being parsed
	fileCursorPos int // The cursor will be directly to the right of the last token parsed when passed into gen()

	tokens    []TokenWithFilePos
	tokensPos int // Current position (index) in the token list
}

type CodeGen struct {
	code       string
	outputPath string
}

// -----
// ENUM
// -----

type EnumVariant struct {
	Key   string
	Value string
}

type PrismaEnum struct {
	Name     string
	Variants []EnumVariant
}

// The enum attempts to parse itself given it's current position in
// the token list. It can advance (mutate) the *Parser to check tokens
// for a match. It then returns a bool indicating if there was a match or not.
// If there is no match, it's the caller's responsibility to reset the
// parser's position to what it was before being passed into parse().
func (e PrismaEnum) parse(p *Parser) (bool, ParseHandler) {
	btk := p.advanceBasic()
	if btk.kind != ENUM {
		return false, nil
	}

	enum := PrismaEnum{
		Name:     btk.Value,
		Variants: []EnumVariant{},
	}

	if _, ok := p.advanceBasicTo(IDENTIFIER, OPEN_CURLY); !ok {
		return false, nil
	}

	for p.currentBasicToken().kind == IDENTIFIER {
		variantName := p.advanceBasic().Value
		enum.Variants = append(enum.Variants, EnumVariant{
			Key:   variantName,
			Value: variantName,
		})
	}

	if p.advanceBasic().kind != CLOSE_CURLY {
		return false, nil
	}

	return true, enum
}

// handle() is called whenever there is a match from parse(). The enum uses
// info in ParseContext to determine what CodeGen needs to be created
// based on the current context.
func (e PrismaEnum) handle(ctx *GenContext) {
	fmt.Println("TODO: Enum.handle()")

	p := NewBasicParser(ctx.sourceFiles[ctx.sourceFilePos])
	p.toFileCursorPos(ctx.fileCursorPos)

	// File cursor starts positioned past the closing curly brace token, and
	// to the immediate left of the next token in line (comment token)
	fmt.Printf("cursorPos before: %d\n", p.fileCursorPos())
	fmt.Printf("left slice before: %s\n", ctx.sourceFiles[ctx.sourceFilePos].content[:p.fileCursorPos()])

	for p.advanceBasicN(-1).kind != ENUM {
		// Advance the parser until we reach the enum token
	}

	// File cursor is now positioned to the immediate left of the enum token
	fmt.Printf("cursorPos after: %d\n", p.fileCursorPos())
	fmt.Printf("left slice after: %s\n", ctx.sourceFiles[ctx.sourceFilePos].content[:p.fileCursorPos()])
}
