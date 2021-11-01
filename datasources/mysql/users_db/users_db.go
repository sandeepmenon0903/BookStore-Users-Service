package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
	Sample string
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "password", "127.0.0.1:3306", "users")
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	log.Println(Client)
	log.Println(err)
}
