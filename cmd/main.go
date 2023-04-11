package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	handler "server-side/internal/handler"
	"server-side/internal/repository"
	"server-side/internal/service"
	"server-side/server"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	servers := service.NewService(repos)
	handlers := handler.NewHandler(servers)

	srv := new(server.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitHandler()); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
