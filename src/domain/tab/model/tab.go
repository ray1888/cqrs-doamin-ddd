package model

import (
	"github.com/google/uuid"
)

type TabOpened struct {
	Guid        uuid.UUID
	TableNumber int
	Waiter      string
}

type OpenTab struct {
	Guid        uuid.UUID
	TableNumber int
	Waiter      string
}
