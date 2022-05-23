package models

import "github.com/lib/pq"

type Satellite struct {
	Name     string         `json:"name" validate:"omitempty"`
	Position Position       `json:"position" validate:"omitempty"`
	Distance float64        `json:"distance" validate:"omitempty"`
	Message  pq.StringArray `json:"message" validate:"omitempty"`
}
