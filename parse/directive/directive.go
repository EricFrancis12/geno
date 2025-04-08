package directive

import "strings"

type Directive struct {
	Name   string
	Params []string
}

// This function extracts Directives from comments that use the following format:
// // #[foo, bar(baz)]
//
// This comment would create 2 directives:
// 1. Directive{Name: "foo", Params: []string{}}
// 2. Directive{Name: "bar", Params: []string{"baz"}}
func ParseCommentDirectives(s string) []Directive {
	directives := []Directive{}

	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "//#") || strings.HasPrefix(s, "// #") {
		s = strings.TrimPrefix(s, "//")
		s = strings.TrimSpace(s)
		s = strings.TrimPrefix(s, "#[")

		// Go char by char to find the closing bracket
		// and trim the string to the closing bracket
		for i := 0; i < len(s); i++ {
			if s[i] == ']' {
				s = s[:i]
				break
			}
		}

		s = strings.TrimSpace(s)

		directive := Directive{}
		parts := strings.SplitN(s, "(", 2)
		directive.Name = strings.TrimSpace(parts[0])
		if len(parts) > 1 {
			paramsStr := strings.TrimSuffix(parts[1], ")")
			directive.Params = strings.Split(paramsStr, ",")
			for i := range directive.Params {
				directive.Params[i] = strings.TrimSpace(directive.Params[i])
			}
		}
		directives = append(directives, directive)
	}

	return directives
}
