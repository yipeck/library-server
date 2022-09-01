package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open database
	db, err := sql.Open("sqlite3", "./library.sqlite3")

	if err != nil {
		log.Fatal(err)
	}

	InitDB(db)
	defer db.Close()

	// Initializes router
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	InitRouter(s, db)
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(ResponseWriterMiddleware)

	log.Fatal(http.ListenAndServe(":80", r))
}

// Initialize database
func InitDB(db *sql.DB) {
	sql, err := os.ReadFile("./library.sql")

	if err != nil {
		log.Fatal(err)
	}

	db.Exec(string(sql))
}

// ResponseWriter Middleware
func ResponseWriterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,OPTIONS,DELETE")

		next.ServeHTTP(w, r)
	})
}
