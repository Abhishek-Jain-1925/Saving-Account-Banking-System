package account

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
)

func Create(accService Service) func(w http.ResponseWriter, r *http.Request) { //Post
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tknStr := r.Header.Get("Authorization")
		user_id, _, err := accService.Authenticate(tknStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Plz, Do Login First !!"))
			log.Print(err)
			return
		}
		//w.Write([]byte(fmt.Sprintf("Hello, %s", response)))

		var req dto.CreateAccountReq
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Print("error !! while creating Account data from json into struct !!")
			return
		}

		err = req.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("\n Error...while Validating input !! Plz, Provide Valid Credentials !!"))
			return
		}

		result, err := accService.CreateAccount(ctx, req, user_id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}

		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
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
			w.Write([]byte("Plz, Do Login First !!"))
			log.Print(err)
			return
		}
		// w.Write([]byte(fmt.Sprintf("Hello, %s", response)))

		var req dto.Transaction
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Print("error !! while creating Account data from json into struct !!")
			return
		}

		err = req.ValidateTransaction()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("\n Error...while Validating input !! Plz, Provide Valid Credentials !!"))
			return
		}

		//var result dto.Transaction
		result, err := accService.DepositMoney(ctx, req, user_id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res := fmt.Sprintf("\nCAUTION : %v", err)
			w.Write([]byte(res))
			return
		}

		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func Withdrawal(accService Service) func(w http.ResponseWriter, r *http.Request) { //PUT
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tknStr := r.Header.Get("Authorization")
		user_id, response, err := accService.Authenticate(tknStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Plz, Do Login First !!"))
			log.Print(err)
			return
		}
		w.Write([]byte(fmt.Sprintf("Hello, %s", response)))

		var req dto.Transaction
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Print("error !! while creating Account data from json into struct !!")
			return
		}

		err = req.ValidateTransaction()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("\n Error...while Validating input !! Plz, Provide Valid Credentials !!"))
			return
		}

		result, err := accService.WithdrawalMoney(ctx, req, user_id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res := fmt.Sprintf("\nCAUTION : %v", err)
			w.Write([]byte(res))
			return
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		// w.Write([]byte(response))
	}
}

func Delete(accService Service) func(w http.ResponseWriter, r *http.Request) { //Post
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tknStr := r.Header.Get("Authorization")
		user_id, response, err := accService.Authenticate(tknStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Plz, Do Login First !!"))
			log.Print(err)
			return
		}
		w.Write([]byte(fmt.Sprintf("Hello, %s", response)))

		var req dto.DeleteAccountReq
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Print("error !! while creating Account data from json into struct !!")
			return
		}

		err = req.ValidateDeleteReq()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("\n Error...while Validating input !! Plz, Provide Valid Credentials !!"))
			return
		}
		response, err = accService.DeleteAccount(ctx, req, user_id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := fmt.Sprintf("\nCAUTION : %v", err)
			w.Write([]byte(response))
			return
		}
		w.Write([]byte(response))
		w.Write([]byte("\n\n Thank You For Banking With Us !! \n We are looking forward for your feedback to improve our Banking Facility !!"))
	}
}

func ViewStatement(w http.ResponseWriter, r *http.Request) { //GET pathparam

}

//Domain => struct and its related vars, const, all data rather than bussiness logic i.e.Service
//rollback of db transaction
//
