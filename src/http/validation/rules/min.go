package rules

import "fmt"

type Min struct {
    Min int
}

func (m Min) Name() string {
    return "min"
}

func (m Min) Validate(field string, value any, _ map[string]any) error {
    s, ok := value.(string)
    if !ok || len(s) < m.Min {
        return fmt.Errorf("%s must be at least %d characters", field, m.Min)
    }
    return nil
}
