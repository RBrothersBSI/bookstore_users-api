package mysql_utils

import (
	"github.com/RBrothersBSI/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr{
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.ContainsAny(err.Error(), errorNoRows){
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number{
	// 1062 is the number of the duplicate email sql error
	case 1062:
		return errors.NewBadRequestError("Email already exists")
	}
	return errors.NewInternalServerError("error processing request")
}