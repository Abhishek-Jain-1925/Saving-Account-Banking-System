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
	DeleteAccount(req dto.DeleteAccountReq, user_id int) (response string, err error)
	DepositMoney(req dto.Transaction, user_id int) (dto.Transaction, error)
	WithdrawalMoney(req dto.Transaction, user_id int) (dto.Transaction, error)
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
	// resStr := "\n *** Account Created Successfully ***"
	// resStr += "\n\n Kindly Note following details -"
	// resStr += fmt.Sprintf("\n- Account Number : %v", acc_no)
	// resStr += fmt.Sprintf("\n- Account Type : %v", req.Account_type)
	// resStr += fmt.Sprintf("\n- Branch ID : %v", req.Branch_id)
	// resStr += fmt.Sprintf("\n- User ID : %v", user_id)

	var res dto.CreateAccountReq
	res.Account_no = int(acc_no)
	res.Account_type = req.Account_type
	res.Balance = req.Balance
	res.Branch_id = req.Branch_id
	res.User_id = user_id

	return res, nil
}

func (db *AccountStore) DeleteAccount(req dto.DeleteAccountReq, user_id int) (response string, err error) {
	var count int64
	QueryExecuter := db.initiateQueryExecutor(db.DB)
	row := QueryExecuter.QueryRow("SELECT user_id FROM account where acc_no=? AND user_id=?", req.Account_no, user_id)
	err = row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("record not found, plz provide valid data")
		}
		return "", fmt.Errorf("something went wrong")
	}

	//For Inserting
	stmt, err := QueryExecuter.Prepare(`DELETE FROM account WHERE acc_no=? AND user_id=?`)
	if err != nil {
		return "", fmt.Errorf("errror While deleting data from db")
	}
	stmt.Exec(req.Account_no, user_id)
	return "\n Account Deleted Successfully !!", nil
}

func (db *AccountStore) DepositMoney(req dto.Transaction, user_id int) (dto.Transaction, error) {

	//For Deposit Money
	QueryExecuter := db.initiateQueryExecutor(db.DB)
	var balance float64
	var res dto.Transaction
	row := QueryExecuter.QueryRow("SELECT balance FROM account WHERE acc_no = ? AND user_id=?", req.Account_no, user_id)
	err := row.Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return dto.Transaction{}, fmt.Errorf("no Record found")
		}
		return dto.Transaction{}, fmt.Errorf("something went wrong")
	}
	TotalBal := balance + req.Amount

	stmt, err := QueryExecuter.Prepare(`UPDATE account SET balance=? WHERE acc_no=? AND user_id=?`)
	if err != nil {
		return dto.Transaction{}, fmt.Errorf("errror While inserting CreateAccount data in db")
	}
	stmt.Exec(TotalBal, req.Account_no, user_id)

	// resStr += "\n*** Money Deposited Successfully ***"
	// resStr += fmt.Sprintf("\n New Total Balance : %.2f", TotalBal)
	res.Account_no = req.Account_no
	res.Amount = TotalBal
	return res, nil
}

func (db *AccountStore) WithdrawalMoney(req dto.Transaction, user_id int) (dto.Transaction, error) {

	//For Withdrawal
	QueryExecuter := db.initiateQueryExecutor(db.DB)
	var balance float64
	row := QueryExecuter.QueryRow("SELECT balance FROM account WHERE acc_no = ? AND user_id=?", req.Account_no, user_id)
	err := row.Scan(&balance)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No record found !!")
			return dto.Transaction{}, fmt.Errorf("no record found")
		}
		return dto.Transaction{}, fmt.Errorf("something went wrong")
	}
	if balance < req.Amount {
		return dto.Transaction{}, fmt.Errorf("insufficient balance")
	}
	//resStr := fmt.Sprintf("\n Previous Balance : %.2f", balance)
	TotalBal := balance - req.Amount

	stmt, err := QueryExecuter.Prepare(`UPDATE account SET balance=? WHERE acc_no=? AND user_id=?`)
	if err != nil {
		return dto.Transaction{}, fmt.Errorf("errror While inserting CreateAccount data in db")
	}
	stmt.Exec(TotalBal, req.Account_no, user_id)

	// resStr += "\n*** Money Withdrawal Successfully ***"
	// resStr += fmt.Sprintf("\n New Total Balance : %.2f", TotalBal)
	var res dto.Transaction
	res.Account_no = req.Account_no
	res.Amount = TotalBal
	return res, nil
}
