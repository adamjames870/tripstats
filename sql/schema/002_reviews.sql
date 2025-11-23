-- +goose Up
CREATE TABLE reviews (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    tripadvisor_review_id int NOT NULL UNIQUE,
    location_id varchar NOT NULL UNIQUE,
    published_date TIMESTAMP NOT NULL,
    tripadvisor_url varchar,
    tripadvisor_title varchar,
    tripadvisor_text varchar,
    rating int NOT NULL,
    CONSTRAINT fk_location_id FOREIGN KEY (location_id) REFERENCES locations(location_id) ON DELETE CASCADE
);

CREATE TABLE sub_reviews (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    review_id uuid NOT NULL,
    subrating_name varchar NOT NULL,
    rating int NOT NULL,
    CONSTRAINT fk_review_id FOREIGN KEY (review_id) REFERENCES reviews(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE sub_reviews;
DROP TABLE reviews;