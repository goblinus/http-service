package repositories

import (
	"github.com/goblinus/http-service/pkg/domain"
	"github.com/google/uuid"
)

type User interface {
	FindUserByUUID(userUUID uuid.UUID) (domain.User, error)
	FindUserByEmail(email string) (domain.User, error)
	Create(user domain.User) uuid.UUID
	Delete(userUUID uuid.UUID) (domain.User, error)
}
