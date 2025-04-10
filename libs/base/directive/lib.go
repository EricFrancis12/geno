package directive

import "github.com/EricFrancis12/geno/libs/base/custom"

func Lib() custom.CustomBaseTokenLib {
	lib := custom.CustomBaseTokenLib{}
	lib.AddToken(CommentDirective{})
	return lib
}
