package validation

type Error struct {
	Field   string
	Rule    string
	Message string
}

func (e Error) Error() string {
	return e.Message
}	
