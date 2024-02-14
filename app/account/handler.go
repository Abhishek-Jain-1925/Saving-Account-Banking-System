package account

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
)

func Create(accService Service) func(w http.ResponseWriter, r *http.Request) { //Post
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tknStr := r.Header.Get("Authorization")
		user_id, _, err := accService.Authenticate(tknStr)
		if err != nil {
			dto.ErrorUnauthorizedAccess(err, w)
			return
		}

		var req dto.CreateAccountReq
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			dto.ErrorInternalServer(err, w)
			return
		}

		err = req.Validate()
		if err != nil {
			dto.ErrorBadRequest(err, w)
			return
		}

		result, err := accService.CreateAccount(ctx, req, user_id)
		if err != nil {
			dto.ErrorBadRequest(err, w)
			return
		}

		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			dto.ErrorInternalServer(err, w)
			return
		}
	}
}

func Deposit(accService Service) func(w http.ResponseWriter, r *http.Request) { //Post
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tknStr := r.Header.Get("Authorization")
		user_id, _, err := accService.Authenticate(tknStr)
		if err != nil {
			dto.ErrorUnauthorizedAccess(err, w)
			return
		}

		var req dto.Transaction
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			dto.ErrorInternalServer(err, w)
			return
		}

		err = req.ValidateTransaction()
		if err != nil {
			dto.ErrorBadRequest(err, w)
			return
		}

		//var result dto.Transaction
		result, err := accService.DepositMoney(ctx, req, user_id)
		if err != nil {
			dto.ErrorBadRequest(err, w)
			return
		}

		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			dto.ErrorInternalServer(err, w)
			return
		}
	}
}

func Withdrawal(accService Service) func(w http.ResponseWriter, r *http.Request) { //PUT
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tknStr := r.Header.Get("Authorization")
		user_id, _, err := accService.Authenticate(tknStr)
		if err != nil {
			dto.ErrorUnauthorizedAccess(err, w)
			return
		}

		var req dto.Transaction
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			dto.ErrorInternalServer(err, w)
			return
		}

		err = req.ValidateTransaction()
		if err != nil {
			dto.ErrorBadRequest(err, w)
			return
		}

		result, err := accService.WithdrawalMoney(ctx, req, user_id)
		if err != nil {
			dto.ErrorBadRequest(err, w)
			return
		}

		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			dto.ErrorInternalServer(err, w)
			return
		}
	}
}

func Delete(accService Service) func(w http.ResponseWriter, r *http.Request) { //Post
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tknStr := r.Header.Get("Authorization")
		user_id, _, err := accService.Authenticate(tknStr)
		if err != nil {
			dto.ErrorUnauthorizedAccess(err, w)
			return
		}

		var req dto.DeleteAccountReq
		// err = json.NewDecoder(r.Body).Decode(&req)
		queryParams := r.URL.Query()
		paramValue := queryParams.Get("acc_no")
		acc_no, err := strconv.Atoi(paramValue)
		req.Account_no = acc_no

		if err != nil {
			dto.ErrorInternalServer(err, w)
			return
		}

		err = req.ValidateDeleteReq()
		if err != nil {
			dto.ErrorBadRequest(err, w)
			return
		}

		response, err := accService.DeleteAccount(ctx, req, user_id)
		if err != nil {
			dto.ErrorBadRequest(err, w)
			return
		}

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			dto.ErrorInternalServer(err, w)
			return
		}
	}
}

func ViewBalance(userService Service) func(w http.ResponseWriter, r *http.Request) { //GET pathparam
	return func(w http.ResponseWriter, r *http.Request) {
		//ctx := r.Context()

		//tknStr := r.Header.Get("Authorization")
		// user_id, response, err := userService.Authenticate(tknStr)
		// if err != nil {
		// 	dto.ErrorUnauthorizedAccess(err, w)
		// 	return
		// }
	}
}
