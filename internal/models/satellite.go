package models

type Satellite struct {
	Name     string   `json:"name" validate:"omitempty"`
	Position Position `json:"position" validate:"omitempty"`
	Distance float64  `json:"distance" validate:"omitempty"`
	Message  []string `json:"message" validate:"omitempty"`
}
