package enduser

import "github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/repository"

type service struct {
	UserRepo repository.UserStorer
}

// All User related funcs that processing From DB in Bussiness Logic
type Service interface {
}

func NewService(UserRepo repository.UserStorer) Service {
	return &service{
		UserRepo: UserRepo,
	}
}

//All Bussiness Logic Related funcs of USER with logic here onwards=>
