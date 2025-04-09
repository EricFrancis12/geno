package geno

type GenEngine[T Token] struct {
	TokenLib TokenLib[T]
	Triggers []GenTrigger[T]
}

func NewGenEngine[T Token](tokenLib TokenLib[T], triggers ...GenTrigger[T]) *GenEngine[T] {
	return &GenEngine[T]{
		TokenLib: tokenLib,
		Triggers: triggers,
	}
}

func (e *GenEngine[T]) AddTrigger(trigger GenTrigger[T]) {
	e.Triggers = append(e.Triggers, trigger)
}

func (e *GenEngine[T]) AddTriggers(triggers ...GenTrigger[T]) {
	for _, trigger := range triggers {
		e.AddTrigger(trigger)
	}
}

func (e GenEngine[T]) Gen(sourceFiles ...SourceFile) []CodeGen {
	ctx := &GenContext[T]{
		WipCodeGen:  []CodeGen{},
		SourceFiles: sourceFiles,
	}

	for i, sf := range sourceFiles {
		ctx.SourceFilePos = i

		// Copy triggers, because they are to be mutated on successful .Parse()
		triggers := e.Triggers

		for _, gt := range triggers {
			p := NewParser(sf, e.TokenLib)

			// Update context positioned tokens
			ctx.PositionedTokens = p.PositionedTokens

			for !p.AtEOF() {
				posBefore := p.Pos()

				tp, ok := any(p).(TokenParser)
				if !ok {
					panic("expected *Parser[T] to be convertable to TokenParser")
				}

				if gt.Parse(tp) != nil {
					// Reset the parser to the last position + 1 to advance to the next token
					p.SetPos(posBefore + 1)
				} else {
					ctx.FileCursorPos = p.CursorPos()
					ctx.Pos = p.Pos()

					gt.OnParse(ctx)
				}
			}
		}
	}

	return ctx.WipCodeGen
}

// A GenTrigger has a Token embeded so it can be used in a
// TokenLib if needed, along with other regular Tokens
type GenTrigger[T Token] interface {
	Token
	OnParse(*GenContext[T])
}

type GenContext[T Token] struct {
	WipCodeGen []CodeGen

	SourceFiles   []SourceFile
	SourceFilePos int // The index of the source file being parsed
	FileCursorPos int // The cursor will be directly to the right of the last token parsed when passed into gen()

	PositionedTokens []TokenWithCursorPos[T]
	Pos              int // Current position (index) in positionedTokens
}

// TODO: add .Authors and .OrigAuthor to track the SourceFile that created this
type CodeGen struct {
	Code       string
	OutputPath string
}

type SourceFile struct {
	Name    string
	Content string
}
