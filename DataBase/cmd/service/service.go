package main

import (
	"DataBase/pkg /memdb"
	"DataBase/pkg /postgres"
	"DataBase/pkg /storage"
	"fmt"
	"log"
	"os"
)

var db storage.Interface

func main() {
	var err error
	pwd := os.Detenv("dbpass")
	if pwd == "" {
		os.Exit(1)
	}
	connstr :=
		"postgres://postgres:" +
			pwd + "@ubuntu-server.northeurope.cloudapp.azure.com/tasks"
	db, err = postgres.New(connstr)
	if err != nil {
		log.Fatal(err)
	}
	db = memdb.DB{}
	tasks, err := db.Tasks(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks)
}
