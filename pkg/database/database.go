package database

import (
	"GDForum/pkg/config"
	"database/sql"
	"fmt"
	"errors"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

//Connect连接数据库
func Connect(dbConfig gorm.Dialector,_logger gormlogger.Interface){
	var err error
	DB,err = gorm.Open(dbConfig,&gorm.Config{
		Logger:_logger,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}
func CurrentDatabase() (dbname string) {
	dbname = DB.Migrator().CurrentDatabase()
	return
}

func DeleteAllTables() error {

	var err error

	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMysqlDatabase()
	case "sqlite":
		deleteAllSqliteTables()
	default:
		panic(errors.New("database connection not supported"))
	}

	return err
}

func deleteAllSqliteTables() error {
	tables := []string{}
	DB.Select(&tables, "SELECT name FROM sqlite_master WHERE type='table'")
	for _, table := range tables {
		DB.Migrator().DropTable(table)
	}
	return nil
}

func deleteMysqlDatabase() error {
	dbname := CurrentDatabase()
	sql := fmt.Sprintf("DROP DATABASE %s;", dbname)
	if err := DB.Exec(sql).Error; err != nil {
		return err
	}
	sql = fmt.Sprintf("CREATE DATABASE %s;", dbname)
	if err := DB.Exec(sql).Error; err != nil {
		return err
	}
	sql = fmt.Sprintf("USE %s;", dbname)
	if err := DB.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}
