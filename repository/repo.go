package repository

import (
	"context"
	"database/sql"
	"log"
)

//Initialization to baseRepo
//contains commit,rollback,transcations details related operation

type BaseRepository struct {
	DB *sql.DB
}

type RepositoryTrasanctions interface {
	initiateQueryExecutor(tx *sql.DB) *sql.DB
	BeginTx(ctx context.Context) (tx *sql.Tx, err error)
	CommitTx(tx *sql.Tx) (err error)
	RollbackTx(tx *sql.Tx) (err error)
	HandleTransaction(ctx context.Context, tx *sql.Tx, incomingErr error) (err error)
	GetConn() (conn *sql.DB)
}

func (repo *BaseRepository) GetConn() (conn *sql.DB) {
	return repo.DB
}

func (repo *BaseRepository) BeginTx(ctx context.Context) (tx *sql.Tx, err error) {

	sqlDB, err := repo.DB.Begin()
	if err != nil {
		log.Printf("error occured while initiating database transaction: %v", err.Error())
		return nil, err
	}

	return sqlDB, nil
}

func (repo *BaseRepository) RollbackTx(tx *sql.Tx) (err error) {
	err = tx.Rollback()
	return
}

func (repo *BaseRepository) CommitTx(tx *sql.Tx) (err error) {
	err = tx.Commit()
	return
}

func (repo *BaseRepository) HandleTransaction(ctx context.Context, tx *sql.Tx, incomingErr error) (err error) {
	if incomingErr != nil {
		err = tx.Rollback()
		if err != nil {
			return
		}
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}
	return
}

func (repo *BaseRepository) initiateQueryExecutor(tx *sql.DB) *sql.DB {
	//Populate the query executor so we can join/use a transaction if one is present.
	//If we are not running inside a transaction then the plain *sql.DB object is used.
	executor := repo.DB
	if tx != nil {
		executor = tx
	}
	return executor
}
