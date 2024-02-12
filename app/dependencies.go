package app

import (
	"database/sql"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/account"
	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/admin"
	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/enduser"
	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/repository"
)

type Dependencies struct {
	UserService    enduser.Service
	AccountService account.Service
	AdminService   admin.Service
}

func NewServices(db *sql.DB) Dependencies {

	//Initialize repo dependencies
	UserRepo := repository.NewUserRepo(db)
	AccountRepo := repository.NewAccountRepo(db)
	AdminRepo := repository.NewAdminRepo(db)

	//Initialize Service Dependencies
	userService := enduser.NewService(UserRepo)
	accountService := account.NewService(AccountRepo)
	adminService := admin.NewService(AdminRepo)

	return Dependencies{
		UserService:    userService,
		AccountService: accountService,
		AdminService:   adminService,
	}
}
