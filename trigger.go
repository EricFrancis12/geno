package geno

// A GenTrigger has a Token embeded so it can be used in a
// TokenLib if needed, along with other regular Tokens
type GenTrigger interface {
	Token
	OnParse
}

type OnParse interface {
	OnParse(*GenContext)
}

type ParseHandler = func(ctx *GenContext)
