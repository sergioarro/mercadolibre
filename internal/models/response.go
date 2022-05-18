package models

type Response struct {
	Position Position `json:"position"`
	Message  string   `json:"message"`
}
