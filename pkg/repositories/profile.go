package repositories

import (
	"github.com/goblinus/http-service/pkg/domain"
	"github.com/google/uuid"
)

type Profile interface {
	FindByUUID(profileUUID uuid.UUID) (domain.Profile, error)
	ValidateByPasswdHash(profileUUID uuid.UUID, hash string) bool
	Create(profile domain.Profile) (uuid.UUID, error)
	Delete(profileUUID uuid.UUID) (domain.Profile, error)
}
