package rules


type Rule interface {
	 Name() string  
	 Validate(attribute string  , value any )bool
}