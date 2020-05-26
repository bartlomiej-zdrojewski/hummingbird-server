package main

import(
    "fmt"
    "net/http"
    "database/sql"
    "github.com/gorilla/mux"
)

type context struct {
    db *sql.DB
    sessions map[string]session
}

func main() {
    ctx := context{}
    defer ctx.closeDatabase()

    r := mux.NewRouter()
    api := r.PathPrefix("/api").Subrouter()
    api.HandleFunc("/login", ctx.handleLoginRequest).Methods(http.MethodPost)
    api.HandleFunc("/register", ctx.handleRegisterRequest).Methods(http.MethodPost)

    err := http.ListenAndServe(":8080", r)
    if err != nil {
        fmt.Print(err.Error())
        // TODO log
        return
    }
}
