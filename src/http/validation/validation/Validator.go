package validation

type Validator struct {
	data  map[string]any
	rules map[string][]Rule
	errs  map[string][]Error
}

func New(data map[string]any) *Validator {
	return &Validator{
		data:  data,
		rules: make(map[string][]Rule),
		errs:  make(map[string][]Error),
	}
}

// THIS MUST EXIST
func (v *Validator) Add(field string, rules ...Rule) {
	v.rules[field] = append(v.rules[field], rules...)
}

func (v *Validator) Validate() bool {
	for field, rules := range v.rules {
		value := v.data[field]

		for _, rule := range rules {
			if err := rule.Validate(field, value, v.data); err != nil {
				v.errs[field] = append(v.errs[field], Error{
					Field:   field,
					Rule:    rule.Name(),
					Message: err.Error(),
				})
			}
		}
	}
	return len(v.errs) == 0
}

func (v *Validator) Errors() map[string][]Error {
	return v.errs
}
