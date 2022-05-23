package repository

const (
	getSatelliteByNameQuery = `SELECT COUNT(name) FROM satellite WHERE name = $1`
	createSatellite         = `INSERT INTO satellite (name, message, distance, position) VALUES ($1, $2, $3, $4) RETURNING *`
	getSatellites           = `SELECT name, message, distance, position FROM satellite`
)
