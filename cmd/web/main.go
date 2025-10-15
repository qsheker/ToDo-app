package main

import (
	"log"

	"github.com/qsheker/ToDo-app/internal/repository"
	"github.com/spf13/viper"
)

// @title ToDo App API
// @version 1.0
// @description REST API for working with todo's

// @host localhost:8081
// @BasePath /
func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}
	//port := viper.GetString("app.port")

	// GORM connection
	db, err := repository.NewDB(repository.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.dbname"),
		SSLMode:  viper.GetString("database.sslmode"),
	})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	if err := repository.AutoMigrate(db); err != nil {
		log.Fatal("Migration failed:", err)
	}

}
func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	return viper.ReadInConfig()
}
