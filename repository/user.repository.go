package repository

import (
	"database/sql"
	"fmt"

	"github.com/samsammn/project7-test/model"
)

type UserRepository interface {
	GetAll() []model.User
	GetById(id string) model.User
	Save(user model.User)
	Update(id string, user model.User)
	Delete(id string)
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

func (repo *userRepo) GetById(id string) model.User {
	user := model.User{}
	result := repo.db.QueryRow("Select * from users where id=?", id)

	result.Scan(&user.Id, &user.Name, &user.Active)

	return user
}

func (repo *userRepo) Save(user model.User) {
	_, err := repo.db.Exec("Insert into users (name, active) values (?, ?)", user.Name, user.Active)

	if err != nil {
		fmt.Println("Error when saving data user :", err)
	}
}

func (repo *userRepo) Update(id string, user model.User) {
	_, err := repo.db.Exec("update users set name=?, active=? where id=?", user.Name, user.Active, id)

	if err != nil {
		fmt.Println("Error when update data user :", err)
	}
}

func (repo *userRepo) Delete(id string) {
	_, err := repo.db.Exec("delete from users where id=?", id)

	if err != nil {
		fmt.Println("Error when delete data user :", err)
	}
}
