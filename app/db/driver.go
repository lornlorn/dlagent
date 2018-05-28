package db

import (
	"app/utils"
	"log"

	"github.com/go-xorm/xorm"
	// _ "github.com/lib/pq" // postgresql driver
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

// Engine xorm
var Engine *xorm.Engine

// InitDB func() error
func InitDB() error {
	var err error
	// dbtype := "postgres"
	// dbstr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", host, port, user, password, dbname)
	// dbstr := "postgres://test:test@localhost:5432/testdb?sslmode=disable"
	// dbstr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)

	dbtype, _ := utils.ReadConf("db", "dbtype")
	dbstr, _ := utils.ReadConf("db", "dbstr")

	log.Printf("%v,%v\n", dbtype, dbstr)

	Engine, err = xorm.NewEngine(dbtype, dbstr)
	if err != nil {
		return err
	}
	Engine.ShowSQL(true)
	err = Engine.Ping()
	if err != nil {
		log.Printf("DB Ping Failed : %v\n", err)
		return err
	}
	return nil
}
