package model

import "go_gql/validator"

func (r RegisterInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("email", r.Email)
	v.IsEmail("emai", r.Email)

	v.Required("password", r.Password)
	v.MinLength("password", r.Password, 6)

	v.Required("username", r.Username)
	v.MinLength("username", r.Username, 2)

	return v.IsValid(), v.Errors
}

func (l LoginInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("email", l.Email)
	v.IsEmail("emai", l.Email)

	v.Required("password", l.Password)
	v.MinLength("password", l.Password, 6)

	return v.IsValid(), v.Errors
}
