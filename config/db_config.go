package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
}

func InitDatabase() *gorm.DB {
	file, err := os.ReadFile("config/database_config.yml")

	if err != nil {
		log.Println(err)
		panic(err)
	}

	var databaseConfig DatabaseConfig
	err = yaml.Unmarshal(file, &databaseConfig)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		databaseConfig.Username, databaseConfig.Password, databaseConfig.Host, databaseConfig.Port, databaseConfig.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	//not use still connecting
	sqlDB.SetMaxIdleConns(10)
	//amount of request to database
	sqlDB.SetMaxOpenConns(100)
	//idle un connecting after 5 minute
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)
	//time of connecting to db
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
