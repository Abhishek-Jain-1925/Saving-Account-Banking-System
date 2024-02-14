package account_test

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/account"
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

func TestCreateHandler(t *testing.T) {
	service := &mockAccountService{}
	handler := account.Create(service)

	t.Run("Valid Request", func(t *testing.T) {
		payload := []byte(`{"account_number":"1234567890","initial_balance":1000}`)
		req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "valid_token")
		rr := httptest.NewRecorder()
		handler(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusOK)
		}
	})

	t.Run("Invalid Token", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/create", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "invalid_token")
		rr := httptest.NewRecorder()
		handler(rr, req)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusUnauthorized)
		}
	})
}

//Deposit, Withdrawal, and Delete.
