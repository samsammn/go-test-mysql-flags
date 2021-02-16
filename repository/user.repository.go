package repository

import (
	"database/sql"
	"fmt"

	"github.com/samsammn/project7-test/model"
)

type UserRepository interface {
	GetAll() []model.User
	Save(user model.User)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(connDb *sql.DB) UserRepository {
	return &userRepo{db: connDb}
}

func (repo *userRepo) GetAll() []model.User {
	var users []model.User
	result, err := repo.db.Query("Select * from users")

	if err != nil {
		fmt.Println("Error when get all users :", err)
	}

	for result.Next() {
		userRow := model.User{}
		result.Scan(&userRow.Id, &userRow.Name, &userRow.Active)

		users = append(users, userRow)
	}

	return users
}

func (repo *userRepo) Save(user model.User) {
	_, err := repo.db.Exec("Insert into users (name, active) values (?, ?)", user.Name, user.Active)

	if err != nil {
		fmt.Println("Error when saving data user :", err)
	}
}
