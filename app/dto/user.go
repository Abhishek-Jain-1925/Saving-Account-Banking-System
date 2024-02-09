package dto

import (
	"fmt"
	"regexp"

	"github.com/dgrijalva/jwt-go"
)

type CreateLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type CreateUser struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	Role     string `json:"role"`
}

type UpdateUser struct {
	User_id  int    `json:"user_id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
}

func (req *CreateLoginRequest) Validate() error {
	if len(req.Username) <= 0 || len(req.Password) <= 0 {
		return fmt.Errorf("please provide anything as input")
	}
	if len(req.Password) < 3 && len(req.Password) > 16 {
		return fmt.Errorf("length of the password field must be in 3 to 16 chars")
	}
	return nil
}

func (req *CreateUser) ValidateUser() error {
	if len(req.Name) <= 0 {
		return fmt.Errorf("name field cannot be empty")
	}
	if len(req.Address) <= 0 {
		return fmt.Errorf("address field cannot be empty")
	}
	if len(req.Email) <= 0 {
		return fmt.Errorf("email field cannot be empty")
	}
	if !isValidEmail(req.Email) {
		return fmt.Errorf("invalid email format")
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
	if len(req.Role) <= 0 || (req.Role != "Customer" && req.Role != "Admin") {
		return fmt.Errorf("please provide a proper role field")
	}

	return nil
}

func isValidEmail(email string) bool {
	//provides a basic check
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(regex).MatchString(email)
}

func (req *UpdateUser) ValidateUpdate() error {
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
