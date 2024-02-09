package account

import (
	"fmt"
	"os"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/repository"
	"github.com/dgrijalva/jwt-go"
)

type service struct {
	AccountRepo repository.AccountStorer
}

type Service interface {
	Authenticate(tknStr string) (response string, err error)
	CreateAccount(req dto.CreateAccountReq) (res string, err error)
}

func NewService(AccountRepo repository.AccountStorer) Service {
	return &service{
		AccountRepo: AccountRepo,
	}
}

// All Account related bussiness logic here onwards=>
func (us *service) Authenticate(tknStr string) (response string, err error) {

	jwtkey := []byte(os.Getenv("jwtkey"))
	claims := &dto.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	if err != nil {
		return "", fmt.Errorf("error in parsing claims")
	}

	if !tkn.Valid {
		return "", fmt.Errorf("invalid token")
	}
	return claims.Username, nil
}

func (as *service) CreateAccount(req dto.CreateAccountReq) (res string, err error) {

	response, err := as.AccountRepo.CreateAccount(req)
	if err != nil {
		return "", err
	}

	return response, nil
}
