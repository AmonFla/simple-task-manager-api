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

type ProjectStateController struct {
	dao   *dao.ProjectStateDao
	model models.ProjectState
}

func NewProjectStateController(s *mux.Router) *ProjectStateController {

	controller := new(ProjectStateController)
	controller.dao = dao.NewProjectState()
	//adding routes
	// Project
	s.HandleFunc("/project", controller.PostProjectState).Methods("POST")
	s.HandleFunc("/project", controller.GetAllProjectState).Methods("GET")
	s.HandleFunc("/project/{ID:[0-9]+}", controller.GetProjectState).Methods("GET")
	s.HandleFunc("/project/{ID:[0-9]+}", controller.PutProjectState).Methods("PUT")
	s.HandleFunc("/project/{ID:[0-9]+}", controller.DeleteProjectState).Methods("DELETE")

	return controller
}

func (pt *ProjectStateController) PostProjectState(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pt.model); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := pt.dao.CreateProjectState(&pt.model); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, pt.model)
}

func (pt *ProjectStateController) GetAllProjectState(w http.ResponseWriter, r *http.Request) {
	data, err := pt.dao.GetProjectStates()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)

}

func (pt *ProjectStateController) GetProjectState(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &pt.model.ID)
	err := pt.dao.GetProjectState(&pt.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, &pt.model)
}

func (pt *ProjectStateController) DeleteProjectState(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &pt.model.ID)
	err := pt.dao.DeleteProjectState(&pt.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithMoMessage(w, http.StatusAccepted)
}

func (pt *ProjectStateController) PutProjectState(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &pt.model.ID)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pt.model); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err := pt.dao.UpdateProjectState(&pt.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, &pt.model)
}
