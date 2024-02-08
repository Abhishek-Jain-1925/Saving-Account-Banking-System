package dto

import (
	"fmt"

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
	return nil
}
