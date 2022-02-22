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

type TaskStateController struct {
	dao   *dao.TaskStateDao
	model models.TaskState
}

func NewTaskStateController(s *mux.Router) *TaskStateController {

	controller := new(TaskStateController)
	controller.dao = dao.NewTaskState()
	//adding routes
	// Project
	s.HandleFunc("/taskstate", controller.PostTaskState).Methods("POST")
	s.HandleFunc("/taskstate", controller.GetAllTaskState).Methods("GET")
	s.HandleFunc("/taskstate/{ID:[0-9]+}", controller.GetTaskState).Methods("GET")
	s.HandleFunc("/taskstate/{ID:[0-9]+}", controller.PutTaskState).Methods("PUT")
	s.HandleFunc("/taskstate/{ID:[0-9]+}", controller.DeleteTaskState).Methods("DELETE")

	return controller
}

func (st *TaskStateController) PostTaskState(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&st.model); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := st.dao.CreateTaskState(&st.model); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, st.model)
}

func (st *TaskStateController) GetAllTaskState(w http.ResponseWriter, r *http.Request) {
	data, err := st.dao.GetTaskStates()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)

}

func (st *TaskStateController) GetTaskState(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &st.model.ID)
	err := st.dao.GetTaskState(&st.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, &st.model)
}

func (st *TaskStateController) DeleteTaskState(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &st.model.ID)
	err := st.dao.DeleteTaskState(&st.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithMoMessage(w, http.StatusAccepted)
}

func (st *TaskStateController) PutTaskState(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &st.model.ID)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&st.model); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err := st.dao.UpdateTaskState(&st.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, &st.model)
}
