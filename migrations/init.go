package main

import (
	"github.com/shobhitsharma/go-gql-starter/db"
	"github.com/shobhitsharma/go-gql-starter/model"
)

func main() {
	d, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer d.Close()

	d.DropTableIfExists(&model.User{})
	d.CreateTable(&model.User{})
}
