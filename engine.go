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

func (e GenEngine[T]) Gen(sourceFiles ...SourceFile) GenResult {
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

				if tk, err := p.Parse(gt); err != nil {
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

	return GenResult{
		CodeGens: ctx.WipCodeGen,
	}
}

type SourceFile struct {
	Name    string
	Content string
}

type GenContext struct {
	WipCodeGen []CodeGen

	SourceFiles   []SourceFile
	SourceFilePos int // The index of the source file being parsed
	FileCursorPos int // The cursor will be directly to the right of the last token parsed when passed into gen()

	TokensFromSource []TokenFromSource[Token]
	Pos              int // Current position (index) in positionedTokens
}

// TODO: Add Authors and OrigAuthor props to track which SourceFile created this
type CodeGen struct {
	Code       string
	OutputPath string
}

type GenResult struct {
	CodeGens []CodeGen
}

func (g GenResult) Join(sep string) string {
	s := ""
	for _, cg := range g.CodeGens {
		s += cg.Code
		s += sep
	}
	return s
}
