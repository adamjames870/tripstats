-- name: SaveReview :one
INSERT INTO reviews (id, created_at, updated_at, tripadvisor_review_id, location_id, published_date,
    tripadvisor_url, tripadvisor_title, tripadvisor_text, rating)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: SaveSubReview :one
INSERT INTO sub_reviews (id, created_at, updated_at, review_id, subrating_name, rating)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetReviewCount :one
SELECT COUNT(*) AS review_count
FROM reviews
WHERE location_id = $1
GROUP BY location_id;