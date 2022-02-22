package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/AmonFla/simple-task-manager-api/dao"
	"github.com/AmonFla/simple-task-manager-api/models"
	"github.com/AmonFla/simple-task-manager-api/utils"
)

type TaskController struct {
	dao   *dao.TaskDao
	model models.Task
}

func NewTaskController(s *mux.Router) *TaskController {

	controller := new(TaskController)
	controller.dao = dao.NewTask()
	//adding routes
	// Task
	s.HandleFunc("/task", controller.PostTask).Methods("POST")
	s.HandleFunc("/task", controller.GetAllTask).Methods("GET")
	s.HandleFunc("/task/{ID:[0-9]+}", controller.GetTask).Methods("GET")
	s.HandleFunc("/task/{ID:[0-9]+}", controller.PutTask).Methods("PUT")
	s.HandleFunc("/task/{ID:[0-9]+}", controller.DeleteTask).Methods("DELETE")

	return controller
}

func (tr *TaskController) PostTask(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tr.model); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := tr.dao.CreateTask(&tr.model); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, tr.model)
}

func (tr *TaskController) GetAllTask(w http.ResponseWriter, r *http.Request) {
	data, err := tr.dao.GetTasks()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)

}

func (tr *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &tr.model.ID)
	err := tr.dao.GetTask(&tr.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, &tr.model)
}

func (tr *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &tr.model.ID)
	err := tr.dao.DeleteTask(&tr.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithMoMessage(w, http.StatusAccepted)
}

func (tr *TaskController) PutTask(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &tr.model.ID)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tr.model); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err := tr.dao.UpdateTask(&tr.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, &tr.model)
}
