package db

import (
	"database/sql"
	"fmt"

	"github.com/megamsquare/go_setup/config"
)

type Config struct {
	Username string `config:"DB_USERNAME" default:"postgres"`
	Password string `config:"DB_PASSWORD" default:"postgres"`
	Database string `config:"DB_NAME" default:"postgres"`
	Address  string `config:"DB_ADDRESS" default:"localhost"`
	Port     int    `config:"DB_PORT" default:"5432"`
}

func Load_config() *Config {
	var conf Config
	config.Load(&conf)
	return &conf
}

func Connect_db(conf *Config) (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Address, conf.Port, conf.Username, conf.Password, conf.Database)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
