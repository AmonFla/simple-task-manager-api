//mian.go

package main

import (
	"fmt"
	"net/http"

	"github.com/AmonFla/simple-task-manager-api/controllers"
	"github.com/AmonFla/simple-task-manager-api/utils"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	s := router.PathPrefix("/api").Subrouter()

	// Configuro cada una de las rutas y metodos que aceptan
	s.HandleFunc("/keepalive", controllers.KeepAlive).Methods("GET")

	// TaskState
	s.HandleFunc("/task_state", controllers.PostTaskState).Methods("POST")
	s.HandleFunc("/task_state", controllers.GetAllTaskState).Methods("GET")
	s.HandleFunc("/task_state/{ID:[0-9]+}", controllers.GetTaskState).Methods("GET")
	s.HandleFunc("/task_state/{ID:[0-9]+}", controllers.PutTaskState).Methods("PUT")
	s.HandleFunc("/task_state/{ID:[0-9]+}", controllers.DeleteTaskState).Methods("DELETE")

	//Inicializo el server
	port := utils.GoDotEnvVariable("PORT")
	fmt.Printf("Starting server: " + port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic(err)
	}
}
