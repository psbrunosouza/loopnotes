package repositories

import (
	"loop-notes-api/internal/entities"

	"gorm.io/gorm"
)

type TaskRepository interface {
	List(tasks *[]entities.Task) *gorm.DB
	Create(task *entities.Task) *gorm.DB
	// Delete(task *entities.Task) *gorm.DB
	Find(task *entities.Task) *gorm.DB
	// Update(task *entities.Task) * gorm.DB
}

type repository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
} 

func (repository *repository) List(tasks *[]entities.Task) *gorm.DB {
	result := repository.db.Find(tasks)
	if result.Error != nil {
		return result
	}
	return nil
}

func (repository *repository) Create(task *entities.Task) *gorm.DB {
	if result := repository.db.Create(task); result.Error != nil {
		return result
	}
	return nil
}


func (repository *repository) Find(task *entities.Task) *gorm.DB {
	if result := repository.db.First(task); result.Error != nil {
		return result
	}
	return nil
}