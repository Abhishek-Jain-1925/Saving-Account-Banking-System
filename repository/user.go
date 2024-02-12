package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserStore struct {
	BaseRepository
}

// All User related DB activity like sigup,list,signup,update
// have to specify methods in interface then perform operations
type UserStorer interface {
	RepositoryTrasanctions

	GetLoginDetails() (response map[string]string, err error)
	AddUser(req dto.CreateUser) (dto.Response, error)
	UpdateUser(req dto.UpdateUser, user_id int) (dto.UpdateUser, error)

	TokenDetails(email string) (user_id int, role string, err error)
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
	QueryExecuter := db.initiateQueryExecutor(db.DB)
	rows, err := QueryExecuter.Query("SELECT email, password FROM user")
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

func (db *UserStore) AddUser(req dto.CreateUser) (dto.Response, error) {

	//To get Existing value
	var count int64
	QueryExecuter := db.initiateQueryExecutor(db.DB)
	row := QueryExecuter.QueryRow("SELECT MAX(user_id) FROM user")
	err := row.Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			count = 0
		}
		return dto.Response{}, fmt.Errorf("something went wrong")
	}
	//Hashing of Password
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		log.Println("error !! while hashing password, Error : ", err)
		return dto.Response{}, fmt.Errorf("error !! while hashing password, Error : %v", err)
	}

	//For Inserting
	stmt, err := QueryExecuter.Prepare(`INSERT INTO user VALUES(?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		return dto.Response{}, fmt.Errorf("errror While inserting sign-up data in db")
	}
	user_id := count + 1
	stmt.Exec(user_id, req.Name, req.Address, req.Email, string(hashPwd), req.Mobile, req.Role, time.Now().Unix(), time.Now().Unix())

	// resStr := "\n *** Signed Up Successfully ***"
	// resStr += "\n\n Kindly Note following details -"
	// resStr += fmt.Sprintf("\n- User ID : %v", (count + 1))
	// resStr += fmt.Sprintf("\n- Email(username) : %v", req.Email)
	// resStr += fmt.Sprintf("\n- Password : %v", req.Password)
	// resStr += fmt.Sprintf("\n- Role : %v", req.Role)
	// resStr += "\n\n*** You can login now. ***"

	var res dto.Response
	res.User_id = int(user_id)
	res.Name = req.Name
	res.Address = req.Address
	res.Email = req.Email
	res.Password = req.Password
	res.Mobile = req.Mobile
	res.Role = req.Role

	return res, nil
}

func (db *UserStore) UpdateUser(req dto.UpdateUser, user_id int) (dto.UpdateUser, error) {
	//Hashing of password
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return dto.UpdateUser{}, fmt.Errorf(err.Error())
	}

	// For Updating User Info.
	QueryExecuter := db.initiateQueryExecutor(db.DB)
	stmt, err := QueryExecuter.Prepare(`UPDATE user SET name=?, address=?, password=?, mobile=?, updated_at=? WHERE user_id=?`)
	if err != nil {
		return dto.UpdateUser{}, fmt.Errorf("error while updating user data in db: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Name, req.Address, string(hashPwd), req.Mobile, time.Now().Unix(), user_id)
	if err != nil {
		return dto.UpdateUser{}, fmt.Errorf("error executing update statement: %v", err)
	}

	var res dto.UpdateUser
	res.User_id = user_id
	res.Name = req.Name
	res.Address = req.Address
	res.Mobile = req.Mobile
	res.Password = req.Password

	return res, nil
}

func (db *UserStore) TokenDetails(email string) (user_id int, role string, err error) {
	QueryExecuter := db.initiateQueryExecutor(db.DB)
	row := QueryExecuter.QueryRow("SELECT user_id,role FROM user where email=?", email)
	err = row.Scan(&user_id, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, "", fmt.Errorf("record not found")
		}
		return 0, "", fmt.Errorf("something went wrong")
	}

	return user_id, role, nil
}
