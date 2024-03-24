package model

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	ID            uuid.UUID
	Username      string
	Password      string
	Email         string
	EmailVerified bool
	Phone         string
	Created       time.Time
	Updated       time.Time
}
