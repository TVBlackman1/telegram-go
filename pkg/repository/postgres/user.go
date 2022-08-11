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
	fmt.Println(uuidStr)
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
	var usersDb []repository.UserDb
	pagination, err := req.SelectIn(&usersDb)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad request: %s", err.Error())
	}

	users := make([]repository.UserDbDto, len(usersDb))
	for i, userDb := range usersDb {
		users[i] = repository.UserDbDto{
			Id:     userDb.Id,
			ChatId: userDb.ChatId,
		}
		if userDb.Name.Valid {
			users[i].Name = userDb.Name.String
		}
		if userDb.StateId.Valid {
			users[i].StateId = userDb.StateId.String
		}
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

func (rep *UserRepository) GetOne(query repository.UserQuery) (repository.UserDbDto, error) {
	selectedFields := "id, name, chat_id, state_id"
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
	fmt.Println(request.String())
	err := rep.db.Get(&userDb, request.String())
	// TODO nullstrings to other place
	if err != nil {
		fmt.Fprintf(os.Stderr, "Bad request: %s\n", err.Error())
		return repository.UserDbDto{}, err
	}
	user := repository.UserDbDto{
		Id:     userDb.Id,
		ChatId: userDb.ChatId,
	}
	if userDb.Name.Valid {
		user.Name = userDb.Name.String
	}
	if userDb.StateId.Valid {
		user.StateId = userDb.StateId.String
	}
	fmt.Println(user)
	return user, nil
}

func addListQueryConditions(logicBuilder *strings.Builder, query repository.UserListQuery) {
	if query.Name != "" {
		fmt.Fprintf(logicBuilder, " where name like '%%%s%%'", query.Name)
	}
}

func addQueryConditions(logicBuilder *strings.Builder, query repository.UserQuery) {
	alreadyWithCondition := false
	if query.Name != "" {
		addSQLKeywords(logicBuilder, alreadyWithCondition)
		fmt.Fprintf(logicBuilder, "name='%s'", query.Name)
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
