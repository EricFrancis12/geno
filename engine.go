package geno

type GenEngine[T Token] struct {
	TokenLib TokenLib[T]
	Triggers []GenTrigger
}

func NewGenEngine[T Token](tokenLib TokenLib[T], triggers ...GenTrigger) *GenEngine[T] {
	return &GenEngine[T]{
		TokenLib: tokenLib,
		Triggers: triggers,
	}
}

func (e *GenEngine[T]) AddTrigger(trigger GenTrigger) {
	e.Triggers = append(e.Triggers, trigger)
}

func (e *GenEngine[T]) AddTriggers(triggers ...GenTrigger) {
	for _, trigger := range triggers {
		e.AddTrigger(trigger)
	}
}

func (e GenEngine[T]) Gen(sourceFiles ...SourceFile) []CodeGen {
	ctx := &GenContext{
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
			// ctx.TokensFromSource = p.TokensFromSource
			for _, tfs := range p.TokensFromSource {
				ctx.TokensFromSource = append(ctx.TokensFromSource, tfs.Generalize())
			}

			for !p.AtEOF() {
				posBefore := p.Pos()

				tp := p.Generalize()

				if tk, err := gt.Parse(tp); err != nil {
					// Reset the parser to the last position + 1 to advance to the next token
					p.SetPos(posBefore + 1)
				} else {
					ctx.FileCursorPos = p.CursorPos()
					ctx.Pos = p.Pos()

					// Check and run on parse effect if present
					op, ok := tk.(OnParse)
					if ok {
						op.OnParse(ctx)
					}
				}
			}
		}
	}

	return ctx.WipCodeGen
}

// A GenTrigger has a Token embeded so it can be used in a
// TokenLib if needed, along with other regular Tokens
type GenTrigger interface {
	Token
	OnParse
}

type OnParse interface {
	OnParse(*GenContext)
}

type GenContext struct {
	WipCodeGen []CodeGen

	SourceFiles   []SourceFile
	SourceFilePos int // The index of the source file being parsed
	FileCursorPos int // The cursor will be directly to the right of the last token parsed when passed into gen()

	TokensFromSource []TokenFromSource[Token]
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
