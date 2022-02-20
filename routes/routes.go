package routes

import (
	"github.com/AmonFla/simple-task-manager-api/controllers"
	"github.com/gorilla/mux"
)

func AddRoutes(s *mux.Router) {
	// Configuro cada una de las rutas y metodos que aceptan
	s.HandleFunc("/keepalive", controllers.KeepAlive).Methods("GET")

	// TaskState
	s.HandleFunc("/taskstate", controllers.PostTaskState).Methods("POST")
	s.HandleFunc("/taskstate", controllers.GetAllTaskState).Methods("GET")
	s.HandleFunc("/taskstate/{ID:[0-9]+}", controllers.GetTaskState).Methods("GET")
	s.HandleFunc("/taskstate/{ID:[0-9]+}", controllers.PutTaskState).Methods("PUT")
	s.HandleFunc("/taskstate/{ID:[0-9]+}", controllers.DeleteTaskState).Methods("DELETE")

	// ProjectState
	s.HandleFunc("/projectstate", controllers.PostProjectState).Methods("POST")
	s.HandleFunc("/projectstate", controllers.GetAllProjectState).Methods("GET")
	s.HandleFunc("/projectstate/{ID:[0-9]+}", controllers.GetProjectState).Methods("GET")
	s.HandleFunc("/projectstate/{ID:[0-9]+}", controllers.PutProjectState).Methods("PUT")
	s.HandleFunc("/projectstate/{ID:[0-9]+}", controllers.DeleteProjectState).Methods("DELETE")
}
