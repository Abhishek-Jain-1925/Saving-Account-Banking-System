package enduser

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
)

var e dto.Error

// In below method Service Param is nothing but enduser.Service
func Login(userService Service) func(w http.ResponseWriter, r *http.Request) { //POST
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req dto.CreateLoginRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// log.Print("error !! while decoding login data from json into struct !!")
			e.Error_code = 500
			e.Error_msg = "error, while decoding data into json, plz provide valid credentails"
			_ = json.NewEncoder(w).Encode(e)
			return
		}

		err = req.Validate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("Plz, Provide Valid Credentials !!"))
			e.Error_code = 400
			e.Error_msg = err.Error()
			// w.Write([]byte("Plz,Provide Valid Data"))
			_ = json.NewEncoder(w).Encode(e)

			return
		}

		response, err := userService.CreateLogin(ctx, req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			e.Error_code = 400
			e.Error_msg = "Plz, provide valid credentials"
			// w.Write([]byte("Plz,Provide Valid Data"))
			_ = json.NewEncoder(w).Encode(e)
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
			w.WriteHeader(http.StatusInternalServerError)
			// log.Print("error !! while decoding Signup data from json into struct !!")
			e.Error_code = 500
			e.Error_msg = "error, while decoding data into json, plz provide valid credentails"
			_ = json.NewEncoder(w).Encode(e)
			return
		}

		err = req.ValidateUser()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// response := fmt.Sprintf("\nCAUTION : %v", err)
			// w.Write([]byte(response))
			e.Error_code = 400
			e.Error_msg = err.Error()
			_ = json.NewEncoder(w).Encode(e)
			return
		}

		result, err := userService.CreateSignup(ctx, req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			//log.Println(err)
			e.Error_code = 400
			e.Error_msg = err.Error()
			_ = json.NewEncoder(w).Encode(e)
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
			e.Error_code = 401
			e.Error_msg = "Plz, Do Login First"
			_ = json.NewEncoder(w).Encode(e)
			return
		}
		//w.Write([]byte(fmt.Sprintf("Hello, %s", response)))

		//Updating User Info
		var req dto.UpdateUser
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			e.Error_code = 500
			e.Error_msg = "error, while decoding data into json, plz provide valid credentails"
			_ = json.NewEncoder(w).Encode(e)
			return
		}

		err = req.ValidateUpdate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			e.Error_code = 400
			e.Error_msg = err.Error()
			_ = json.NewEncoder(w).Encode(e)
			return
		}

		result, err := userService.UpdateUser(ctx, req, user_id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			e.Error_code = 400
			e.Error_msg = err.Error()
			_ = json.NewEncoder(w).Encode(e)
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
