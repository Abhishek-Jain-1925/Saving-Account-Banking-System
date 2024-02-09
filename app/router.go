package app

import (
	"net/http"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/account"
	user "github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/enduser"
	"github.com/gorilla/mux"
)

func NewRouter(deps Dependencies) *mux.Router {

	r := mux.NewRouter()
	//User Related Activity
	r.HandleFunc("/login", user.Login(deps.UserService)).Methods(http.MethodPost)
	r.HandleFunc("/signup", user.Signup(deps.UserService)).Methods(http.MethodPost)
	r.HandleFunc("/updateEndUser", user.Update(deps.UserService)).Methods(http.MethodPut)

	//Account Related Activity
	r.HandleFunc("/account/create", account.Create(deps.AccountService)).Methods(http.MethodPost)
	r.HandleFunc("/account/deposite", account.Deposit)
	r.HandleFunc("/account/withdraw", account.Withdrawal)
	r.HandleFunc("/account/delete", account.Delete(deps.AccountService)).Methods(http.MethodDelete)
	r.HandleFunc("account/statement", account.ViewStatement)

	//Admin Side Activity
	r.HandleFunc("admin/statement", account.ViewStatement)
	r.HandleFunc("/getUsersList", user.List)

	return r
}
