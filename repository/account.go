package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
)

type AccountStore struct {
	BaseRepository
}

// All Account related DB activity like create account,deposite,withdraw,delete,view statement
// have to specify methods in interface then perform operations
type AccountStorer interface {
	RepositoryTrasanctions

	CreateAccount(req dto.CreateAccountReq) (response string, err error)
}

// To store all account related info in DB
type Account struct {
	Acc_no     int
	User_id    int
	Branch_id  int
	Acc_type   string
	Balance    float64
	Created_at int
	Updated_at int
}

func NewAccountRepo(db *sql.DB) AccountStorer {
	return &AccountStore{
		BaseRepository: BaseRepository{db},
	}
}

func (db *AccountStore) CreateAccount(req dto.CreateAccountReq) (response string, err error) {

	//For Inserting
	stmt, err := db.DB.Prepare(`INSERT INTO account VALUES(?,?,?,?,?,?,?)`)
	if err != nil {
		return "", fmt.Errorf("errror While inserting CreateAccount data in db")
	}
	stmt.Exec(100, req.User_id, req.Branch_id, req.Account_type, req.Balance, time.Now().Unix(), time.Now().Unix())

	return "Account Created Successfully !!", nil
}
