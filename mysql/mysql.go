package mysql

import (
    "database/sql"
    "github.com/go-sql-driver/mysql"
)

var dbMap map[string]*sql.DB = make(map[string]*sql.DB)

//
func InitDB(dbName string, cfg *mysql.Config) error {
    db, err := sql.Open("mysql", cfg.FormatDSN())
    if nil != err {
        return err
    }
    db.SetMaxIdleConns(20)
    db.SetMaxOpenConns(50)
    dbMap[dbName] = db
    return nil
}

func GetDB(dbName string) *sql.DB {
    db, ok := dbMap[dbName]
    if ok {
        return db
    }else {
        return nil
    }
}
