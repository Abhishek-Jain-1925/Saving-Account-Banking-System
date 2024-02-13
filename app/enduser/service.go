package enduser

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	UserRepo repository.UserStorer
}

// All User related funcs that processing From DB in Bussiness Logic
type Service interface {
	Authenticate(tknStr string) (user_id int, response string, err error)
	CreateLogin(ctx context.Context, req dto.CreateLoginRequest) (res string, err error)
	CreateSignup(ctx context.Context, req dto.CreateUser) (dto.Response, error)
	UpdateUser(ctx context.Context, req dto.UpdateUser, user_id int) (dto.UpdateUser, error)
}

func NewService(UserRepo repository.UserStorer) Service {
	return &service{
		UserRepo: UserRepo,
	}
}

// All Bussiness Logic Related funcs of USER with logic here onwards=>
func (us *service) Authenticate(tknStr string) (user_id int, response string, err error) {

	jwtkey := []byte(os.Getenv("jwtkey"))
	claims := &dto.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	if err != nil {
		return 0, "", fmt.Errorf("error in parsing claims")
	}

	if !tkn.Valid {
		return 0, "", fmt.Errorf("invalid token")
	}
	return claims.User_id, claims.Username, nil
}

func (us *service) CreateLogin(ctx context.Context, req dto.CreateLoginRequest) (res string, err error) {

	tx, err := us.UserRepo.BeginTx(ctx)
	LoginMap, err := us.UserRepo.GetLoginDetails()
	if err != nil {
		return "", err
	}
	var jwtkey = os.Getenv("jwtkey")

	expectedPwd, ok := LoginMap[req.Username]
	bcrErr := bcrypt.CompareHashAndPassword([]byte(expectedPwd), []byte(req.Password))
	if !ok || bcrErr != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	expirationTime := time.Now().Add(time.Minute * 10)
	//Getting Additional data from DB like user_id, role
	uid, role, err := us.UserRepo.TokenDetails(req.Username)

	claims := &dto.Claims{
		Username: req.Username,
		User_id:  uid,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(jwtkey))

	if err != nil {
		return "", fmt.Errorf("error in parse token, %s", err)
	}

	defer func() {
		txErr := us.UserRepo.HandleTransaction(ctx, tx, err)
		if txErr != nil {
			err = txErr
			return
		}
	}()
	return tokenStr, nil
}

func (us *service) CreateSignup(ctx context.Context, req dto.CreateUser) (dto.Response, error) {
	tx, err := us.UserRepo.BeginTx(ctx)
	if err != nil {
		return dto.Response{}, err
	}

	response, err := us.UserRepo.AddUser(req)
	if err != nil {
		return dto.Response{}, err
	}

	defer func() {
		txErr := us.UserRepo.HandleTransaction(ctx, tx, err)
		if txErr != nil {
			err = txErr
			return
		}
	}()
	return response, nil
}

func (us *service) UpdateUser(ctx context.Context, req dto.UpdateUser, user_id int) (dto.UpdateUser, error) {
	tx, err := us.UserRepo.BeginTx(ctx)
	if err != nil {
		return dto.UpdateUser{}, err
	}

	response, err := us.UserRepo.UpdateUser(req, user_id)
	if err != nil {
		return dto.UpdateUser{}, err
	}

	defer func() {
		txErr := us.UserRepo.HandleTransaction(ctx, tx, err)
		if txErr != nil {
			err = txErr
			return
		}
	}()
	return response, nil
}
