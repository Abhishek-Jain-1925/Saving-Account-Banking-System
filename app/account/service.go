package account

import (
	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/repository"
)

type service struct {
	AccountRepo repository.AccountStorer
}

type Service interface {
}

func NewService(AccountRepo repository.AccountStorer) Service {
	return &service{
		AccountRepo: AccountRepo,
	}
}

//All Account related bussiness logic here onwards=>
