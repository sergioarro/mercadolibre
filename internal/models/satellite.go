package models

type Satellite struct {
	Name     string   `json:"name" validate:"required"`
	Position Position `json:"position" validate:"omitempty"`
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
}
