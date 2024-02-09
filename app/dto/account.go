package dto

import "fmt"

type CreateAccountReq struct {
	User_id      int     `json:"user_id"`
	Branch_id    int     `json:"branch_id"`
	Account_type string  `json:"acc_type"`
	Balance      float64 `json:"balance"`
}

type DeleteAccountReq struct {
	Account_no int `json:"acc_no"`
	User_id    int `json:"user_id"`
}

func (req *CreateAccountReq) Validate() error {
	if len(req.Account_type) <= 0 || (req.Account_type != "Savings") {
		return fmt.Errorf("please provide Valid Account type")
	}
	if req.Balance < 0 {
		return fmt.Errorf("initial account balance should be greater than equal Zero")
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
