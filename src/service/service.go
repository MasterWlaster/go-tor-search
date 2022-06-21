package service

import (
	"tor_search/src"
	"tor_search/src/repository"
)

type Service struct {
	Search ISearch
	Logger ILogger
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Search: NewSearch(repository),
		Logger: NewLogger(repository)}
}

type ISearch interface {
	Search(word string, limit int) ([]src.Result, error)
}

type ILogger interface {
	Log(err error)
}
