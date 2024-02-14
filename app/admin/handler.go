package admin

import (
	"encoding/json"
	"net/http"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
)

func ListUsers(AdminService Service) func(w http.ResponseWriter, r *http.Request) { //GET
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tknStr := r.Header.Get("Authorization")
		_, err := AdminService.Authenticate(tknStr)
		if err != nil {
			dto.ErrorUnauthorizedAccess(err, w)
			return
		}

		resp, err := AdminService.ListUsers(ctx)
		if err != nil {
			dto.ErrorBadRequest(err, w)
			return
		}

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			dto.ErrorInternalServer(err, w)
			return
		}
	}
}

func Update(adminService Service) func(w http.ResponseWriter, r *http.Request) { //PUT
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tknStr := r.Header.Get("Authorization")

		_, err := adminService.Authenticate(tknStr)
		if err != nil {
			dto.ErrorUnauthorizedAccess(err, w)
			return
		}

		//Updating User Info
		var req dto.UpdateUserInfo
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			dto.ErrorInternalServer(err, w)
			return
		}

		err = req.ValidateUpdate()
		if err != nil {
			dto.ErrorBadRequest(err, w)
			return
		}

		result, err := adminService.UpdateUser(ctx, req)
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
