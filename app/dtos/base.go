package dto

import (
	"time"

	"main.go/app/models"
)

type BaseResponse struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Age       int        `json:"age"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type BaseRequest struct {
	Name string `json:"name" validate:"required,min=3"`
	Age  int    `json:"age" validate:"required,min=0,max=999"`
}

func ToBase(u *models.Base) *BaseResponse {
	return &BaseResponse{
		ID:        u.ID,
		Name:      u.Name,
		Age:       u.Age,
		CreatedAt: u.CreatedAt,
		UpdatedAt: &u.UpdatedAt,
	}
}

func ToBases(u []*models.Base) []*BaseResponse {
	res := make([]*BaseResponse, len(u))
	for i, data := range u {
		res[i] = ToBase(data)
	}
	return res
}
