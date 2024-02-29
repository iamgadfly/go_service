package models

import "time"

type Project struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required,min=1,max=30"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreateProjectReq struct {
	Name string `json:"name"`
}
