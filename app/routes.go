package app

import (
	"database/sql"
	"net/http"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/common"
	user "github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/customer"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router, db *sql.DB) {

	//User Related Activity
	r.HandleFunc("/login", user.Login).Methods(http.MethodPost)
	r.HandleFunc("/signup", user.Signup)
	r.HandleFunc("/account/create", user.CreateAccount)
	r.HandleFunc("/account/deposite", user.Deposite)
	r.HandleFunc("/account/withdraw", user.Withdrawal)
	r.HandleFunc("/account/delete", user.DeleteAccount)
	r.HandleFunc("account/statement", common.ViewStatement)

	//Admin Side Activity
	r.HandleFunc("admin/statement", common.ViewStatement)
}
