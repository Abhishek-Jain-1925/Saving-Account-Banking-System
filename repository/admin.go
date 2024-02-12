package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Abhishek-Jain-1925/Saving-Account-Banking-System/app/dto"
	"golang.org/x/crypto/bcrypt"
)

type AdminStore struct {
	BaseRepository
}

type AdminStorer interface {
	RepositoryTrasanctions

	ListUsers(ctx context.Context) ([]dto.Response, error)
	UpdateUserInfo(req dto.UpdateUserInfo) (dto.UpdateUserInfo, error)
}

func NewAdminRepo(db *sql.DB) AdminStorer {
	return &AdminStore{
		BaseRepository: BaseRepository{db},
	}
}

func (db *AdminStore) ListUsers(ctx context.Context) ([]dto.Response, error) {

	//To get All user values
	QueryExecuter := db.initiateQueryExecutor(db.DB)
	rows, err := QueryExecuter.Query("SELECT user_id, name, address, email, password, mobile, role FROM user")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var result []dto.Response
	for rows.Next() {
		var res dto.Response
		if err := rows.Scan(&res.User_id, &res.Name, &res.Address, &res.Email, &res.Password, &res.Mobile, &res.Role); err != nil {
			log.Print("error while scanning row: ", err)
			continue
		}
		result = append(result, res)
	}
	return result, nil
}

func (db *AdminStore) UpdateUserInfo(req dto.UpdateUserInfo) (dto.UpdateUserInfo, error) {
	//Hashing of password
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return dto.UpdateUserInfo{}, fmt.Errorf(err.Error())
	}

	// For Updating User Info.
	QueryExecuter := db.initiateQueryExecutor(db.DB)
	stmt, err := QueryExecuter.Prepare(`UPDATE user SET name=?, address=?,email=?, password=?, mobile=?,role=?, updated_at=? WHERE user_id=?`)
	if err != nil {
		return dto.UpdateUserInfo{}, fmt.Errorf("error while updating user data in db: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Name, req.Address, req.Email, string(hashPwd), req.Mobile, req.Role, time.Now().Unix(), req.User_id)
	if err != nil {
		return dto.UpdateUserInfo{}, fmt.Errorf("error executing updateUserInfo at Admin side: %v", err)
	}

	var res dto.UpdateUserInfo
	res.User_id = req.User_id
	res.Name = req.Name
	res.Address = req.Address
	res.Email = req.Email
	res.Password = req.Password
	res.Role = req.Role

	return res, nil
}
