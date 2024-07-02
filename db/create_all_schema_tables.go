package db

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sharmajsr/oms/common"
)

func createTable(tableSchemaFile string) {
	// logger.Info("Creating table ", "tableSchemaFile", tableSchemaFile)

	var asset_schema_path string
	asset_schema_path = common.CFG_KEY_DB_SCHEMA_PATH

	asset_schema_path = filepath.Join(asset_schema_path, tableSchemaFile)

	blob, err := os.ReadFile(asset_schema_path)
	if err != nil {
		panic(err)
	}
	if _, err = sqldb.Exec(string(blob)); err != nil {
		panic(err)
	} else {
		fmt.Println("Created table ", "tableSchemaFile", tableSchemaFile)
	}
}

func createAllSqlTables() {
	createTable("users.sql")
	createTable("products.sql")
	createTable("orders.sql")
}
