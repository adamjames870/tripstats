-- name: SaveLocation :one
INSERT INTO locations (id, location_id, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetLocationFromId :one
SELECT * FROM location_info
WHERE id = $1;

-- name: ResetLocations :exec
DELETE FROM locations;

-- name: ResetLocationInfo :exec
DELETE FROM location_info;