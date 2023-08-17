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

// github.com/gin-gonic/gin
// github.com/jmoiron/sqlx
// github.com/joho/godotenv
// github.com/spf13/viper
// github.com/asaskevich/govalidator
// github.com/golang-jwt/jwt/v5
// github.com/gin-contrib/cors
// github.com/cloudinary/cloudinary-go/v2

// migrate database
// migrate -path ./migrations -database "postgresql://fazztrack:123456@localhost/coffeeshop_database?port=5432&sslmode=disable&search_path=public" -verbose up
// migrate -path ./migrations -database "postgresql://fazztrack:123456@localhost/coffeeshop_database?port=5432&sslmode=disable&search_path=public" -verbose down
