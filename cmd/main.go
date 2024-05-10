package main

// github.com/lib/pq
import (
	sever "example/hello"
	"example/hello/pkg/handler"
	"example/hello/pkg/repository"
	"example/hello/pkg/repository/db"
	"example/hello/pkg/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	//migrate -path ./schemas -database 'mysql://homestead:secret@tcp(localhost:3306)/my_custom_db' down
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := db.NewMysqlDb(db.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.name"),
	})
	if err != nil {
		log.Fatalf("error initializing db: %s", err.Error())
	}

	// repository -> работа с бд
	// service -> сама логика
	// handlers -> обработчики (на данный момент роуты)
	repos := repository.NewRepository(db)    // в зависимости от env, работаем с нужной бд
	services := service.NewService(repos)    // передаём в сервисы выбранную бд
	handlers := handler.NewHandler(services) // обработчик выбранного сервиса

	srv := new(sever.Server)
	port := viper.GetInt("port")

	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
