package validation

type Rule interface {
	Validate(field string, value any, data map[string]any) error
	Name() string
}
