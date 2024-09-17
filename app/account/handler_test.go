package account_test

import (
	"context"
	"errors"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
)

type mockAccountService struct{}

func (m *mockAccountService) Authenticate(token string) (int, string, error) {
	if token == "valid_token" {
		return 1, "user_role", nil
	}
	return 0, "", errors.New("invalid token")
}

func (m *mockAccountService) CreateAccount(ctx context.Context, req dto.CreateAccountReq, userID int) (dto.CreateAccountReq, error) {
	return dto.CreateAccountReq{}, nil
}

func (m *mockAccountService) DepositMoney(ctx context.Context, req dto.Transaction, userID int) (dto.TransactionResponse, error) {
	return dto.TransactionResponse{}, nil
}

func (m *mockAccountService) WithdrawalMoney(ctx context.Context, req dto.Transaction, userID int) (dto.TransactionResponse, error) {
	return dto.TransactionResponse{}, nil
}

func (m *mockAccountService) DeleteAccount(ctx context.Context, req dto.DeleteAccountReq, userID int) (dto.DeleteAccount, error) {
	return dto.DeleteAccount{}, nil
}


//Deposit, Withdrawal, and Delete.
