package enduser

import (
	"fmt"
	"os"
	"time"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/repository"
	"github.com/dgrijalva/jwt-go"
)

type service struct {
	UserRepo repository.UserStorer
}

var Map1 = map[string]string{
	"Samnit": "Sam@123",
	"user2":  "password2",
}

// All User related funcs that processing From DB in Bussiness Logic
type Service interface {
	Authenticate(tknStr string) (response string, err error)
	CreateLogin(req dto.CreateLoginRequest) (res string, err error)
	CreateSignup(req dto.CreateUser) (res string, err error)
	UpdateUser(req dto.UpdateUser) (response string, err error)
}

func NewService(UserRepo repository.UserStorer) Service {
	return &service{
		UserRepo: UserRepo,
	}
}

// All Bussiness Logic Related funcs of USER with logic here onwards=>
func (us *service) CreateLogin(req dto.CreateLoginRequest) (res string, err error) {

	LoginMap, err := us.UserRepo.GetLoginDetails()
	if err != nil {
		return "", err
	}

	var jwtkey = os.Getenv("jwtkey")

	expectedPwd, ok := LoginMap[req.Username]
	if !ok || expectedPwd != req.Password {
		return "", fmt.Errorf("incorrect password")
	}

	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &dto.Claims{
		Username: req.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(jwtkey))
	if err != nil {

		return "", fmt.Errorf("error in parse token, %s", err)
	}

	return tokenStr, nil
}

func (us *service) CreateSignup(req dto.CreateUser) (res string, err error) {

	response, err := us.UserRepo.AddUser(req)
	if err != nil {
		return "", err
	}
	return response, nil
}

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

func (us *service) UpdateUser(req dto.UpdateUser) (response string, err error) {
	response, err = us.UserRepo.UpdateUser(req)
	if err != nil {
		return "", err
	}

	return response, nil
}
