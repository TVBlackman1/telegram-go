package repository

import (
	"database/sql"
	"fmt"

	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/repository/utils"
	"github.com/google/uuid"
)

type UserRepository interface {
	Add(query CreateUserDto) (uuid.UUID, error)
	Remove(interface{})
	Edit(interface{})
	GetList(query UserListQuery) UsersDbMetaDto
	GetOne(query UserQuery) (UserDbDto, error)
	SetNewStateUUID(userUUID uuid.UUID, stateUUID uuid.UUID) error
	// TODO SetNewStateOnFly(state interface{}) error
}

type CreateUserDto struct {
	Id      uuid.UUID    `db:"id"`
	Name    string       `db:"name"`
	Login   string       `db:"login"`
	ChatId  types.ChatId `db:"chat_id"`
	StateId uuid.UUID    `db:"state_id"`
}

type UserDbDto struct {
	Id      uuid.UUID
	Name    string
	Login   string
	ChatId  types.ChatId
	StateId uuid.UUID
	State   StateDbDto
}

type UserDb struct {
	Id      uuid.UUID      `db:"id"`
	Name    sql.NullString `db:"name"`
	Login   sql.NullString `db:"login"`
	ChatId  types.ChatId   `db:"chat_id"`
	StateId uuid.UUID      `db:"state_id"`
	State   StateDbDto     `db:"state"`
}

type UsersDbMetaDto struct {
	Data []UserDbDto
	Meta utils.Pagination
}

type UserListQuery struct {
	Pagination utils.QueryPagination
	Name       string
	State      string
}

type UserQuery struct {
	Login  string
	ChatId types.ChatId
}

const (
	USERS_LIMIT     = 30
	USERS_TABLENAME = "users"
)

func UserDbToUserDbDto(userDb UserDb) UserDbDto {
	user := UserDbDto{
		Id:      userDb.Id,
		ChatId:  userDb.ChatId,
		StateId: userDb.StateId,
		State:   userDb.State,
	}
	if userDb.Name.Valid {
		user.Name = userDb.Name.String
	}
	if userDb.Login.Valid {
		user.Login = userDb.Login.String
	}
	fmt.Printf("%+v", user)
	return user
}
