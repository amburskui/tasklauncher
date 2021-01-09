package persistence

import (
	"github.com/amburskui/tasklauncher/internal/app/domain"

	"gopkg.in/mgo.v2"
)

// TaskRepository struct
type TaskRepository struct {
	session *mgo.Session
}

// NewTaskRepository returns a new TaskRepository instance
func NewTaskRepository(session *mgo.Session) *TaskRepository {
	return &TaskRepository{session}
}

// Insert task
func (repo *TaskRepository) Insert(task domain.Task) error {
	err := repo.session.Clone().DB("test").C("tasks").Insert(task)
	return err
}

func (repo *TaskRepository) NewTask() (*task, error) {
	return nil, nil
}
