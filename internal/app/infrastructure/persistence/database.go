package persistence

import "gopkg.in/mgo.v2"

// Repositories struct
type Repositories struct {
	mdb            *mgo.Session
	taskRepository *TaskRepository
}

// NewRepositories returns a new instance of a MongoDB repository
func NewRepositories(DbName string) (*Repositories, error) {
	mdb, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		return nil, err
	}
	mdb.SetMode(mgo.Monotonic, true)

	taskRepo := NewTaskRepository(mdb.Clone())
	return &Repositories{taskRepo}, nil
}
