package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/sharmajsr/oms/common"
)

var onceInit sync.Once
var sqldb *sql.DB
var queries *Queries

// func connectDb() {

// }

func genUrl(dbName string, dbPass string, dbUser string, dbHost string, dbPort string) string {

	var url string
	if dbPass == "" {
		url = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=disable", dbUser, dbHost, dbPort, dbName)
	} else {
		url = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	}

	return url

}

func Init() {
	var err error
	onceInit.Do(func() {
		dbName := common.CFG_KEY_PG_NAME
		dbPass := common.CFG_KEY_PG_PWD
		dbUser := common.CFG_KEY_PG_USER
		dbHost := common.CFG_KEY_PG_HOST
		dbPort := common.CFG_KEY_PG_PORT
		db_url := genUrl(dbName, dbPass, dbUser, dbHost, dbPort)
		fmt.Println(db_url)
		sqldb, err = sql.Open("postgres", db_url)
		if err != nil {
			fmt.Printf("unable to create sql connection %s", err)
		}

		queries = New(sqldb)
		fmt.Println("after opening connection")
		createAllSqlTables()
	})
}

func GetQueries() *Queries {
	return queries
}

func GetSqlDb() *sql.DB {
	return sqldb
}
