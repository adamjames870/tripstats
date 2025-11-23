-- name: SaveLocationInfo :one
INSERT INTO location_info (id, created_at, updated_at, name, web_url, rating, num_reviews)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateLocationInfo :one
UPDATE location_info
SET updated_at = $4,
    rating = $2,
    num_reviews = $3
WHERE id = $1
RETURNING *;

