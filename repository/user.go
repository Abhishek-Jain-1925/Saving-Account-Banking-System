package repository

import "database/sql"

type UserStore struct {
	BaseRepository
}

// All User related DB activity like sigup,list,signup,update
// have to specify methods in interface then perform operations
type UserStorer interface {
	RepositoryTrasanctions
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
