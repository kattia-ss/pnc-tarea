package repositories

import (
	"layersapi/entities"
)

// Contiene la interfaz UserRepository. Define los métodos esperados:
// GetAll, GetById, Create, Update y Delete.

type UserRepository interface {
	GetAll() ([]entities.User, error)
	GetById(id string) (entities.User, error)
	Create(user entities.User) error
	Update(id, name, email string) error
	Delete(id string) error
}
