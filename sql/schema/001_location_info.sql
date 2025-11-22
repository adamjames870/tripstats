-- +goose Up
CREATE TABLE locations (
   id uuid PRIMARY KEY,
   location_id varchar NOT NULL UNIQUE,
   name varchar NOT NULL
);

CREATE TABLE location_info (
    id varchar PRIMARY KEY,
    name varchar NOT NULL,
    web_url varchar,
    rating float,
    num_reviews int,
    CONSTRAINT fk_location_id FOREIGN KEY (id) REFERENCES locations(location_id)
);

-- +goose Down
DROP TABLE location_info;
DROP TABLE locations;