package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"tor_search/src/controller"
	"tor_search/src/repository"
	"tor_search/src/service"
)

func main() {
	fmt.Println("\nСтарт...")

	err := initConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println("Подключение к БД...")
	db, err := repository.ConnectPostgres(repository.DbConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Name:     viper.GetString("db.name"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		SslMode:  viper.GetString("db.ssl_mode"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Успешно подключено!")

	db.SetMaxOpenConns(2 * runtime.NumCPU())

	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println(fmt.Sprintf("Не удалось закрыть подключение к бд: %s", err))
		}
	}()

	r := repository.NewRepository(db)
	s := service.NewService(r)
	c := controller.NewConsoleController(s)

	go c.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	fmt.Println("Завершение работы...")
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("app-config")
	return viper.ReadInConfig()
}
