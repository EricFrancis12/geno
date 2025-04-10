package directive

import "github.com/EricFrancis12/geno/libs/base"

func Lib() base.CustomBaseTokenLib {
	lib := base.CustomBaseTokenLib{}
	lib.AddToken(CommentDirective{})
	return lib
}
