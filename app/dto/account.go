package dto

import (
	"errors"
	"fmt"
)

type CreateAccountReq struct {
	Account_no   int     `json:"acc_no,omitempty"`
	User_id      int     `json:"user_id"`
	Branch_id    int     `json:"branch_id"`
	Account_type string  `json:"acc_type"`
	Balance      float64 `json:"balance"`
}

type DeleteAccountReq struct {
	Account_no int `json:"acc_no"`
	User_id    int `json:"user_id"`
}

type Transaction struct {
	Account_no int     `json:"acc_no"`
	Amount     float64 `json:"amount"`
}

func (req *CreateAccountReq) Validate() error {
	if req.User_id <= 0 {
		return errors.New("user_id must be greater than 0")
	}
	if req.Branch_id <= 0 {
		return errors.New("branch_id must be greater than 0")
	}
	if req.Account_no < 0 {
		return errors.New("account_no cannot be negative")
	}
	if len(req.Account_type) <= 0 || (req.Account_type != "Savings") {
		return fmt.Errorf("please provide Valid Account type")
	}
	if req.Balance < 0 {
		return errors.New("balance cannot be negative")
	}
	return nil
}

func (req *DeleteAccountReq) ValidateDeleteReq() error {
	if req.Account_no <= 0 {
		return fmt.Errorf("please provide Valid Account No")
	}
	if req.User_id < 0 {
		return fmt.Errorf("please provide Valid User_ID")
	}
	return nil
}

func (req *Transaction) ValidateTransaction() error {

	if req.Account_no < 0 {
		return fmt.Errorf("account number never be negative")
	}
	if req.Amount < 0 {
		return fmt.Errorf("amount never be negative")
	}
	return nil
}
