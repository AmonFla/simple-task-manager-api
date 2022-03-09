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

type NoteController struct {
	dao   *dao.NoteDao
	model models.Note
}

func NewNoteController(s *mux.Router) *NoteController {

	controller := new(NoteController)
	controller.dao = dao.NewNote()
	//adding routes
	// Note
	s.HandleFunc("/note", controller.PostNote).Methods("POST")
	s.HandleFunc("/note", controller.GetAllNote).Methods("GET")
	s.HandleFunc("/note/{ID:[0-9]+}", controller.GetNote).Methods("GET")
	s.HandleFunc("/note/{ID:[0-9]+}", controller.PutNote).Methods("PUT")
	s.HandleFunc("/note/{ID:[0-9]+}", controller.DeleteNote).Methods("DELETE")

	return controller
}

func (note *NoteController) PostNote(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&note.model); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()
	note.model.CreatedAt = time.Now()
	if err := note.dao.CreateNote(&note.model); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, note.model)
}

func (note *NoteController) GetAllNote(w http.ResponseWriter, r *http.Request) {
	data, err := note.dao.GetNotes()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)

}

func (note *NoteController) GetNote(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &note.model.ID)
	err := note.dao.GetNote(&note.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, &note.model)
}

func (note *NoteController) DeleteNote(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &note.model.ID)
	err := note.dao.DeleteNote(&note.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithMoMessage(w, http.StatusAccepted)
}

func (note *NoteController) PutNote(w http.ResponseWriter, r *http.Request) {
	//Obtengo las variables definidas en la ruta
	vars := mux.Vars(r)

	fmt.Sscan(vars["ID"], &note.model.ID)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&note.model); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err := note.dao.UpdateNote(&note.model)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, &note.model)
}
