-- name: GetLocationFromId :one
SELECT * FROM location_info
WHERE id = $1;