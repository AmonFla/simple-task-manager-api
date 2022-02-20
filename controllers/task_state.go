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

var st models.TaskState

func PostTaskState(w http.ResponseWriter, r *http.Request) {
	dao := dao.NewTaskState()
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&st); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := dao.CreateTaskState(&st); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, st)
}

func GetAllTaskState(w http.ResponseWriter, r *http.Request) {
	dao := dao.NewTaskState()
	data, err := dao.GetTaskStates()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)

}

func GetTaskState(w http.ResponseWriter, r *http.Request) {
	dao := dao.NewTaskState()
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &st.ID)
	err := dao.GetTaskState(&st)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, &st)
}

func DeleteTaskState(w http.ResponseWriter, r *http.Request) {
	dao := dao.NewTaskState()
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &st.ID)
	err := dao.DeleteTaskState(&st)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithMoMessage(w, http.StatusAccepted)
}

func PutTaskState(w http.ResponseWriter, r *http.Request) {
	dao := dao.NewTaskState()
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &st.ID)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&st); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err := dao.UpdateTaskState(&st)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, &st)
}
