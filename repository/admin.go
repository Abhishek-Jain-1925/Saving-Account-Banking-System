package repository

type AdminStore struct {
	BaseRepository
}

type AdminStorer interface {
	RepositoryTrasanctions
}
