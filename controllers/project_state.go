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

var pt models.ProjectState

func PostProjectState(w http.ResponseWriter, r *http.Request) {
	dao := dao.NewProjectState()
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pt); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := dao.CreateProjectState(&pt); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, pt)
}

func GetAllProjectState(w http.ResponseWriter, r *http.Request) {
	dao := dao.NewProjectState()
	data, err := dao.GetProjectStates()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)

}

func GetProjectState(w http.ResponseWriter, r *http.Request) {
	dao := dao.NewProjectState()
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &pt.ID)
	err := dao.GetProjectState(&pt)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, &pt)
}

func DeleteProjectState(w http.ResponseWriter, r *http.Request) {
	dao := dao.NewProjectState()
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &pt.ID)
	err := dao.DeleteProjectState(&pt)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithMoMessage(w, http.StatusAccepted)
}

func PutProjectState(w http.ResponseWriter, r *http.Request) {
	dao := dao.NewProjectState()
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &pt.ID)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pt); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err := dao.UpdateProjectState(&pt)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, &pt)
}
