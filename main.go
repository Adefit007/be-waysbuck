package main

import (
	"fmt"
	"net/http"
	"waysbuck/database"
	"waysbuck/pkg/mysql"
	"waysbuck/routes"

	"github.com/gorilla/mux"
)

func main() {
	// On http (API)
	r := mux.NewRouter()
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// }).Methods("GET")
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
