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

	CreateAccount(req dto.CreateAccountReq, user_id int) (dto.CreateAccountReq, error)
	DeleteAccount(req dto.DeleteAccountReq, user_id int) (dto.DeleteAccount, error)
	DepositMoney(req dto.Transaction, user_id int) (dto.TransactionResponse, error)
	WithdrawalMoney(req dto.Transaction, user_id int) (dto.TransactionResponse, error)
}

func NewAccountRepo(db *sql.DB) AccountStorer {
	return &AccountStore{
		BaseRepository: BaseRepository{db},
	}
}

func (db *AccountStore) CreateAccount(req dto.CreateAccountReq, user_id int) (dto.CreateAccountReq, error) {

	//To get Existing value
	var count int64
	QueryExecuter := db.initiateQueryExecutor(db.DB)
	row := QueryExecuter.QueryRow("SELECT MAX(acc_no) FROM account")
	err := row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			count = 100
		}
	}
	//For Inserting
	stmt, err := QueryExecuter.Prepare(`INSERT INTO account VALUES(?,?,?,?,?,?,?)`)
	if err != nil {
		return dto.CreateAccountReq{}, fmt.Errorf("errror While inserting CreateAccount data in db")
	}
	acc_no := (count + 1)
	stmt.Exec(acc_no, user_id, req.Branch_id, req.Account_type, req.Balance, time.Now().Unix(), time.Now().Unix())

	res := dto.CreateAccountReq{
		Account_no:   int(acc_no),
		Account_type: req.Account_type,
		Balance:      req.Balance,
		Branch_id:    req.Branch_id,
		User_id:      user_id,
	}
	return res, nil
}

func (db *AccountStore) DeleteAccount(req dto.DeleteAccountReq, user_id int) (dto.DeleteAccount, error) {
	var count int64
	QueryExecuter := db.initiateQueryExecutor(db.DB)
	row := QueryExecuter.QueryRow("SELECT user_id FROM account where acc_no=? AND user_id=?", req.Account_no, user_id)
	err := row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return dto.DeleteAccount{}, fmt.Errorf("record not found, plz provide valid data")
		}
		return dto.DeleteAccount{}, fmt.Errorf("something went wrong")
	}

	//For Inserting
	stmt, err := QueryExecuter.Prepare(`DELETE FROM account WHERE acc_no=? AND user_id=?`)
	if err != nil {
		return dto.DeleteAccount{}, fmt.Errorf("errror While deleting data from db")
	}
	stmt.Exec(req.Account_no, user_id)
	res := dto.DeleteAccount{
		Msg: "account deleted successfully",
	}
	return res, nil
}

func (db *AccountStore) DepositMoney(req dto.Transaction, user_id int) (dto.TransactionResponse, error) {
	//For Deposit Money
	QueryExecuter := db.initiateQueryExecutor(db.DB)
	var balance float64

	row := QueryExecuter.QueryRow("SELECT balance FROM account WHERE acc_no = ? AND user_id=?", req.Account_no, user_id)
	err := row.Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return dto.TransactionResponse{}, fmt.Errorf("no Record found")
		}
		return dto.TransactionResponse{}, fmt.Errorf("something went wrong")
	}
	TotalBal := balance + req.Amount

	stmt, err := QueryExecuter.Prepare(`UPDATE account SET balance=? WHERE acc_no=? AND user_id=?`)
	if err != nil {
		return dto.TransactionResponse{}, fmt.Errorf("errror While inserting CreateAccount data in db")
	}
	stmt.Exec(TotalBal, req.Account_no, user_id)
	res := dto.TransactionResponse{
		Account_no: req.Account_no,
		Balance:    TotalBal,
	}
	return res, nil
}

func (db *AccountStore) WithdrawalMoney(req dto.Transaction, user_id int) (dto.TransactionResponse, error) {

	//For Withdrawal
	QueryExecuter := db.initiateQueryExecutor(db.DB)
	var balance float64
	row := QueryExecuter.QueryRow("SELECT balance FROM account WHERE acc_no = ? AND user_id=?", req.Account_no, user_id)
	err := row.Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return dto.TransactionResponse{}, fmt.Errorf("no record found")
		}
		return dto.TransactionResponse{}, fmt.Errorf("something went wrong")
	}

	if balance < req.Amount {
		return dto.TransactionResponse{}, fmt.Errorf("insufficient balance")
	}
	TotalBal := balance - req.Amount

	stmt, err := QueryExecuter.Prepare(`UPDATE account SET balance=? WHERE acc_no=? AND user_id=?`)
	if err != nil {
		return dto.TransactionResponse{}, fmt.Errorf("errror While inserting CreateAccount data in db")
	}
	stmt.Exec(TotalBal, req.Account_no, user_id)

	res := dto.TransactionResponse{
		Account_no: req.Account_no,
		Balance:    TotalBal,
	}
	return res, nil
}
