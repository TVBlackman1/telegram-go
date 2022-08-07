package postgres

import (
	"fmt"
	"os"
	"strings"

	"github.com/TVBlackman1/telegram-go/pkg/repository"
	"github.com/TVBlackman1/telegram-go/pkg/repository/utils"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (rep *UserRepository) Add(interface{}) {

}

func (rep *UserRepository) Remove(interface{}) {

}

func (rep *UserRepository) GetList(query repository.UserListQuery) repository.UsersDbMetaDto {
	selectedFields := "id, name, chat_id, state_id"
	var logicBuilder strings.Builder
	fmt.Fprintf(&logicBuilder, "FROM %s", repository.USERS_TABLENAME)
	if query.Name != "" {
		fmt.Fprintf(&logicBuilder, " where name like '%%%s%%'", query.Name)
	}
	req := &utils.RequestWithPagination{
		Db:         rep.db,
		LogicPart:  logicBuilder.String(),
		Selected:   selectedFields,
		Pagination: query.Pagination,
	}
	var users []repository.UserDbDto
	pagination, err := req.SelectIn(&users)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad request: %s", err.Error())
	}

	fmt.Printf("%+v\n", pagination)
	fmt.Println(users)
	return repository.UsersDbMetaDto{
		Data: users,
		Meta: pagination,
	}
}

func (rep *UserRepository) Edit(interface{}) {

}
