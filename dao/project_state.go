package dao

import (
	"database/sql"

	daoConnection "github.com/AmonFla/simple-task-manager-api/dao/connection"
	"github.com/AmonFla/simple-task-manager-api/models"
)

type ProjectStateDao struct {
	DB *sql.DB
}

func NewProjectState() *ProjectStateDao {
	dao := new(ProjectStateDao)
	dao.DB = daoConnection.FactoryDao()
	return dao
}

func (dao *ProjectStateDao) GetProjectState(pt *models.ProjectState) error {
	return dao.DB.QueryRow("SELECT name FROM project_state WHERE id=$1", pt.ID).Scan(&pt.Name)
}

func (dao *ProjectStateDao) UpdateProjectState(pt *models.ProjectState) error {
	_, err := dao.DB.Exec("UPDATE project_state SET name=$1 WHERE id=$2", pt.Name, pt.ID)
	return err
}

func (dao *ProjectStateDao) DeleteProjectState(pt *models.ProjectState) error {
	_, err := dao.DB.Exec("DELETE FROM project_state WHERE id=$1", pt.ID)
	return err
}

func (dao *ProjectStateDao) CreateProjectState(pt *models.ProjectState) error {
	err := dao.DB.QueryRow(
		"INSERT INTO project_state (name) VALUES($1) RETURNING id",
		pt.Name).Scan(&pt.ID)

	if err != nil {
		return err
	}

	return nil
}

func (dao *ProjectStateDao) GetProjectStates() ([]models.ProjectState, error) {
	rows, err := dao.DB.Query("SELECT id, name FROM project_state")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	states := []models.ProjectState{}

	for rows.Next() {
		var pt models.ProjectState
		if err := rows.Scan(&pt.ID, &pt.Name); err != nil {
			return nil, err
		}
		states = append(states, pt)
	}

	return states, nil
}
