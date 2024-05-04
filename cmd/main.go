package main

import (
	sever "example/hello"
	"example/hello/pkg/handler"
	"example/hello/pkg/repository"
	"example/hello/pkg/service"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	// repository -> работа с бд
	// service -> сама логика
	// handlers -> обработчики (на данный момент роуты)
	repos := repository.NewRepository()      // в зависимости от env, работаем с нужной бд
	services := service.NewService(repos)    // передаём в сервисы выбранную бд
	handlers := handler.NewHandler(services) // обработчик выбранного сервиса

	srv := new(sever.Server)
	port, _ := strconv.Atoi(viper.GetString("port"))

	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
