package domain

import (
	"github.com/google/uuid"
	"time"
)

type Profile struct {
	Uuid      uuid.UUID `json:"uuid"`
	Login     string    `json:"login"`
	Passwd    string    `json:"passwd"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
