package main

import "golang.org/x/exp/utf8string"

func init() {
	funcMap["slice"] = &tmplFuncStruct{
		short:    "Construct a substring from a string.",
		examples: []string{`{{ "ᛁᚳ᛫ᛗᚨᚷ᛫ᚷᛚᚨᛋ᛫" | %s 3 7 }}`},
		fn: func(low int, high int, in string) string {
			s := utf8string.NewString(in)
			return s.Slice(low, high)
		},
	}
}
