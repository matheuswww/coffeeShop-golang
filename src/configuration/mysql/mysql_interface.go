package mysql

import "database/sql"

func NewMysql(host, name, password string, port int) Mysql {
	return &mysql{
		host:     host,
		name:     name,
		password: password,
		port:     port,
	}
}

type Mysql interface {
	NewMysqlConnection() (*sql.DB, error)
}

type mysql struct {
	host     string
	name     string
	password string
	port     int
}
