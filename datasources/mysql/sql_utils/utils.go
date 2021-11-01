package utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/sandeepmenon0903/BookStore-Users-Service/utils/errors"
)

const (
	emptyResultCommand = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), emptyResultCommand) {
			return errors.NotFound("Record Doesnot Exist")
		}
		return errors.NewInternalServerError("Errors Parsing DataBase Response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.BadRequest("Record Already Exist in the DataBase")
	}
	return errors.NewInternalServerError("Error Processing Request")
}
