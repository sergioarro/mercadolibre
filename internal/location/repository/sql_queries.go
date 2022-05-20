package repository

const (
	getSatelliteByNameQuery = `SELECT COUNT(name) FROM satellite WHERE name = $1`
)
