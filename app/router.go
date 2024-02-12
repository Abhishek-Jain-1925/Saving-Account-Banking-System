package app

import (
	"net/http"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/account"
	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/admin"
	user "github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/enduser"
	"github.com/gorilla/mux"
)

func NewRouter(deps Dependencies) *mux.Router {

	r := mux.NewRouter()
	//User Related Activity
	r.HandleFunc("/login", user.Login(deps.UserService)).Methods(http.MethodPost)
	r.HandleFunc("/signup", user.Signup(deps.UserService)).Methods(http.MethodPost)
	r.HandleFunc("/update_user", user.Update(deps.UserService)).Methods(http.MethodPut)

	//Account Related Activity
	r.HandleFunc("/account/create", account.Create(deps.AccountService)).Methods(http.MethodPost)
	r.HandleFunc("/account/deposit", account.Deposit(deps.AccountService)).Methods(http.MethodPut)
	r.HandleFunc("/account/withdrawal", account.Withdrawal(deps.AccountService)).Methods(http.MethodPut)
	r.HandleFunc("/account/delete", account.Delete(deps.AccountService)).Methods(http.MethodDelete)
	r.HandleFunc("account/statement", account.ViewStatement)

	//Admin Side Activity
	r.HandleFunc("admin/statement", account.ViewStatement)
	r.HandleFunc("/admin/user_list", admin.ListUsers(deps.AdminService)).Methods(http.MethodGet)
	r.HandleFunc("/admin/update_user", admin.Update(deps.AdminService)).Methods(http.MethodPut)

	return r
}
