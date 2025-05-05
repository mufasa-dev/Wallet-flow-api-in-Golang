package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Historic struct {
	gorm.Model
	Action  string
	Comment string
	UserId  uint
}

type HistoricResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Action    string    `json:"action"`
	Comment   string    `json:"comment"`
	UserId    uint      `json:"user_id"`
}
