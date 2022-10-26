package repository

import (
	"database/sql"

	"github.com/google/uuid"
)

type JokeRepository interface {
	Add(interface{}) (uuid.UUID, error)
	Remove(interface{})
	Edit(interface{})
	GetList(interface{})
	Count(interface{}) (int, error)
	GetOne(JokeQuery) (JokeDbDto, error)
}

type JokeDbDto struct {
	Id          uuid.UUID
	Name        string
	Text        string
	Author      string
	Explanation string
}

type JokeDb struct {
	Id          uuid.UUID      `db:"id"`
	Name        sql.NullString `db:"name"`
	Text        sql.NullString `db:"text"`
	Author      sql.NullString `db:"author"`
	Explanation sql.NullString `db:"explanation"`
}

type JokeQuery struct {
	Offset int
}

const (
	JOKES_TABLENAME = "jokes"
)

func JokeDbToJokeDbDto(jokeDb JokeDb) JokeDbDto {
	ret := JokeDbDto{
		Id: jokeDb.Id,
	}
	if jokeDb.Name.Valid {
		ret.Name = jokeDb.Name.String
	}
	if jokeDb.Text.Valid {
		ret.Text = jokeDb.Text.String
	}
	if jokeDb.Author.Valid {
		ret.Author = jokeDb.Author.String
	}
	if jokeDb.Explanation.Valid {
		ret.Explanation = jokeDb.Explanation.String
	}
	return ret
}
