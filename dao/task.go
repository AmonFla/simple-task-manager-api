package dao

import (
	"database/sql"

	daoConnection "github.com/AmonFla/simple-task-manager-api/dao/connection"
	"github.com/AmonFla/simple-task-manager-api/models"
)

type TaskDao struct {
	DB *sql.DB
}

func NewTask() *TaskDao {
	dao := new(TaskDao)
	dao.DB = daoConnection.FactoryDao()
	return dao
}

func (dao *TaskDao) GetTask(tsk *models.Task) error {
	return dao.DB.QueryRow("SELECT title, description, project_id, user_id, created_at, updated_at FROM Task WHERE id=$1",
		tsk.ID).Scan(&tsk.Title, &tsk.Description, &tsk.ProjectId, &tsk.UserId, &tsk.CreatedAt, &tsk.UpdatedAt)
}

func (dao *TaskDao) UpdateTask(tsk *models.Task) error {
	_, err := dao.DB.Exec("UPDATE Task SET title=$1, description=$2, project_id=$3, user_id=$4,  updated_at=$5 WHERE id=$6",
		tsk.Title, tsk.Description, tsk.ProjectId, tsk.UserId, tsk.UpdatedAt, tsk.ID)
	return err
}

func (dao *TaskDao) DeleteTask(tsk *models.Task) error {
	_, err := dao.DB.Exec("DELETE FROM Task WHERE id=$1", tsk.ID)
	return err
}

func (dao *TaskDao) CreateTask(tsk *models.Task) error {
	err := dao.DB.QueryRow(
		"INSERT INTO Task (title, description, project_id, user_id, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6) RETURNING id",
		tsk.Title, tsk.Description, tsk.ProjectId, tsk.UserId, tsk.CreatedAt, tsk.UpdatedAt).Scan(&tsk.ID)

	if err != nil {
		return err
	}

	if tsk.States.ID > 0 {
		_, err = dao.DB.Exec("INSERT INTO task_task_state (state_id, task_id, user_id) VALUES($1,$2,$3)", tsk.States.ID, tsk.ID, 1)
	}

	if err != nil {
		return err
	}

	return nil
}

func (dao *TaskDao) GetTasks() ([]models.Task, error) {
	rows, err := dao.DB.Query("SELECT id, title FROM Task")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	states := []models.Task{}

	for rows.Next() {
		var tsk models.Task
		if err := rows.Scan(&tsk.ID, &tsk.Title); err != nil {
			return nil, err
		}
		states = append(states, tsk)
	}

	return states, nil
}

func (dao *TaskDao) AddStateToTask(ID string, state string) error {

	_, err := dao.DB.Exec("INSERT INTO task_task_state (state_id, task_id, user_id) VALUES($1,$2,$3)", ID, state, 1)
	if err != nil {
		return err
	}

	return nil
}
