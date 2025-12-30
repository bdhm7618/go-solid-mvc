package factory

import (
	"strconv"

	"mvc.com/src/http/validation/rules"
	"mvc.com/src/http/validation/validation"
)

type RuleBuilder func(args ...string) validation.Rule

var registry = map[string]RuleBuilder{}

func Register(name string, builder RuleBuilder) {
	registry[name] = builder
}

func Create(name string, args ...string) validation.Rule {
	if builder, ok := registry[name]; ok {
		return builder(args...)
	}
	panic("unknown rule: " + name)
}
func init() {
	Register("required", func(_ ...string) validation.Rule {
		return rules.Required{}
	})

	Register("min", func(args ...string) validation.Rule {
		min, _ := strconv.Atoi(args[0])
		return rules.Min{Min: min}
	})
}
