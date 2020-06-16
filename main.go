package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

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
	ct.Register("config", conf)
	ct.Register("database", SetupDatabase(*conf))
	ct.Register("repository", new(example.Repository))
	ct.Register("service", new(example.Service))
	ct.Register("handler", new(example.Handler))

	routerHandler := router.RouterHandler{}
	ct.Register("router", routerHandler)
	err := ct.Start()
	if err != nil {
		log.Fatalf("error starting : %v", err)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", conf.ServerPort), router.Router(&routerHandler)))
}
