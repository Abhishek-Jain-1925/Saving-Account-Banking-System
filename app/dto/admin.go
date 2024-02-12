package dto

import "fmt"

type UpdateUserInfo struct {
	User_id  int    `json:"user_id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Role     string `json:"role"`
}

func (req *UpdateUserInfo) ValidateUpdate() error {
	if len(req.Name) <= 0 {
		return fmt.Errorf("name field cannot be empty")
	}
	if len(req.Address) <= 0 {
		return fmt.Errorf("address field cannot be empty")
	}
	if len(req.Password) <= 0 {
		return fmt.Errorf("password field cannot be empty")
	}
	if len(req.Password) < 3 || len(req.Password) > 16 {
		return fmt.Errorf("length of the password field must be between 3 and 16 characters")
	}
	if len(req.Mobile) <= 0 {
		return fmt.Errorf("mobile field cannot be empty")
	}
	return nil
}
