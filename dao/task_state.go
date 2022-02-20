package dao

import (
	"database/sql"

	daoConnection "github.com/AmonFla/simple-task-manager-api/dao/connection"
	"github.com/AmonFla/simple-task-manager-api/models"
)

type TaskStateDao struct {
	DB *sql.DB
}

func NewTaskState() *TaskStateDao {
	dao := new(TaskStateDao)
	dao.DB = daoConnection.FactoryDao()
	return dao
}

func (dao *TaskStateDao) GetTaskState(st *models.TaskState) error {
	return dao.DB.QueryRow("SELECT name FROM task_state WHERE id=$1", st.ID).Scan(&st.Name)
}

func (dao *TaskStateDao) UpdateTaskState(st *models.TaskState) error {
	_, err := dao.DB.Exec("UPDATE task_state SET name=$1 WHERE id=$2", st.Name, st.ID)
	return err
}

func (dao *TaskStateDao) DeleteTaskState(st *models.TaskState) error {
	_, err := dao.DB.Exec("DELETE FROM task_state WHERE id=$1", st.ID)
	return err
}

func (dao *TaskStateDao) CreateTaskState(st *models.TaskState) error {
	err := dao.DB.QueryRow(
		"INSERT INTO task_state (name) VALUES($1) RETURNING id",
		st.Name).Scan(&st.ID)

	if err != nil {
		return err
	}

	return nil
}

func (dao *TaskStateDao) GetTaskStates() ([]models.TaskState, error) {
	rows, err := dao.DB.Query("SELECT id, name FROM task_state")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	states := []models.TaskState{}

	for rows.Next() {
		var st models.TaskState
		if err := rows.Scan(&st.ID, &st.Name); err != nil {
			return nil, err
		}
		states = append(states, st)
	}

	return states, nil
}
