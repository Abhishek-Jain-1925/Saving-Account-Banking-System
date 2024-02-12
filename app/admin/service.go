package admin

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/repository"
	"github.com/dgrijalva/jwt-go"
)

type service struct {
	AdminRepo repository.AdminStorer
}
type Service interface {
	Authenticate(tknStr string) (response string, err error)
	ListUsers(ctx context.Context) ([]dto.Response, error)
	UpdateUser(ctx context.Context, req dto.UpdateUserInfo) (dto.UpdateUserInfo, error)
}

func NewService(AdminRepo repository.AdminStorer) Service {
	return &service{
		AdminRepo: AdminRepo,
	}
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

	if !strings.EqualFold(claims.Role, "admin") {
		return "", fmt.Errorf("access denied")
	}
	return claims.Username, nil
}

func (adm *service) ListUsers(ctx context.Context) ([]dto.Response, error) {
	tx, err := adm.AdminRepo.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	response, err := adm.AdminRepo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		txErr := adm.AdminRepo.HandleTransaction(ctx, tx, err)
		if txErr != nil {
			err = txErr
			return
		}
	}()
	return response, nil
}

func (us *service) UpdateUser(ctx context.Context, req dto.UpdateUserInfo) (dto.UpdateUserInfo, error) {
	tx, err := us.AdminRepo.BeginTx(ctx)
	if err != nil {
		return dto.UpdateUserInfo{}, fmt.Errorf(err.Error())
	}

	response, err := us.AdminRepo.UpdateUserInfo(req)
	if err != nil {
		return dto.UpdateUserInfo{}, err
	}

	defer func() {
		txErr := us.AdminRepo.HandleTransaction(ctx, tx, err)
		if txErr != nil {
			err = txErr
			return
		}
	}()
	return response, nil
}
