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
	existingUser, err := rep.GetOne(repository.UserQuery{
		Login: query.Login,
	})
	alreadyExist := err == nil
	if alreadyExist {
		return existingUser.Id, utils.ErrAlreadyExists
	}

	request := fmt.Sprintf("INSERT INTO %s(id, name, login, chat_id) VALUES ('%s', '%s', '%s', '%d') RETURNING id;\n",
		repository.USERS_TABLENAME,
		query.Id,
		query.Name,
		query.Login,
		query.ChatId,
	)
	var uuidStr string
	err = rep.db.Get(&uuidStr, request)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return uuid.UUID{}, err
	}
	return uuid.Parse(uuidStr)
}

func (rep *UserRepository) Remove(interface{}) {

}

func (rep *UserRepository) GetList(query repository.UserListQuery) repository.UsersDbMetaDto {
	selectedFields := "id, name, login, chat_id, state_id"
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
	var usersDb []repository.UserDb
	pagination, err := req.SelectIn(&usersDb)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad request: %s", err.Error())
	}

	users := make([]repository.UserDbDto, len(usersDb))
	for i, userDb := range usersDb {
		users[i] = repository.UserDbToUserDbDto(userDb)
	}

	return repository.UsersDbMetaDto{
		Data: users,
		Meta: pagination,
	}
}

func (rep *UserRepository) Edit(interface{}) {

}

func (rep *UserRepository) GetOne(query repository.UserQuery) (repository.UserDbDto, error) {
	selectedFields := "id, name, login, chat_id, state_id"
	var logicBuilder strings.Builder
	utils.AddPrimaryTableToBuilder(&logicBuilder, repository.USERS_TABLENAME)
	addQueryConditions(&logicBuilder, query)

	var request strings.Builder
	request.WriteString("SELECT ")
	request.WriteString(selectedFields)
	request.WriteRune(' ')
	request.WriteString(logicBuilder.String())
	utils.AddLimit1(&request)

	var userDb repository.UserDb
	err := rep.db.Get(&userDb, request.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad request: %s\n", err.Error())
		return repository.UserDbDto{}, err
	}
	user := repository.UserDbToUserDbDto(userDb)
	return user, nil
}

func (rep *UserRepository) SetNewStateUUID(userUUID uuid.UUID, stateUUID uuid.UUID) error {
	request := fmt.Sprintf(`
	with ctx as (
		select id
		from %s
		where id='%s'
		limit 1
	)
	update %s as u
	set state_id='%s'
	from ctx
	where ctx.id = u.id
	returning u.id;
	`,
		repository.USERS_TABLENAME,
		userUUID,
		repository.USERS_TABLENAME,
		stateUUID,
	)
	var uuidStr string
	err := rep.db.Get(&uuidStr, request)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return err
	}
	return nil
}

func addListQueryConditions(logicBuilder *strings.Builder, query repository.UserListQuery) {
	if query.Name != "" {
		fmt.Fprintf(logicBuilder, " where name like '%%%s%%'", query.Name)
	}
}

func addQueryConditions(logicBuilder *strings.Builder, query repository.UserQuery) {
	alreadyWithCondition := false
	if query.Login != "" {
		addSQLKeywords(logicBuilder, alreadyWithCondition)
		fmt.Fprintf(logicBuilder, "login='%s'", query.Login)
		alreadyWithCondition = true
	}
	if query.ChatId != 0 {
		addSQLKeywords(logicBuilder, alreadyWithCondition)
		fmt.Fprintf(logicBuilder, "chat_id='%d'", query.ChatId)
		alreadyWithCondition = true
	}
}

func addSQLKeywords(logicBuilder *strings.Builder, alreadyWithCondition bool) {
	if alreadyWithCondition {
		fmt.Fprint(logicBuilder, " AND ")
	} else {
		fmt.Fprint(logicBuilder, " WHERE ")
	}
}
