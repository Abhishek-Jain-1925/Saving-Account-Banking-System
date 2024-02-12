package enduser

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
)

// In below method Service Param is nothing but enduser.Service
func Login(userService Service) func(w http.ResponseWriter, r *http.Request) { //POST
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req dto.CreateLoginRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Print("error !! while decoding login data from json into struct !!")
			return
		}

		err = req.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error...while Validating input !! Plz, Provide Valid Credentials !!"))
			return
		}

		response, err := userService.CreateLogin(ctx, req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			w.Write([]byte("Plz,Provide Valid Data"))
			return
		}

		resStr := "*** Logged in successfully ***"
		resStr += fmt.Sprintf("\n\n Your Token for further Banking : \n%v", response)
		w.Write([]byte(resStr))
	}
}

func Signup(userService Service) func(w http.ResponseWriter, r *http.Request) { //Post
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req dto.CreateUser

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Print("error !! while decoding Signup data from json into struct !!")
			return
		}

		err = req.ValidateUser()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := fmt.Sprintf("Plz, Provide Valid Credentials !! \nCAUTION : %v", err)
			w.Write([]byte(response))
			return
		}

		result, err := userService.CreateSignup(ctx, req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		// w.Write([]byte(response))
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func Update(userService Service) func(w http.ResponseWriter, r *http.Request) { //PUT
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tknStr := r.Header.Get("Authorization")

		user_id, _, err := userService.Authenticate(tknStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Plz, Do Login First !!"))
			log.Print(err)
			return
		}
		//w.Write([]byte(fmt.Sprintf("Hello, %s", response)))

		//Updating User Info
		var req dto.UpdateUser
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Print("error !! while decoding Update data from json into struct !!")
			return
		}

		err = req.ValidateUpdate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := fmt.Sprintf("Plz, Provide Valid Credentials !! \nCAUTION : %v", err)
			w.Write([]byte(response))
			return
		}

		result, err := userService.UpdateUser(ctx, req, user_id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// w.Write([]byte(response))
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func List(w http.ResponseWriter, r *http.Request) { //Post

}

//updateEndUSer PUT pathparam
//list enduser GET queryParam
//getenduser admin/ GET pathparam
