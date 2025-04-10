package handler

import (
	"github.com/EricFrancis12/geno"
	"github.com/EricFrancis12/geno/libs/base"
)

type BaseHandlerToken struct {
	base.BaseToken
	ParseHandlers []geno.ParseHandler
}

func NewHandlerBaseToken(parseHandlers ...geno.ParseHandler) BaseHandlerToken {
	return BaseHandlerToken{
		ParseHandlers: parseHandlers,
	}
}

func (t BaseHandlerToken) OnParse(ctx *geno.GenContext) {
	for _, ph := range t.ParseHandlers {
		if ph != nil {
			ph(ctx)
		}
	}
}
