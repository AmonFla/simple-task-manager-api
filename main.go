//mian.go

package main

import (
	"fmt"
	"net/http"

	"github.com/AmonFla/simple-task-manager-api/routes"
	"github.com/AmonFla/simple-task-manager-api/utils"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	s := router.PathPrefix("/api").Subrouter()

	routes.AddRoutes(s)

	//Inicializo el server
	port := utils.GoDotEnvVariable("PORT")
	fmt.Printf("Starting server: " + port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic(err)
	}
}
