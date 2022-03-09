package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	s.HandleFunc("/project/{ID:[0-9]+}/users", controller.AddUserToProject).Methods("POST")
	s.HandleFunc("/project/{ID:[0-9]+}/users/{user:[0-9]+}", controller.DeleteUserFromProject).Methods("DELETE")
	s.HandleFunc("/project/{ID:[0-9]+}/users/{user:[0-9]+}/active", controller.ActiveUserFromProject).Methods("PUT")
	s.HandleFunc("/project/{ID:[0-9]+}/state/{state:[0-9]+}", controller.AddStateToProject).Methods("PUT")

	return controller
}

func (pr *ProjectController) PostProject(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pr.model); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()
	pr.model.CreatedAt = time.Now()
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
	pr.model.UpdatedAt = time.Now()
	err := pr.dao.UpdateProject(&pr.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	pr.dao.GetProject(&pr.model)
	utils.RespondWithJSON(w, http.StatusOK, &pr.model)
}

func (pr *ProjectController) AddUserToProject(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &pr.model.ID)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pr.model); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err := pr.dao.AddUserToProject(&pr.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	pr.dao.GetProject(&pr.model)
	utils.RespondWithJSON(w, http.StatusOK, &pr.model)
}

func (pr *ProjectController) DeleteUserFromProject(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)
	err := pr.dao.DeleteUserFromProject(vars["ID"], vars["user"])
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Sscan(vars["ID"], &pr.model.ID)
	pr.dao.GetProject(&pr.model)
	utils.RespondWithJSON(w, http.StatusOK, &pr.model)
}

func (pr *ProjectController) ActiveUserFromProject(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)
	err := pr.dao.ActiveUserFromProject(vars["ID"], vars["user"])
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Sscan(vars["ID"], &pr.model.ID)
	pr.dao.GetProject(&pr.model)
	utils.RespondWithJSON(w, http.StatusOK, &pr.model)
}

func (pr *ProjectController) AddStateToProject(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)
	err := pr.dao.AddStateToProject(vars["ID"], vars["state"])
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Sscan(vars["ID"], &pr.model.ID)
	pr.dao.GetProject(&pr.model)
	utils.RespondWithJSON(w, http.StatusOK, &pr.model)
}
