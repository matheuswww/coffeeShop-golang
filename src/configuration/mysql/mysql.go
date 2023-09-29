package mysql

import (
	"database/sql"
	"errors"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func (m *mysql) NewMysqlConnection() (*sql.DB, error) {
	err := m.configConn()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("mysql", "root:"+m.password+"@tcp("+m.host+")/"+m.name)
	if err != nil {
		logger.Error("MYSQL DB CONNECT ERROR!!!", err, zap.String("journey", "databaseConnect"))
		return nil, err
	}
	if err := db.Ping(); err != nil {
		logger.Error("MYSQL DB CONNECT ERROR!!!", err, zap.String("journey", "databaseConnect"))
		return nil, err
	}
	logger.Info("MYSQL DB IS RUNNING!!!")
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Minute * 1)
	return db, nil
}

func (m *mysql) configConn() error {
	mode := os.Getenv("MODE")
	if mode == "PROD" {
		return nil
	} else if mode == "DEV" {
		m.host = "172.17.0.3"
		m.name = "coffeeShop"
		m.password = "senha"
		m.port = 8080
		return nil
	}
	return errors.New("invalid MODE")
}
