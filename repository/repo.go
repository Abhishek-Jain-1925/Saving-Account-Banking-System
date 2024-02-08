package repository

import "database/sql"

//Initialization to baseRepo
//contains commit,rollback,transcations details related operation

type BaseRepository struct {
	DB *sql.DB
}

type RepositoryTrasanctions interface {
}
