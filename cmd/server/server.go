package main

import (
	"GoNews/configs"
	"GoNews/pkg/api"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/memdb"
	"GoNews/pkg/storage/mongo"
	"GoNews/pkg/storage/postgresql"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Инициализация конфига для получения URL БД
	if err := configs.InitConfig(); err != nil {
		log.Fatalf("error initialization: %s", err.Error())
	}
	postgresURL := fmt.Sprintf("postgresql://%s:%s@%s:%v/%s", viper.Get("postgres.username"), viper.Get("postgres.password"), viper.Get("postgres.host"), viper.Get("postgres.port"), viper.Get("postgres.database"))
	mongoURL := fmt.Sprintf("mongodb://%s:%s@%s:%v", viper.Get("mongo.username"), viper.Get("mongo.password"), viper.Get("mongo.host"), viper.Get("mongo.port"))
	// Создаём объект сервера.
	var srv server

	// Создаём объекты баз данных.
	//
	// БД в памяти.
	db := memdb.New()

	// Реляционная БД PostgreSQL.
	db2, err := postgresql.New(postgresURL)
	if err != nil {
		log.Fatal(err)
	}
	// Документная БД MongoDB.
	db3, err := mongo.New(mongoURL)
	if err != nil {
		log.Fatal(err)
	}
	_, _, _ = db, db2, db3

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db2

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	// Предаём серверу маршрутизатор запросов,
	// поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	http.ListenAndServe(":8080", srv.api.Router())
}
