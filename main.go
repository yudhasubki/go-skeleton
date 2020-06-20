package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/yudhasubki/go-skeleton/server"

	"github.com/yudhasubki/go-skeleton/container"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yudhasubki/go-skeleton/router"

	conf "github.com/yudhasubki/go-skeleton/config"
	"github.com/yudhasubki/go-skeleton/domain/example"
)

func SetupDatabase(conf conf.Config) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", conf.DatabaseUsername, conf.DatabasePassword, conf.DatabaseHost, conf.DatabaseName))
	if err != nil {
		log.Fatalf("%v", err)
	}
	return db
}

func main() {
	ct := container.NewContainer()
	conf := conf.Get()
	db := SetupDatabase(*conf)
	ct.Register("config", conf)
	ct.Register("database", db)
	ct.Register("repository", new(example.Repository))
	ct.Register("service", new(example.Service))
	ct.Register("handler", new(example.Handler))

	routerHandler := router.RouterHandler{}
	ct.Register("router", routerHandler)
	err := ct.Start()
	if err != nil {
		log.Fatalf("error starting : %v", err)
	}

	serve := server.Server{Config: conf, Router: &routerHandler, DB: db}
	serve.EnableGracefulShutdown(*ct)
	serve.Serve()
}
