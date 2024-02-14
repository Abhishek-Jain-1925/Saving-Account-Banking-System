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
	subrouter := r.PathPrefix("/account").Subrouter()
	subrouter.HandleFunc("/create", account.Create(deps.AccountService)).Methods(http.MethodPost)
	subrouter.HandleFunc("/deposit", account.Deposit(deps.AccountService)).Methods(http.MethodPut)
	subrouter.HandleFunc("/withdrawal", account.Withdrawal(deps.AccountService)).Methods(http.MethodPut)
	subrouter.HandleFunc("/delete", account.Delete(deps.AccountService)).Methods(http.MethodDelete)
	subrouter.HandleFunc("/balance", account.ViewBalance(deps.AccountService)).Methods(http.MethodGet)

	//Admin Side Activity
	subrouter1 := r.PathPrefix("/admin").Subrouter()
	//r.HandleFunc("admin/statement", account.ViewStatement)
	subrouter1.HandleFunc("/user_list", admin.ListUsers(deps.AdminService)).Methods(http.MethodGet)
	subrouter1.HandleFunc("/update_user", admin.Update(deps.AdminService)).Methods(http.MethodPut)

	return r
}
