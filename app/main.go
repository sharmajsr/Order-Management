package main

import (
	"github.com/sharmajsr/oms/api_server"
)

func main() {
	// var configFile string
	// var initDb int

	// flag.StringVar(&configFile, "config", "", "config file")
	// flag.IntVar(&initDb, "initDb", 0, "DB init")
	// flag.Parse()

	// if configFile == "" {

	// 	os.Exit(1)
	// }

	// viper.SetConfigFile(configFile)
	// viper.SetConfigType("yaml")
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	// logger.Fatal("error reading config file", err)
	// }
	// if initDb > 0 {
	// 	db.Init()
	// }

	api_server.RunServer()
}
