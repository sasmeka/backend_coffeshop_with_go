package pkg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Postgres_Database() (*sqlx.DB, error) {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s", host, user, password, dbName, port)

	return sqlx.Connect("postgres", config)

}
