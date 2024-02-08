package app

import (
	"database/sql"
	"fmt"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/enduser"
	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/repository"
)

type Dependencies struct {
	UserService enduser.Service
}

func NewServices(db *sql.DB) Dependencies {

	//Initialize repo dependencies
	UserRepo := repository.NewUserRepo(db)
	AccountRepo := repository.NewAccountRepo(db)
	fmt.Println(UserRepo, AccountRepo)

	//Initialize Service Dependencies
	userService := enduser.NewService(db)

	return Dependencies{
		UserService: userService,
	}
}
