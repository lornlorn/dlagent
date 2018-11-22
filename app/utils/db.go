package utils

import (
	"github.com/go-xorm/xorm"       // _ "github.com/lib/pq" // postgresql driver
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"log"
)

// Engine xorm
var Engine *xorm.Engine

// InitDB func() error
func InitDB(dbtype string, dbstr string) error {
	var err error
	// dbtype := "postgres"
	// dbstr := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", host, port, user, password, dbname)
	// dbstr := "postgres://test:test@localhost:5432/testdb?sslmode=disable"
	// dbstr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)

	log.Printf("%v,%v\n", dbtype, dbstr)

	Engine, err = xorm.NewEngine(dbtype, dbstr)
	if err != nil {
		return err
	}
	Engine.ShowSQL(true)
	err = Engine.Ping()
	if err != nil {
		log.Printf("db.driver.InitDB -> Engine.Ping Error : %v\n", err)
		return err
	}
	return nil
}
