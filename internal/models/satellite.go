package models

type Satellite struct {
	Name     string   `json:"name"`
	Position Position `json:"position"`
}
