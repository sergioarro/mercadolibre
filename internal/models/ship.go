package models

type Ship struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}
