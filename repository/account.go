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
	DeleteAccount(req dto.DeleteAccountReq) (response string, err error)
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
	stmt.Exec(101, req.User_id, req.Branch_id, req.Account_type, req.Balance, time.Now().Unix(), time.Now().Unix())

	return "\n Account Created Successfully !!", nil
}

func (db *AccountStore) DeleteAccount(req dto.DeleteAccountReq) (response string, err error) {

	//For Inserting
	stmt, err := db.DB.Prepare(`DELETE FROM account WHERE acc_no=? AND user_id=?`)
	if err != nil {
		return "", fmt.Errorf("errror While inserting CreateAccount data in db")
	}
	stmt.Exec(req.Account_no, req.User_id)

	return "\n Account Deleted Successfully !!", nil
}
