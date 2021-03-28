CREATE TABLE IF NOT EXISTS car
(
    id    SERIAL PRIMARY KEY,
    brand VARCHAR NOT NULL,
    model VARCHAR NOT NULL,
    price INTEGER NOT NULL,
    status VARCHAR,
    mileage INTEGER
);