package parser

import (
	"strings"

	"mvc.com/src/http/validation/factory"
	"mvc.com/src/http/validation/validation"
)

func Parse(ruleString string) []validation.Rule {
    parts := strings.Split(ruleString, "|")
    var rules []validation.Rule

    for _, part := range parts {
        tokens := strings.Split(part, ":")
        rules = append(rules, factory.Create(tokens[0], tokens[1:]...))
    }

    return rules
}
