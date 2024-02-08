package app

import (
	"database/sql"
	"net/http"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/account"
	user "github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/enduser"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router, db *sql.DB) {

	//User Related Activity
	r.HandleFunc("/login", user.Login).Methods(http.MethodPost)
	r.HandleFunc("/signup", user.Signup)
	r.HandleFunc("/updateEndUser", user.UpdateEndUSer)
	r.HandleFunc("/getUsersList", user.ListEndUser)

	//Account Related Activity
	r.HandleFunc("/account/create", account.CreateAccount)
	r.HandleFunc("/account/deposite", account.Deposite)
	r.HandleFunc("/account/withdraw", account.Withdrawal)
	r.HandleFunc("/account/delete", account.DeleteAccount)
	r.HandleFunc("account/statement", account.ViewStatement)

	//Admin Side Activity
	r.HandleFunc("admin/statement", account.ViewStatement)
}
