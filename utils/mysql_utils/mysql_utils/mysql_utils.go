package mysql_utils

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/meetmanok/bookstore_users-api/utils/errors"
	"strings"
)

const (
	ErrorNoRows      = "no rows in result set"
)

func ParseError(err error) *errors.RestErr{
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows){
			return errors.NewNotFoundErr(fmt.Sprintf("no record matching given id"))
		}
		return errors.NewInternalServerErr("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestErr(fmt.Sprintf("invalid data"))
	case 1292:
		return errors.NewBadRequestErr(sqlErr.Message)
	}
	return errors.NewInternalServerErr("error processing request")
}