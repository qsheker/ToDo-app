package repository

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Injector() *gorm.DB {
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}

	db, err := NewDB(Config{
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

	if err := AutoMigrate(db); err != nil {
		log.Fatal("Migration failed:", err)
	}
	return db
}
func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("C:\\Users\\aldik\\Desktop\\GoRest\\configs")
	return viper.ReadInConfig()
}
