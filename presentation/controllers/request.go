package controllers

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s is required of (type %s)", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}

	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}

	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}

	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role == "" || r.Company == "" || r.Location == "" || r.Link == "" || r.Salary <= 0 {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}
