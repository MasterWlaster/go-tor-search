package service

import (
	"tor_search/src"
	"tor_search/src/repository"
)

type Search struct {
	repository *repository.Repository
}

func NewSearch(repository *repository.Repository) *Search {
	return &Search{repository: repository}
}

func (s *Search) Search(word string, limit int) ([]src.Result, error) {
	if limit == 0 {
		return make([]src.Result, 0), nil
	}

	return s.repository.Memory.GetSearchResults(word, limit, 0)
}
