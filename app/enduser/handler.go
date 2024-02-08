package enduser

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
)

// In below method Service Param is nothing but enduser.Service
func Login(userService Service) func(w http.ResponseWriter, r *http.Request) { //POST
	return func(w http.ResponseWriter, r *http.Request) {

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

		response, err := userService.CreateLogin(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error !! while creating token in service layer !! Plz.Provide Valid Data"))
			return
		}

		w.Write([]byte(response))
	}
}

func Signup(userService Service) func(w http.ResponseWriter, r *http.Request) { //Post
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.CreateUser

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Print("error !! while decoding login data from json into struct !!")
			return
		}

		err = req.ValidateUser()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error...while Validating input !! Plz, Provide Valid Credentials !!"))
			return
		}

		response, err := userService.CreateSignup(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte(response))
	}
}

func Update(w http.ResponseWriter, r *http.Request) { //Post

}

func List(w http.ResponseWriter, r *http.Request) { //Post

}

//updateEndUSer PUT pathparam
//list enduser GET queryParam
//getenduser admin/ GET pathparam
