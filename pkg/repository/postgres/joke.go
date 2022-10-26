package postgres

import (
	"fmt"
	"os"

	"github.com/TVBlackman1/telegram-go/pkg/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type JokeRepository struct {
	db *sqlx.DB
}

func NewJokeRepository(db *sqlx.DB) *JokeRepository {
	return &JokeRepository{db}
}

func (rep *JokeRepository) Add(query interface{}) (uuid.UUID, error) {
	return uuid.NewUUID()
}

func (rep *JokeRepository) Remove(interface{}) {
}

func (rep *JokeRepository) GetList(interface{}) {
}

func (rep *JokeRepository) Edit(interface{}) {

}

func (rep *JokeRepository) Count(interface{}) (int, error) {
	request := fmt.Sprintf("select count(*) from %s;\n",
		repository.JOKES_TABLENAME,
	)
	var getCount GetCount
	err := rep.db.Get(&getCount, request)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 0, err
	}
	return getCount.Count, nil
}

func (rep *JokeRepository) GetOne(query repository.JokeQuery) (repository.JokeDbDto, error) {
	// tODO to common db invoke
	request := fmt.Sprintf("select * from %s offset %d limit 1;\n",
		repository.JOKES_TABLENAME,
		query.Offset,
	)
	var jokeDb repository.JokeDb
	err := rep.db.Get(&jokeDb, request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad request: %s\n", err.Error())
		return repository.JokeDbDto{}, err
	}
	jokes := repository.JokeDbToJokeDbDto(jokeDb)
	return jokes, nil
}
