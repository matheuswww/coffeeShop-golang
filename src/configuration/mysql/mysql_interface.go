package mysql

import "database/sql"

func NewMysql() Mysql {
	return &mysql{}
}

type Mysql interface {
	NewMysqlConnection() (*sql.DB, error)
	configConn() error
}

type mysql struct {
	host     string
	name     string
	password string
	port     int
}
