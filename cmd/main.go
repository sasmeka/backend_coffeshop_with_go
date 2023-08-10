package main

import (
	"log"
	"sasmeka/coffeeshop/internal/routers"
	"sasmeka/coffeeshop/pkg"

	"github.com/asaskevich/govalidator"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	database, err := pkg.Postgres_Database()
	if err != nil {
		log.Fatal(err)
	}
	router := routers.Routers(database)
	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/main.go
// go get -u github.com/gin-gonic/gin
// github.com/jmoiron/sqlx v1.3.5 // indirect
// github.com/joho/godotenv

// migrate database

// migrate -path ./migrations -database "postgresql://fazztrack:123456@localhost/coffeeshop_database?sslmode=disable&search_path=public" -verbose up
// migrate -path ./migrations -database "postgresql://fazztrack:123456@localhost/coffeeshop_database?sslmode=disable&search_path=public" -verbose down
