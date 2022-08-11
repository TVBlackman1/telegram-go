package repository

import (
	"database/sql"

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
}

type CreateUserDto struct {
	Id     uuid.UUID    `db:"id"`
	Name   string       `db:"name"`
	ChatId types.ChatId `db:"chat_id"`
	// StateId string `db:"state_id"`
}

type UserDbDto struct {
	Id      uuid.UUID
	Name    string
	ChatId  types.ChatId
	StateId string
}

type UserDb struct {
	Id      uuid.UUID      `db:"id"`
	Name    sql.NullString `db:"name"`
	ChatId  types.ChatId   `db:"chat_id"`
	StateId sql.NullString `db:"state_id"`
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
	Name   string
	ChatId types.ChatId
}

const (
	USERS_LIMIT     = 30
	USERS_TABLENAME = "users"
)

func UserDbToUserDbDto(userDb UserDb) UserDbDto {
	user := UserDbDto{
		Id:     userDb.Id,
		ChatId: userDb.ChatId,
	}
	if userDb.Name.Valid {
		user.Name = userDb.Name.String
	}
	if userDb.StateId.Valid {
		user.StateId = userDb.StateId.String
	}
	return user
}
