package repository

import (
	"github.com/jmoiron/sqlx"
	"tor_search/src"
)

type Repository struct {
	Memory IMemory
	Logger ILoggerRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Memory: NewPostgresDb(db),
		Logger: NewConsoleLogger()}
}

type IMemory interface {
	GetSearchResults(word string, limit, offset int) ([]src.Result, error)
}

type ILoggerRepository interface {
	Log(err error)
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
	SslMode  string
}
