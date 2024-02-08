package account

import "net/http"

func CreateAccount(w http.ResponseWriter, r *http.Request) { //POST

}

func Deposite(w http.ResponseWriter, r *http.Request) { //PUT pathparam

}

func Withdrawal(w http.ResponseWriter, r *http.Request) { //PUT pathparam

}

func DeleteAccount(w http.ResponseWriter, r *http.Request) { //Delete

}

func ViewStatement(w http.ResponseWriter, r *http.Request) { //GET pathparam

}

//Domain => struct and its related vars, const, all data rather than bussiness logic i.e.Service
//rollback of db transaction
//
