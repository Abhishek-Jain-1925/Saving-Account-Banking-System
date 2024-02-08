package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
)

type UserStore struct {
	BaseRepository
}

// All User related DB activity like sigup,list,signup,update
// have to specify methods in interface then perform operations
type UserStorer interface {
	RepositoryTrasanctions

	GetLoginDetails() (response map[string]string, err error)
	AddUser(req dto.CreateUser) (response string, err error)
}

// All info want to process on DB
type User struct {
	User_id    int
	Name       string
	Address    string
	Email      string
	Password   string
	Mobile     string
	Role       string
	Created_at int
	Updated_at int
}

func NewUserRepo(db *sql.DB) UserStorer {
	return &UserStore{
		BaseRepository: BaseRepository{db},
	}
}

func (db *UserStore) GetLoginDetails() (response map[string]string, err error) {
	rows, err := db.DB.Query("SELECT email, password FROM user")
	if err != nil {
		log.Print("error while fetching login details from database: ", err)
		return nil, fmt.Errorf("error while fetching login details from database")
	}

	LoginMap := make(map[string]string)
	for rows.Next() {
		var email, pwd string
		if err := rows.Scan(&email, &pwd); err != nil {
			log.Print("error while scanning row: ", err)
			continue
		}
		LoginMap[email] = pwd
	}

	return LoginMap, nil

}

func (db *UserStore) AddUser(req dto.CreateUser) (response string, err error) {

	//For Inserting
	stmt, err := db.DB.Prepare(`INSERT INTO user VALUES(?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return "", fmt.Errorf("errror While inserting sign-up data in db")
	}
	stmt.Exec(1, req.Name, req.Address, req.Email, req.Password, req.Mobile, req.Role, time.Now().Unix(), time.Now().Unix())
	return "Signed up Successfully !!", nil
}
