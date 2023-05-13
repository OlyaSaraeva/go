package main

import (
	"database/sql"
	"log"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Импортируем для возможности подключения к MySQL
	"github.com/jmoiron/sqlx"
	"github.com/gorilla/mux"
)

const (
	port = ":3000"
	dbDriverName = "mysql"
)

func main() {
	db, err := openDB() // Открываем соединение к базе данных в самом начале
	if err != nil {
		log.Fatal(err)
	}

	dbx := sqlx.NewDb(db, dbDriverName) // Расширяем стандартный клиент к базе

	mux := mux.NewRouter()
	mux.HandleFunc("/home", index(dbx)) // Передаём клиент к базе данных в ф-ию обработчик запроса

	mux.HandleFunc("/post/{postID}", post(dbx))
	mux.HandleFunc("/api/post", createPost(dbx)).Methods(http.MethodPost)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/admin", admin)
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Println("Start server")
	http.ListenAndServe(port, mux)
}

func openDB() (*sql.DB, error) {
	// Здесь прописываем соединение к базе данных
	return sql.Open(dbDriverName, "root:root@tcp(localhost:3306)/blog?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true")
}

 
