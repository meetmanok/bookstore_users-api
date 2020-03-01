package users

import (
	"fmt"
	"github.com/meetmanok/bookstore_users-api/datasources/mysql/users_db"
	"github.com/meetmanok/bookstore_users-api/logger"
	"github.com/meetmanok/bookstore_users-api/utils/errors"
	"github.com/meetmanok/bookstore_users-api/utils/mysql_utils/mysql_utils"
)

var (
	usersDB = make(map[int64]*User)
)

const (
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?,?,?,?,?,?);"
	queryUpdateUser  = "UPDATE users SET first_name=?, last_name=?, email=?, date_created=? WHERE id=?;"
	searchUserById   = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	deleteUserById   = "DELETE FROM users WHERE id=?;"
	queryUserByStatus   = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(searchUserById)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerErr("database error")
	}

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to get user by id", getErr)
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, saveErr := insertResult.LastInsertId()
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	_, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Id)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(deleteUserById)
	if err != nil {
		return errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	_, deleteErr := stmt.Exec(user.Id)
	if deleteErr != nil {
		return mysql_utils.ParseError(deleteErr)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr){
	stmt, err := users_db.Client.Prepare(queryUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerErr(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	defer rows.Close()
	if err != nil {
		return nil, errors.NewInternalServerErr(err.Error())
	}

	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		result = append(result, user)
	}
	if len(result) == 0{
		return nil, errors.NewNotFoundErr(fmt.Sprintf("no users matching status %s", status))
	}
	return result, nil
}
