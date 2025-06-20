package handlers

import (
	database "github.com/williamligtenberg/workout-tracker/database"

	"github.com/go-sql-driver/mysql"
)

type User struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(u *User) error {
	stmt, err := database.DB.Prepare("INSERT INTO users (uuid, username, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.UUID, u.Username, u.Email, u.Password)
	return err
}

func IsDuplicateUserError(err error) bool {
	mysqlErr, ok := err.(*mysql.MySQLError)
	return ok && mysqlErr.Number == 1062
}
