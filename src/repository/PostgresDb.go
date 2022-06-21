package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"tor_search/src"
)

const (
	urls    = `urls`
	words   = `words`
	indexes = `indexes`
)

type PostgresDb struct {
	db *sqlx.DB
}

func NewPostgresDb(db *sqlx.DB) *PostgresDb {
	return &PostgresDb{db: db}
}

func (p *PostgresDb) GetSearchResults(word string, limit, offset int) ([]src.Result, error) {
	var lim interface{}

	if limit < 0 {
		lim = "ALL"
	} else {
		lim = limit
	}

	query := fmt.Sprintf(
		"SELECT t.url, t.word, t.count FROM %s t WHERE t.word LIKE $1 ORDER BY t.count DESC LIMIT %v OFFSET %d",
		indexes, lim, offset)

	rows, err := p.db.Query(query, "%"+word+"%")
	if err != nil {
		return nil, err
	}

	res := make([]src.Result, 0)
	r := src.Result{}

	for rows.Next() {
		err = rows.Scan(&r.Url, &r.Word, &r.Count)
		if err != nil {
			return nil, err
		}

		res = append(res, r)
	}

	return res, nil
}

func ConnectPostgres(c DbConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.Host, c.Port, c.Username, c.Name, c.Password, c.SslMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
