package main

import "fmt"

type PrismaEnum struct {
	Name     string
	Variants []EnumVariant
}

type EnumVariant struct {
	Key   string
	Value string
}

// The enum attempts to parse itself given it's current position in
// the token list. It can advance (mutate) the *Parser to check tokens
// for a match. It then returns a bool indicating if there was a match or not.
// If there is no match, it's the caller's responsibility to reset the
// parser's position to what it was before being passed into parse().
func (e PrismaEnum) parse(p *Parser) (bool, ParseHandler) {
	btk := p.advanceBase()
	if btk.kind != ENUM {
		return false, nil
	}

	enum := PrismaEnum{
		Name:     btk.Value,
		Variants: []EnumVariant{},
	}

	if _, ok := p.advanceBaseTo(IDENTIFIER, OPEN_CURLY); !ok {
		return false, nil
	}

	for p.currentBaseToken().kind == IDENTIFIER {
		variantName := p.advanceBase().Value
		enum.Variants = append(enum.Variants, EnumVariant{
			Key:   variantName,
			Value: variantName,
		})
	}

	if p.advanceBase().kind != CLOSE_CURLY {
		return false, nil
	}

	return true, enum
}

// handle() is called whenever there is a match from parse(). The enum uses
// info in ParseContext to determine what CodeGen needs to be created
// based on the current context.
func (e PrismaEnum) handle(ctx *GenContext) {
	fmt.Println("TODO: Enum.handle()")

	p := NewBaseParser(ctx.sourceFiles[ctx.sourceFilePos])
	p.seekToNearestToken(ctx.fileCursorPos)

	// File cursor starts positioned past the closing curly brace token, and
	// to the immediate left of the next token in line (comment token)
	fmt.Printf("cursorPos before: %d\n", p.cursorPos())
	fmt.Printf("left slice before: %s\n", ctx.sourceFiles[ctx.sourceFilePos].content[:p.cursorPos()])

	for p.advanceBaseN(-1).kind != ENUM {
		// Advance the parser until we reach the enum token
	}

	// File cursor is now positioned to the immediate left of the enum token
	fmt.Printf("cursorPos after: %d\n", p.cursorPos())
	fmt.Printf("left slice after: %s\n", ctx.sourceFiles[ctx.sourceFilePos].content[:p.cursorPos()])
}
