package main

import (
	"fmt"
	"mvc.com/src/http/validation/parser"
	"mvc.com/src/http/validation/validation"
)

func main() {
	data := map[string]any{
		"name":  "Ab",
		"email": nil,
	}

	v := validation.New(data)

	v.Add("name", parser.Parse("required|min:3")...)
	v.Add("email", parser.Parse("required")...)

	if !v.Validate() {
		fmt.Println(v.Errors())
	}

}
