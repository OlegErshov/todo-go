package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo_go "todo-go"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(user todo_go.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO users (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id")
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Auth) GetUser(username, password string) (todo_go.User, error) {
	var user todo_go.User
	query := fmt.Sprintf("SELECT * FROM users WHERE username=$1 AND password_hash=$2")
	err := r.db.Get(&user, query, username, password)
	return user, err
}
