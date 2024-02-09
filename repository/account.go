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
	DepositMoney(req dto.Transaction) (response string, err error)
	WithdrawalMoney(req dto.Transaction) (response string, err error)
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

	//To get Existing value
	var count int64
	row := db.DB.QueryRow("SELECT MAX(user_id) FROM account")
	err = row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			count = 0
		}
		return "", fmt.Errorf("something went wrong")
	}
	//For Inserting
	stmt, err := db.DB.Prepare(`INSERT INTO account VALUES(?,?,?,?,?,?,?)`)
	if err != nil {
		return "", fmt.Errorf("errror While inserting CreateAccount data in db")
	}
	acc_no := (count + 1)
	stmt.Exec(acc_no, req.User_id, req.Branch_id, req.Account_type, req.Balance, time.Now().Unix(), time.Now().Unix())
	resStr := "\n *** Account Created Successfully ***"
	resStr += "\n\n Kindly Note following details -"
	resStr += fmt.Sprintf("\n- Account Number : %v", acc_no)
	resStr += fmt.Sprintf("\n- Account Type : %v", req.Account_type)
	resStr += fmt.Sprintf("\n- Branch ID : %v", req.Branch_id)
	resStr += fmt.Sprintf("\n- User ID : %v", req.User_id)

	return resStr, nil
}

func (db *AccountStore) DeleteAccount(req dto.DeleteAccountReq) (response string, err error) {

	var count int64
	row := db.DB.QueryRow("SELECT user_id FROM account where acc_no=?", req.Account_no)
	err = row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("record not found, plz provide valid data")
		}
		return "", fmt.Errorf("something went wrong")
	}

	//For Inserting
	stmt, err := db.DB.Prepare(`DELETE FROM account WHERE acc_no=? AND user_id=?`)
	if err != nil {
		return "", fmt.Errorf("errror While deleting data from db")
	}
	stmt.Exec(req.Account_no, req.User_id)

	return "\n Account Deleted Successfully !!", nil
}

func (db *AccountStore) DepositMoney(req dto.Transaction) (response string, err error) {

	//For Deposit Money
	var balance float64
	row := db.DB.QueryRow("SELECT balance FROM account WHERE acc_no = ?", req.Account_no)
	err = row.Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no Record found")
		}
		return "", fmt.Errorf("something went wrong")
	}
	resStr := fmt.Sprintf("\n Current Balance : %.2f", balance)

	TotalBal := balance + req.Amount

	stmt, err := db.DB.Prepare(`UPDATE account SET balance=? WHERE acc_no=?`)
	if err != nil {
		return "", fmt.Errorf("errror While inserting CreateAccount data in db")
	}
	stmt.Exec(TotalBal, req.Account_no)
	resStr += "\n*** Money Deposited Successfully ***"
	resStr += fmt.Sprintf("\n New Total Balance : %.2f", TotalBal)
	return resStr, nil
}

func (db *AccountStore) WithdrawalMoney(req dto.Transaction) (response string, err error) {

	//For Withdrawal
	var balance float64
	row := db.DB.QueryRow("SELECT balance FROM account WHERE acc_no = ?", req.Account_no)
	err = row.Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No record found !!")
			return
		}
		return
	}
	if balance < req.Amount {
		return "", fmt.Errorf("insufficient balance")
	}
	resStr := fmt.Sprintf("\n Current Balance : %.2f", balance)

	TotalBal := balance - req.Amount

	stmt, err := db.DB.Prepare(`UPDATE account SET balance=? WHERE acc_no=?`)
	if err != nil {
		return "", fmt.Errorf("errror While inserting CreateAccount data in db")
	}
	stmt.Exec(TotalBal, req.Account_no)

	resStr += "\n*** Money Withdrawal Successfully ***"
	resStr += fmt.Sprintf("\n New Total Balance : %.2f", TotalBal)
	return resStr, nil
}
