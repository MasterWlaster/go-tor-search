package service

import "tor_search/src/repository"

type Logger struct {
	repository *repository.Repository
}

func NewLogger(repository *repository.Repository) *Logger {
	return &Logger{repository: repository}
}

func (l *Logger) Log(err error) {
	if err == nil {
		return
	}
	l.repository.Logger.Log(err)
}
