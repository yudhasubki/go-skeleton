package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

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
	conf := conf.Get()
	db := SetupDatabase(*conf)
	handler := example.NewHandler(db, conf)
	rh := router.RouterHandler{Handler: handler}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", conf.ServerPort), router.Router(&rh)))
}
