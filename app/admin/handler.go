package admin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
)

func ListUsers(AdminService Service) func(w http.ResponseWriter, r *http.Request) { //GET
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tknStr := r.Header.Get("Authorization")
		_, err := AdminService.Authenticate(tknStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(string(err.Error())))
			return
		}
		//w.Write([]byte(fmt.Sprintf("Hello, %s", response)))

		resp, err := AdminService.ListUsers(ctx)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Plz,Provide Valid Data"))
			return
		}
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func Update(adminService Service) func(w http.ResponseWriter, r *http.Request) { //PUT
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tknStr := r.Header.Get("Authorization")

		response, err := adminService.Authenticate(tknStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(fmt.Sprintf("Hello, %s", response)))

		//Updating User Info
		var req dto.UpdateUserInfo
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Print("error !! while decoding Update data from json into struct !!")
			return
		}

		err = req.ValidateUpdate()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := fmt.Sprintf("\nCAUTION : %v", err)
			w.Write([]byte(response))
			return
		}

		result, err := adminService.UpdateUser(ctx, req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// w.Write([]byte(response))
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			return
		}
	}
}
