package validator

type Validator struct {
	Errors map[string]string
}

type Validation interface {
	Validate() (bool, map[string]string)
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}
