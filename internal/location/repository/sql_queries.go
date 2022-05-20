package repository

const (
	getSatelliteByNameQuery = `SELECT name, message, position, distance FROM satellite WHERE name = $1`
)
