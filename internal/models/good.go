package models

import (
	"time"
)

type Good struct {
	ID           int       `json:"id"`
	ProjectId    int       `json:"project_id" db:"project_id"`
	Name         string    `json:"name" db:"name" validate:"required,min=1,max=30"`
	Description  string    `json:"description" db:"description"`
	Priority     int       `json:"priority" db:"priority"`
	Removed      bool      `json:"removed" db:"removed"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	RemovedCount int       `json:"removed_count" db:"removed_count"`
}

type RemoveResp struct {
	ID         int  `json:"id"`
	CampaginId int  `json:"campaginId"` // видимо id проекта
	Removed    bool `json:"removed"`
}

type GoodList struct {
	Total   int    `json:"total" db:"total"`
	Removed int    `json:"removed" db:"removed_count"`
	Limit   int    `json:"limit" db:"limit"`
	Offset  int    `json:"offset" db:"offset"`
	Goods   []Good `json:"goods"`
}

type PriorityResp struct {
	ID       int `json:"id"`
	Priority int `json:"priority"`
}

type Reprioritiize struct {
	Priorities []PriorityResp `json:"priorities"`
}

type PriorityReq struct {
	NewPriority int `json:"newPriority" validate:"required"`
}

type UpdateGoodReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateGoodReq struct {
	Name string `json:"name"`
}
