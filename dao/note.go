package dao

import (
	"database/sql"

	daoConnection "github.com/AmonFla/simple-task-manager-api/dao/connection"
	"github.com/AmonFla/simple-task-manager-api/models"
)

type NoteDao struct {
	DB *sql.DB
}

func NewNote() *NoteDao {
	dao := new(NoteDao)
	dao.DB = daoConnection.FactoryDao()
	return dao
}

func (dao *NoteDao) GetNote(note *models.Note) error {
	return dao.DB.QueryRow("SELECT comment, user_id, task_id, created_at FROM notes WHERE id=$1",
		note.ID).Scan(&note.Comment, &note.UserId, &note.TaskId, &note.CreatedAt)
}

func (dao *NoteDao) UpdateNote(note *models.Note) error {
	_, err := dao.DB.Exec("UPDATE notes SET comment=$1, user_id=$2, task_id=$3, created_at=$4 WHERE id=$5",
		note.Comment, note.UserId, note.TaskId, note.CreatedAt, note.ID)
	return err
}

func (dao *NoteDao) DeleteNote(note *models.Note) error {
	_, err := dao.DB.Exec("DELETE FROM notes WHERE id=$1", note.ID)
	return err
}

func (dao *NoteDao) CreateNote(note *models.Note) error {
	err := dao.DB.QueryRow(
		"INSERT INTO notes (comment, user_id, task_id, created_at) VALUES($1,$2,$3,$4) RETURNING id",
		note.Comment, note.UserId, note.TaskId, note.CreatedAt).Scan(&note.ID)

	if err != nil {
		return err
	}

	return nil
}

func (dao *NoteDao) GetNotes() ([]models.Note, error) {
	rows, err := dao.DB.Query("SELECT id, comment, created_at, user_id,task_id FROM notes")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	states := []models.Note{}

	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.ID, &note.Comment, &note.CreatedAt, &note.UserId, &note.TaskId); err != nil {
			return nil, err
		}
		states = append(states, note)
	}

	return states, nil
}
