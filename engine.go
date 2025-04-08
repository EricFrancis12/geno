package geno

type SourceFile struct {
	Name    string
	Content string
}

type GenEngine struct {
	Triggers []GenTrigger
}

func NewGenEngine(triggers ...GenTrigger) *GenEngine {
	return &GenEngine{
		Triggers: triggers,
	}
}

func (e *GenEngine) AddTrigger(trigger GenTrigger) {
	e.Triggers = append(e.Triggers, trigger)
}

func (e *GenEngine) AddTriggers(triggers ...GenTrigger) {
	for _, trigger := range triggers {
		e.AddTrigger(trigger)
	}
}

func (e *GenEngine) Gen(sourceFiles ...SourceFile) []CodeGen {
	ctx := &GenContext{
		WipCodeGen:  []CodeGen{},
		SourceFiles: sourceFiles,
	}

	for i, sf := range sourceFiles {
		ctx.SourceFilePos = i
		for _, gt := range e.Triggers {
			p := NewBaseParser(sf)

			ctx.PositionedTokens = p.PositionedTokens

			for !p.AtEOF() {
				posBefore := p.GetPos()

				if ok, h := gt.Parse(p); ok && h != nil {
					ctx.FileCursorPos = p.CursorPos()
					ctx.Pos = p.GetPos()

					h.Handle(ctx)
				} else {
					// Reset the parser to the last position + 1 to advance to the next token
					p.SetPos(posBefore + 1)
				}
			}
		}
	}

	return ctx.WipCodeGen
}

type ParseHandler interface {
	Handle(*GenContext)
}

type GenTrigger interface {
	Token
	ParseHandler
}

type GenContext struct {
	WipCodeGen []CodeGen

	SourceFiles   []SourceFile
	SourceFilePos int // The index of the source file being parsed
	FileCursorPos int // The cursor will be directly to the right of the last token parsed when passed into gen()

	PositionedTokens []PositionedToken
	Pos              int // Current position (index) in positionedTokens
}

type CodeGen struct {
	Code       string
	OutputPath string
}
