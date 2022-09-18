package postgres

import (
	"fmt"
	"os"

	"github.com/TVBlackman1/telegram-go/pkg/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type StateRepository struct {
	db *sqlx.DB
}

func NewStateRepository(db *sqlx.DB) *StateRepository {
	return &StateRepository{db}
}

func (rep *StateRepository) Add(query repository.CreateStateDto) (uuid.UUID, error) {
	request := fmt.Sprintf("INSERT INTO %s(id, name, context) VALUES ('%s', '%s', '%s') RETURNING id;\n",
		repository.STATES_TABLENAME,
		query.Id,
		query.Name,
		query.Context,
	)
	var uuidStr string
	err := rep.db.Get(&uuidStr, request)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return uuid.UUID{}, err
	}
	return uuid.Parse(uuidStr)
}

func (rep *StateRepository) Remove(interface{}) {

}

func (rep *StateRepository) GetList(interface{}) {
}

func (rep *StateRepository) Edit(interface{}) {

}
