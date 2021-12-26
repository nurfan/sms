package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nurfan/sms/repository"
	route "github.com/nurfan/sms/transport/http"
	db "github.com/nurfan/sms/util/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	// load .env file
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	os.Setenv("FILE_SERVER", "localhost:"+os.Getenv("STATIC_PORT"))
}

func main() {
	conn, err := db.GetSqlxConnection()
	if err != nil {
		log.Fatalf("failed to connect database")
	}
	defer conn.Close()

	// setup repo
	repo := repository.NewRepositoryPsql(conn)
	go func() {
		http.Handle("/storage/", http.StripPrefix("/storage/", http.FileServer(http.Dir("./storage"))))
		log.Fatal(http.ListenAndServe(":"+os.Getenv("STATIC_PORT"), nil))
	}()

	route.Serve(repo)
}
