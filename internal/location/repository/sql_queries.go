package repository

const (
	getLocationBySatellites = `SELECT n.satellite_id,
       n.name,
       n.position
FROM satellite n
WHERE satellite_id = $1`
)
