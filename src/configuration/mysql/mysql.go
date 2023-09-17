package mysql

import (
	"database/sql"
	"matheuswww/coffeeShop-golang/src/configuration/logger"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func (m *mysql) NewMysqlConnection() (*sql.DB,error) {
	err := godotenv.Load()
	if err != nil {
		logger.Error("ENV LOADING ERROR!!!",err,zap.String("journey","databaseConnect"))
		return nil,err
	}
	db,err := sql.Open("mysql","root:"+m.password+"@tcp("+m.host+")/"+m.name)
	if err != nil {
		logger.Error("MYSQL DB CONNECT ERROR!!!",err,zap.String("journey","databaseConnect"))
		return nil,err
	}
	if err := db.Ping(); err != nil {
		logger.Error("MYSQL DB CONNECT ERROR!!!",err,zap.String("journey","databaseConnect"))
		return nil, err
	}
	logger.Info("MYSQL DB IS RUNNING!!!")
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Minute * 1)
	return db,nil
}