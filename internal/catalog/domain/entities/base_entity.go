package entities

import "github.com/google/uuid"

type BaseEntity struct {
	ID    uuid.UUID
	Title string
}
