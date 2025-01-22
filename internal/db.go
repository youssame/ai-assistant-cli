package internal

import (
	"database/sql"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var ElasticsearchClient *elasticsearch.Client = nil

func db() *sql.DB {
	db, err := sql.Open("sqlite3", os.Getenv("ASSISTANT_DB_HOST"))
	if err != nil {
		log.Fatalln("Error while connecting to the DB")
	}
	return db
}

// Query get the result of query
func Query(query string, args ...any) *sql.Rows {
	database := db()
	defer database.Close()
	res, err := database.Query(query, args)
	if err != nil {
		fmt.Println(query, args)
		log.Fatalln("Error while exec a query in the DB", err)
	}
	return res
}

func init() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	ElasticsearchClient = es
}
