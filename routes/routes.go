package routes

import (
	"github.com/AmonFla/simple-task-manager-api/controllers"
	"github.com/gorilla/mux"
)

func AddRoutes(s *mux.Router) {
	// Configuro cada una de las rutas y metodos que aceptan
	s.HandleFunc("/keepalive", controllers.KeepAlive).Methods("GET")
	// TaskState
	controllers.NewTaskStateController(s)
	// ProjectState
	controllers.NewProjectStateController(s)
	//Project
	controllers.NewProjectController(s)
	//Task
	controllers.NewTaskController(s)
}
