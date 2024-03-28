package db

import (
	"4th_Assignment/ent"
	"log"
	"os"

	"entgo.io/ent/dialect"
	_ "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

var DBConn *ent.Client

func ConnectDb() {
	dsn := "postgres://user_demo:postgres@localhost:5432/portfolio?sslmode=disable"
	//dsn := "postgres://postgres:root@localhost:5432/postgres?sslmode=disable"
	client, err := ent.Open(dialect.Postgres, dsn)
	if err != nil {
		log.Fatal("Failed opening connection to postgres: ", err)
		os.Exit(1)
	}
	log.Println("Connected to database")
	DBConn = client

}
