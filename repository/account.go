package repository

import "database/sql"

type AccountStore struct {
	BaseRepository
}

// All Account related DB activity like create account,deposite,withdraw,delete,view statement
// have to specify methods in interface then perform operations
type AccountStorer interface {
	RepositoryTrasanctions
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
