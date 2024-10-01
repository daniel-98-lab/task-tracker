package models

import (
    "time"
	"yask-tracker/internal/enums"
)

type Task struct {
    Id          int       		`json:"id"`
    Description string    		`json:"description"`
    Status      enums.Status    `json:"status"`
    CreatedAt   time.Time 		`json:"created_at"`
    UpdatedAt   time.Time 		`json:"updated_at"`
}
