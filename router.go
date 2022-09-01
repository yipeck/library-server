package main

import (
	"database/sql"

	"github.com/gorilla/mux"

	"github.com/yipeck/library-module/author"
	"github.com/yipeck/library-module/book"
	"github.com/yipeck/library-module/category"
	"github.com/yipeck/library-module/country"
	"github.com/yipeck/library-module/publisher"
	"github.com/yipeck/library-module/user"
)

// Initialize Site Router
func InitRouter(r *mux.Router, db *sql.DB) {
	AuthorRouter(r, db)
	BookRouter(r, db)
	CategoryRouter(r, db)
	CountryRouter(r, db)
	PublisherRouter(r, db)
	UserRouter(r, db)
}

// Author Router
func AuthorRouter(r *mux.Router, db *sql.DB) {
	s := r.PathPrefix("/author").Subrouter()

	s.HandleFunc("", author.FetchAuthors(db)).Methods("GET")
	s.HandleFunc("", author.AddAuthor(db)).Methods("POST")
	s.HandleFunc("/update", author.UpdateAuthor(db)).Methods("POST")
	s.HandleFunc("", author.DeleteAuthor(db)).Methods("OPTIONS", "DELETE")
	s.HandleFunc("/count/{cid}", author.CountByCountry(db)).Methods("GET")
}

// Book Router
func BookRouter(r *mux.Router, db *sql.DB) {
	s := r.PathPrefix("/book").Subrouter()

	s.HandleFunc("/bought", book.FetchBooksBoughtThisMonth(db)).Methods("GET")
	s.HandleFunc("/{type}/{value}", book.FetchBooksByType(db)).Methods("GET")
	s.HandleFunc("/{id}", book.FetchBookById(db)).Methods("GET")
	s.HandleFunc("", book.AddBook(db)).Methods("POST")
	s.HandleFunc("/{id}", book.DeleteBook(db)).Methods("OPTIONS", "DELETE")
	s.HandleFunc("/status", book.UpdateReadStatus(db)).Methods("PUT")
	s.HandleFunc("/count/{type}/{value}", book.CountBooksByType(db)).Methods("GET")
}

// Category Router
func CategoryRouter(r *mux.Router, db *sql.DB) {
	s := r.PathPrefix("/category").Subrouter()

	s.HandleFunc("", category.FetchCategories(db)).Methods("GET")
	s.HandleFunc("", category.AddCategory(db)).Methods("POST")
	s.HandleFunc("/attr", category.UpdateCategory(db)).Methods("POST")
	s.HandleFunc("", category.DeleteCategory(db)).Methods("OPTIONS", "DELETE")
}

// Country Router
func CountryRouter(r *mux.Router, db *sql.DB) {
	s := r.PathPrefix("/country").Subrouter()

	s.HandleFunc("", country.FetchCountries(db)).Methods("GET")
	s.HandleFunc("", country.AddCountry(db)).Methods("POST")
	s.HandleFunc("/attr", country.UpdateCountry(db)).Methods("POST")
	s.HandleFunc("", country.DeleteCountry(db)).Methods("OPTIONS", "DELETE")
}

// Publisher Router
func PublisherRouter(r *mux.Router, db *sql.DB) {
	s := r.PathPrefix("/publisher").Subrouter()

	s.HandleFunc("", publisher.FetchPublisheres(db)).Methods("GET")
	s.HandleFunc("", publisher.AddPublisher(db)).Methods("POST")
	s.HandleFunc("/attr", publisher.UpdatePublisher(db)).Methods("POST")
	s.HandleFunc("", publisher.DeletePublisher(db)).Methods("OPTIONS", "DELETE")
}

// User Router
func UserRouter(r *mux.Router, db *sql.DB) {
	s := r.PathPrefix("/user").Subrouter()

	s.HandleFunc("/signin", user.SignIn(db)).Methods("POST")
	s.HandleFunc("/signup", user.SignUp(db)).Methods("POST")
	s.HandleFunc("", user.FetchUserById(db)).Methods("GET")
	s.HandleFunc("/exist", user.IsExist(db)).Methods("POST")
	s.HandleFunc("/attr", user.UpdateUser(db)).Methods("POST")
}
