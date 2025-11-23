-- +goose Up
CREATE TABLE locations (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    location_id varchar NOT NULL UNIQUE,
    name varchar NOT NULL
);

CREATE TABLE location_info (
    id varchar PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name varchar NOT NULL,
    web_url varchar,
    rating float,
    num_reviews int,
    CONSTRAINT fk_location_id FOREIGN KEY (id) REFERENCES locations(location_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE location_info;
DROP TABLE locations;