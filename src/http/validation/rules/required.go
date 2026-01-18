package rules

type Required struct{}

func (r Required) Name() string {
	return "required"
}

func (r Required) Validate(attribute string, value any) bool {
	return value != nil
}