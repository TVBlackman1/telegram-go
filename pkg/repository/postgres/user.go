package postgres

import (
	"fmt"
	"os"
	"strings"

	"github.com/TVBlackman1/telegram-go/pkg/repository"
	"github.com/TVBlackman1/telegram-go/pkg/repository/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (rep *UserRepository) Add(query repository.CreateUserDto) (uuid.UUID, error) {
	request := fmt.Sprintf("INSERT INTO %s(id, name, chat_id) VALUES ('%s', '%s', '%d') RETURNING id;\n",
		repository.USERS_TABLENAME,
		query.Id,
		query.Name,
		query.ChatId,
	)
	var uuidStr string
	err := rep.db.Get(&uuidStr, request)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return uuid.UUID{}, err
	}
	return uuid.Parse(uuidStr)

}

func (rep *UserRepository) Remove(interface{}) {

}

func (rep *UserRepository) GetList(query repository.UserListQuery) repository.UsersDbMetaDto {
	selectedFields := "id, name, chat_id, state_id"
	var logicBuilder strings.Builder
	utils.AddPrimaryTableToBuilder(&logicBuilder, repository.USERS_TABLENAME)
	addListQueryConditions(&logicBuilder, query)

	query.Pagination.CheckValues(repository.USERS_LIMIT)
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

func addListQueryConditions(logicBuilder *strings.Builder, query repository.UserListQuery) {
	if query.Name != "" {
		fmt.Fprintf(logicBuilder, " where name like '%%%s%%'", query.Name)
	}
}
