package main

type SourceFile struct {
	name    string
	content string
}

type GenEngine struct {
	triggers []GenTrigger
}

func NewGenEngine(triggers ...GenTrigger) *GenEngine {
	return &GenEngine{
		triggers: triggers,
	}
}

func (e *GenEngine) addTrigger(trigger GenTrigger) {
	e.triggers = append(e.triggers, trigger)
}

func (e *GenEngine) addTriggers(triggers ...GenTrigger) {
	for _, trigger := range triggers {
		e.addTrigger(trigger)
	}
}

func (e *GenEngine) gen(sourceFiles ...SourceFile) []CodeGen {
	ctx := &GenContext{
		wipCodeGen:  []CodeGen{},
		sourceFiles: sourceFiles,
	}

	for i, sf := range sourceFiles {
		ctx.sourceFilePos = i
		for _, gt := range e.triggers {
			p := NewBaseParser(sf)

			ctx.positionedTokens = p.positionedTokens

			for !p.atEOF() {
				posBefore := p.getPos()

				if ok, h := gt.parse(p); ok && h != nil {
					ctx.fileCursorPos = p.cursorPos()
					ctx.pos = p.getPos()

					h.handle(ctx)
				} else {
					// Reset the parser to the last position
					p.setPos(posBefore)
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

	positionedTokens []PositionedToken
	pos              int // Current position (index) in positionedTokens
}

type CodeGen struct {
	code       string
	outputPath string
}
