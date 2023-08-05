package main

import (
	"log"
	"sasmeka/coffeeshop/internal/routers"
	"sasmeka/coffeeshop/pkg"

	_ "github.com/joho/godotenv/autoload"
)

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
