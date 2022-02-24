package dao

import (
	"database/sql"

	daoConnection "github.com/AmonFla/simple-task-manager-api/dao/connection"
	"github.com/AmonFla/simple-task-manager-api/models"
)

type ProjectDao struct {
	DB *sql.DB
}

func NewProject() *ProjectDao {
	dao := new(ProjectDao)
	dao.DB = daoConnection.FactoryDao()
	return dao
}

func (dao *ProjectDao) GetProject(pr *models.Project) error {
	dao.DB.QueryRow("SELECT name, description, created_at, updated_at, closed_at FROM project WHERE id=$1",
		pr.ID).Scan(&pr.Name, &pr.Description, &pr.CreatedAt, &pr.UpdatedAt, &pr.ClosedAt)
	//TODO add state status and users list
	return nil
}

func (dao *ProjectDao) UpdateProject(pr *models.Project) error {
	_, err := dao.DB.Exec("UPDATE project SET name=$1, description=$2, created_at=$3, updated_at=$4, closed_at=$5 WHERE id=$6",
		pr.Name, pr.Description, pr.CreatedAt, pr.UpdatedAt, pr.ClosedAt, pr.ID)
	return err
}

func (dao *ProjectDao) DeleteProject(pr *models.Project) error {
	_, err := dao.DB.Exec("DELETE FROM project WHERE id=$1", pr.ID)
	return err
}

func (dao *ProjectDao) CreateProject(pr *models.Project) error {
	err := dao.DB.QueryRow(
		"INSERT INTO project (name, description, created_at, updated_at, closed_at) VALUES($1,$2,$3,$4,$5) RETURNING id",
		pr.Name, pr.Description, pr.CreatedAt, pr.UpdatedAt, pr.ClosedAt).Scan(&pr.ID)

	if err != nil {
		return err
	}

	if pr.States.ID > 0 {
		_, err = dao.DB.Exec("INSERT INTO project_project_state (state_id, project_id, user_id) VALUES($1,$2,$3)", pr.States.ID, pr.ID, 1)
	}

	if err != nil {
		return err
	}

	for _, user := range pr.Users {
		_, err = dao.DB.Exec("INSERT INTO project_user (project_id, user_id, active) VALUES($1,$2,TRUE)", pr.ID, user)
	}

	if err != nil {
		return err
	}

	return nil
}

func (dao *ProjectDao) GetProjects() ([]models.Project, error) {
	rows, err := dao.DB.Query("SELECT id, name FROM project")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	states := []models.Project{}

	for rows.Next() {
		var pr models.Project
		if err := rows.Scan(&pr.ID, &pr.Name); err != nil {
			return nil, err
		}
		states = append(states, pr)
	}

	return states, nil
}

func (dao *ProjectDao) AddUserToProject(pr *models.Project) error {

	for _, user := range pr.Users {
		_, err := dao.DB.Exec("INSERT INTO project_user (project_id, user_id, active) VALUES($1,$2,TRUE) ON CONFLICT DO NOTHING", pr.ID, user)
		if err != nil {
			return err
		}
	}
	//TODO add state status and users list

	return nil
}

func (dao *ProjectDao) DeleteUserFromProject(ID string, user string) error {

	_, err := dao.DB.Exec("DELETE FROM project_user WHERE project_id = $1 AND user_id = $2 ", ID, user)
	if err != nil {
		return err
	}

	return nil
}

func (dao *ProjectDao) ActiveUserFromProject(ID string, user string) error {

	_, err := dao.DB.Exec("UPDATE  project_user SET active = not active WHERE project_id = $1 AND user_id = $2 ", ID, user)
	if err != nil {
		return err
	}

	return nil
}

func (dao *ProjectDao) AddStateToProject(ID string, state string) error {

	_, err := dao.DB.Exec("INSERT INTO project_project_state (state_id, project_id, user_id) VALUES($1,$2,$3)", ID, state, 1)
	if err != nil {
		return err
	}

	return nil
}
