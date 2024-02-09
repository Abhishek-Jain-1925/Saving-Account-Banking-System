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

		tknStr := r.Header.Get("Authorization")
		response, err := accService.Authenticate(tknStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Plz, Do Login First !!"))
			log.Print(err)
			return
		}
		w.Write([]byte(fmt.Sprintf("Hello, %s", response)))

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

		response, err = accService.CreateAccount(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte(response))
	}
}

func Deposit(w http.ResponseWriter, r *http.Request) { //PUT pathparam

}

func Withdrawal(w http.ResponseWriter, r *http.Request) { //PUT pathparam

}

func DeleteAccount(w http.ResponseWriter, r *http.Request) { //Delete

}

func ViewStatement(w http.ResponseWriter, r *http.Request) { //GET pathparam

}

//Domain => struct and its related vars, const, all data rather than bussiness logic i.e.Service
//rollback of db transaction
//
