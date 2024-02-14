package enduser

import (
	"encoding/json"
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
			dto.ErrorInternalServer(err, w)
			return
		}

		err = req.Validate()
		if err != nil {
			dto.ErrorBadRequest(err, w)
			return
		}

		response, err := userService.CreateLogin(ctx, req)
		if err != nil {
			dto.ErrorBadRequest(err, w)
			return
		}

		res := dto.LoginToken{
			IssuedToken: response,
		}

		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			dto.ErrorInternalServer(err, w)
			return
		}
	}
}

func Signup(userService Service) func(w http.ResponseWriter, r *http.Request) { //Post
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req dto.CreateUser

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			dto.ErrorInternalServer(err, w)
			return
		}

		err = req.ValidateUser()
		if err != nil {
			dto.ErrorBadRequest(err, w)
			return
		}

		result, err := userService.CreateSignup(ctx, req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
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

func Update(userService Service) func(w http.ResponseWriter, r *http.Request) { //PUT
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tknStr := r.Header.Get("Authorization")

		user_id, _, err := userService.Authenticate(tknStr)
		if err != nil {
			dto.ErrorUnauthorizedAccess(err, w)
			return
		}

		//Updating User Info
		var req dto.UpdateUser
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			dto.ErrorBadRequest(err, w)
			return
		}

		err = req.ValidateUpdate()
		if err != nil {
			dto.ErrorBadRequest(err, w)
			return
		}

		result, err := userService.UpdateUser(ctx, req, user_id)
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
