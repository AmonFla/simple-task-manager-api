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

type ProjectController struct {
	dao   *dao.ProjectDao
	model models.Project
}

func NewProjectController(s *mux.Router) *ProjectController {

	controller := new(ProjectController)
	controller.dao = dao.NewProject()
	//adding routes
	// Project
	s.HandleFunc("/project", controller.PostProject).Methods("POST")
	s.HandleFunc("/project", controller.GetAllProject).Methods("GET")
	s.HandleFunc("/project/{ID:[0-9]+}", controller.GetProject).Methods("GET")
	s.HandleFunc("/project/{ID:[0-9]+}", controller.PutProject).Methods("PUT")
	s.HandleFunc("/project/{ID:[0-9]+}", controller.DeleteProject).Methods("DELETE")

	return controller
}

func (pr *ProjectController) PostProject(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pr.model); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := pr.dao.CreateProject(&pr.model); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, pr.model)
}

func (pr *ProjectController) GetAllProject(w http.ResponseWriter, r *http.Request) {
	data, err := pr.dao.GetProjects()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)

}

func (pr *ProjectController) GetProject(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &pr.model.ID)
	err := pr.dao.GetProject(&pr.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, &pr.model)
}

func (pr *ProjectController) DeleteProject(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &pr.model.ID)
	err := pr.dao.DeleteProject(&pr.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithMoMessage(w, http.StatusAccepted)
}

func (pr *ProjectController) PutProject(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &pr.model.ID)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pr.model); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err := pr.dao.UpdateProject(&pr.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, &pr.model)
}