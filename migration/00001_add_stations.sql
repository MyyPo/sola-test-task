-- +goose Up
CREATE TABLE stations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    external_uid VARCHAR(128) UNIQUE,
    public BOOLEAN NOT NULL DEFAULT FALSE,
    title VARCHAR(255),
    description TEXT,
    address VARCHAR(1024),
    latitude NUMERIC,
    longitude NUMERIC
);

ALTER TABLE stations
ADD CONSTRAINT check_public_station_requirements
CHECK (
    (public AND title IS NOT NULL AND description IS NOT NULL AND address IS NOT NULL AND latitude IS NOT NULL AND longitude IS NOT NULL)
    OR NOT public
);

-- +goose Down
DROP TABLE IF EXISTS stations;
