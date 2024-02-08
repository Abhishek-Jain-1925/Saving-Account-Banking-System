package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitializeDB() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "../repository/bank.db")
	if err != nil {
		log.Fatal("error !! while creating with database !!")
		return nil, err
	}
	db = database
	defer db.Close()

	statement, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS user(
		user_id INTEGER PRIMARY KEY,
		name VARCHAR(30) NOT NULL,
		address VARCHAR(50) NOT NULL,
		email VARCHAR(30) NOT NULL,
		password VARCHAR(100) NOT NULL,
		mobile CHAR(10) NOT NULL,
		role VARCHAR(10) NOT NULL,
		created_at INTEGER NOT NULL,
		updated_at INTEGER NOT NULL
	)	
	`)
	if err != nil {
		log.Fatalln("error !! while creating user table !! Due to : ", err)
		return nil, err
	}
	statement.Exec()

	statement, err = db.Prepare(`
	CREATE TABLE IF NOT EXISTS branch(
		id INTEGER PRIMARY KEY,
		name VARCHAR(20) NOT NULL,
		location VARCHAR(15) NOT NULL,
		created_at INTEGER NOT NULL,
		updated_at INTEGER NOT NULL
	)
	`)
	if err != nil {
		log.Fatalln("error !! while creating branch table !! Due to : ", err)
		return nil, err
	}
	statement.Exec()

	statement, err = db.Prepare(`
	CREATE TABLE IF NOT EXISTS account(
		acc_no INTEGER PRIMARY KEY,
		user_id INTEGER NOT NULL,
		branch_id INETGER NOT NULL,
		acc_type VARCHAR(10) NOT NULL,
		balance FLOAT NOT NULL,
		created_at INTEGER NOT NULL,
		updated_at INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES user(user_id),
		FOREIGN KEY(branch_id) REFERENCES branch(id)
	)
	`)
	if err != nil {
		log.Fatalln("error !! while creating account table !! Due to : ", err)
		return nil, err
	}
	statement.Exec()

	statement, err = db.Prepare(`
	CREATE TABLE IF NOT EXISTS tbltransaction(
		transaction_id INTEGER PRIMARY KEY,
		acc_no INTEGER NOT NULL,
		amount FLOAT NOT NULL,
		type VARCHAR(10) NOT NULL,
		updated_at INTEGER NOT NULL,
		FOREIGN KEY(acc_no) REFERENCES account(acc_no)
	)
	`)
	if err != nil {
		log.Fatalln("error !! while creating tbltransaction table !! Due to : ", err)
		return nil, err
	}
	statement.Exec()

	log.Fatalln("Successfully Initialized Database !!")

	statement.Close()

	return db, nil
}
