package repository

import (
	"github.com/TVBlackman1/telegram-go/pkg/lib/presenter/types"
	"github.com/TVBlackman1/telegram-go/pkg/repository/utils"
	"github.com/google/uuid"
)

type UserRepository interface {
	Add(query CreateUserDto) (uuid.UUID, error)
	Remove(interface{})
	Edit(interface{})
	GetList(query UserListQuery) UsersDbMetaDto
	GetOne(query UserQuery) UserDbDto
}

type CreateUserDto struct {
	Id     uuid.UUID    `db:"id"`
	Name   string       `db:"name"`
	ChatId types.ChatId `db:"chat_id"`
	// StateId string `db:"state_id"`
}

type UserDbDto struct {
	Id      uuid.UUID    `db:"id"`
	Name    string       `db:"name"`
	ChatId  types.ChatId `db:"chat_id"`
	StateId string       `db:"state_id"`
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
