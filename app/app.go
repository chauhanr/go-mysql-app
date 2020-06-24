package app

import (
    "fmt"
    "database/sql"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)

func InitializeApp() {
    fmt.Println("Connecting to mysql instance customer-db-svc")
    db, err := sql.Open("mysql", "root:mysqlPass@tcp(customer-db-svc:3306)/mysql")
    if err != nil {
        fmt.Printf("Error in connecting to database! %s\n", err)
    }
    r := mux.NewRouter()
    s := NewServer(db, r)
    s.Routes()
    fmt.Printf("Starting Server on Port: 8080") 
    http.ListenAndServe(":8080", s.router)
}
