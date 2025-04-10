package directive

import "github.com/EricFrancis12/geno"

func NewEngine(triggers ...geno.GenTrigger) *geno.GenEngine[geno.Token] {
	return geno.NewGenEngine(Lib(), triggers...)
}
