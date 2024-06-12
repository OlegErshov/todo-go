package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	todo_go "todo-go"
	"todo-go/pkg/handler"
	"todo-go/pkg/repository"
	"todo-go/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := repository.NewDb(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("Error connecting to database, %s", err)
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	srv := new(todo_go.Server)
	if err := srv.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Erorr in attempting to start server: %v", err.Error)
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
