package router

import (
    "go-code-quest/internal/interface/http/handler"
    //"net/http"

    "github.com/gorilla/mux"
)

func NewRouter(userHandler *handler.UserHandler) *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
    r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
    return r
}
