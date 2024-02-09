package dto

import "fmt"

type CreateAccountReq struct {
	User_id      int     `json:"user_id"`
	Branch_id    int     `json:"branch_id"`
	Account_type string  `json:"acc_type"`
	Balance      float64 `json:"balance"`
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
