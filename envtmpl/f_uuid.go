package main

import "github.com/google/uuid"

func init() {
	funcMap["uuid"] = &tmplFuncStruct{
		short: "Create a random (v4) UUID.",
		examples: []string{
			`{{ %s }}`,
		},
		fn: uuid.New,
	}
}
