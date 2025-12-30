package rules

import "fmt"

type Required struct{}

func (r Required) Name() string {
    return "required"
}

func (r Required) Validate(field string, value any, _ map[string]any) error {
    if value == nil {
        return fmt.Errorf("%s is required", field)
    }
    return nil
}
