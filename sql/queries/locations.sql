-- name: SaveLocation :one
INSERT INTO locations (id, created_at, updated_at, location_id, name)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetLocationFromId :one
SELECT * FROM location_info
WHERE id = $1;

-- name: ResetLocations :exec
DELETE FROM locations;

-- name: ResetLocationInfo :exec
DELETE FROM location_info;

