-- +goose Up
CREATE TYPE connector_type AS ENUM ('CCS', 'CHAdeMO', 'Type1', 'Type2');

CREATE TABLE connectors (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    external_id VARCHAR(128),
    station_id UUID REFERENCES stations(id) ON DELETE CASCADE NOT NULL,
    type connector_type NOT NULL,
    max_power_kw NUMERIC(8, 2) NOT NULL
);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION check_max_connectors_per_station()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT COUNT(*) FROM connectors WHERE station_id = NEW.station_id) >= 8 THEN
        RAISE EXCEPTION 'Exceeded maximum connectors per station';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER enforce_max_connectors_per_station
BEFORE INSERT OR UPDATE ON connectors
FOR EACH ROW
EXECUTE FUNCTION check_max_connectors_per_station();

-- +goose Down
DROP TABLE IF EXISTS connectors;
DROP FUNCTION IF EXISTS check_max_connectors_per_station();
