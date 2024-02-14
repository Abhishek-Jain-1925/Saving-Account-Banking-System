package account

import (
	"context"
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
	Authenticate(tknStr string) (user_id int, response string, err error)
	CreateAccount(ctx context.Context, req dto.CreateAccountReq, user_id int) (dto.CreateAccountReq, error)
	DeleteAccount(ctx context.Context, req dto.DeleteAccountReq, user_id int) (dto.DeleteAccount, error)
	DepositMoney(ctx context.Context, req dto.Transaction, user_id int) (dto.TransactionResponse, error)
	WithdrawalMoney(ctx context.Context, req dto.Transaction, user_id int) (dto.TransactionResponse, error)
}

func NewService(AccountRepo repository.AccountStorer) Service {
	return &service{
		AccountRepo: AccountRepo,
	}
}

// All Account related bussiness logic here onwards=>
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

func (as *service) CreateAccount(ctx context.Context, req dto.CreateAccountReq, user_id int) (dto.CreateAccountReq, error) {

	tx, _ := as.AccountRepo.BeginTx(ctx)

	response, err := as.AccountRepo.CreateAccount(req, user_id)
	if err != nil {
		return dto.CreateAccountReq{}, err
	}

	defer func() {
		txErr := as.AccountRepo.HandleTransaction(ctx, tx, err)
		if txErr != nil {
			err = txErr
			return
		}
	}()
	return response, nil
}

func (as *service) DeleteAccount(ctx context.Context, req dto.DeleteAccountReq, user_id int) (dto.DeleteAccount, error) {

	tx, _ := as.AccountRepo.BeginTx(ctx)

	response, err := as.AccountRepo.DeleteAccount(req, user_id)
	if err != nil {
		return dto.DeleteAccount{}, err
	}

	defer func() {
		txErr := as.AccountRepo.HandleTransaction(ctx, tx, err)
		if txErr != nil {
			err = txErr
			return
		}
	}()
	return response, nil
}

func (as *service) DepositMoney(ctx context.Context, req dto.Transaction, user_id int) (dto.TransactionResponse, error) {

	tx, _ := as.AccountRepo.BeginTx(ctx)

	response, err := as.AccountRepo.DepositMoney(req, user_id)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	defer func() {
		txErr := as.AccountRepo.HandleTransaction(ctx, tx, err)
		if txErr != nil {
			err = txErr
			return
		}
	}()
	return response, nil
}

func (as *service) WithdrawalMoney(ctx context.Context, req dto.Transaction, user_id int) (dto.TransactionResponse, error) {

	tx, _ := as.AccountRepo.BeginTx(ctx)

	response, err := as.AccountRepo.WithdrawalMoney(req, user_id)
	if err != nil {
		return dto.TransactionResponse{}, err
	}

	defer func() {
		txErr := as.AccountRepo.HandleTransaction(ctx, tx, err)
		if txErr != nil {
			err = txErr
			return
		}
	}()
	return response, nil
}
