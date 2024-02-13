package account

import (
	"encoding/json"
	"net/http"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
)

var e dto.Error

func Create(accService Service) func(w http.ResponseWriter, r *http.Request) { //Post
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tknStr := r.Header.Get("Authorization")
		user_id, _, err := accService.Authenticate(tknStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			e.Error_code = 401
			e.Error_msg = "Plz, Do Login First"
			_ = json.NewEncoder(w).Encode(e)
			return
		}
		//w.Write([]byte(fmt.Sprintf("Hello, %s", response)))

		var req dto.CreateAccountReq
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			e.Error_code = 500
			e.Error_msg = "error, while decoding data into json, plz provide valid credentails"
			_ = json.NewEncoder(w).Encode(e)
			return
		}

		err = req.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("\n Error...while Validating input !! Plz, Provide Valid Credentials !!"))
			e.Error_code = 400
			e.Error_msg = err.Error()
			_ = json.NewEncoder(w).Encode(e)
			return
		}

		result, err := accService.CreateAccount(ctx, req, user_id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// log.Println(err)
			e.Error_code = 400
			e.Error_msg = err.Error()
			_ = json.NewEncoder(w).Encode(e)
			return
		}

		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			e.Error_code = 500
			e.Error_msg = "error, while encoding data from struct"
			_ = json.NewEncoder(w).Encode(e)
		}
		// w.Write([]byte(response))
	}
}

func Deposit(accService Service) func(w http.ResponseWriter, r *http.Request) { //Post
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tknStr := r.Header.Get("Authorization")
		user_id, _, err := accService.Authenticate(tknStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			e.Error_code = 401
			e.Error_msg = "Plz, Do Login First"
			_ = json.NewEncoder(w).Encode(e)
			return
		}
		// w.Write([]byte(fmt.Sprintf("Hello, %s", response)))

		var req dto.Transaction
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			e.Error_code = 500
			e.Error_msg = "error, while decoding data into json, plz provide valid credentails"
			_ = json.NewEncoder(w).Encode(e)
			return
		}

		err = req.ValidateTransaction()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("\n Error...while Validating input !! Plz, Provide Valid Credentials !!"))
			e.Error_code = 400
			e.Error_msg = err.Error()
			_ = json.NewEncoder(w).Encode(e)
			return
		}

		//var result dto.Transaction
		result, err := accService.DepositMoney(ctx, req, user_id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// res := fmt.Sprintf("\nCAUTION : %v", err)
			// w.Write([]byte(res))
			e.Error_code = 400
			e.Error_msg = err.Error()
			_ = json.NewEncoder(w).Encode(e)
			return
		}

		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			e.Error_code = 500
			e.Error_msg = "error, while decoding data from struct"
			_ = json.NewEncoder(w).Encode(e)
		}
	}
}

func Withdrawal(accService Service) func(w http.ResponseWriter, r *http.Request) { //PUT
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tknStr := r.Header.Get("Authorization")
		user_id, _, err := accService.Authenticate(tknStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			e.Error_code = 401
			e.Error_msg = "Plz, Do Login First"
			_ = json.NewEncoder(w).Encode(e)
			return
		}
		//w.Write([]byte(fmt.Sprintf("Hello, %s", response)))

		var req dto.Transaction
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			e.Error_code = 500
			e.Error_msg = "error, while decoding data into json, plz provide valid credentails"
			_ = json.NewEncoder(w).Encode(e)
			return
		}

		err = req.ValidateTransaction()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			e.Error_code = 400
			e.Error_msg = err.Error()
			_ = json.NewEncoder(w).Encode(e)
			return
		}

		result, err := accService.WithdrawalMoney(ctx, req, user_id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// res := fmt.Sprintf("\nCAUTION : %v", err)
			// w.Write([]byte(res))
			e.Error_code = 400
			e.Error_msg = err.Error()
			_ = json.NewEncoder(w).Encode(e)
			return
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			e.Error_code = 500
			e.Error_msg = "error, while decoding data from struct"
			_ = json.NewEncoder(w).Encode(e)
		}
		// w.Write([]byte(response))
	}
}

func Delete(accService Service) func(w http.ResponseWriter, r *http.Request) { //Post
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tknStr := r.Header.Get("Authorization")
		user_id, _, err := accService.Authenticate(tknStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			e.Error_code = 401
			e.Error_msg = "Plz, Do Login First"
			_ = json.NewEncoder(w).Encode(e)
			return
		}
		//w.Write([]byte(fmt.Sprintf("Hello, %s", response)))

		var req dto.DeleteAccountReq
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			e.Error_code = 500
			e.Error_msg = "error, while decoding data into json, plz provide valid credentails"
			_ = json.NewEncoder(w).Encode(e)
			return
		}

		err = req.ValidateDeleteReq()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			e.Error_code = 400
			e.Error_msg = err.Error()
			_ = json.NewEncoder(w).Encode(e)
			return
		}
		response, err := accService.DeleteAccount(ctx, req, user_id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			e.Error_code = 400
			e.Error_msg = err.Error()
			_ = json.NewEncoder(w).Encode(e)
			return
		}
		w.Write([]byte(response))
		w.Write([]byte("\n\n Thank You For Banking With Us !! \n We are looking forward for your feedback to improve our Banking Facility !!"))
	}
}

func ViewStatement(w http.ResponseWriter, r *http.Request) { //GET pathparam

}
