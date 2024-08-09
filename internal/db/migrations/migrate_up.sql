CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    email_address VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    role VARCHAR(50) NOT NULL,
    creation_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "house" (
    id SERIAL PRIMARY KEY,
    address VARCHAR(255) NOT NULL UNIQUE,
    year INT NOT NULL,
    developer VARCHAR(255),
    date_created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_flat_added TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS "flat" (
    number INT,
    house_id INT NOT NULL REFERENCES house(id),
    price INT NOT NULL,
    rooms INT NOT NULL,
    apartment_status VARCHAR(50) NOT NULL,
    date_created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (number, house_id)
);

CREATE OR REPLACE FUNCTION refresh_last_flat_added()
    RETURNS TRIGGER AS $$
BEGIN
UPDATE house
SET last_flat_added = CURRENT_TIMESTAMP
WHERE id = NEW.house_id;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER after_flat_insert
    AFTER INSERT ON flat
    FOR EACH ROW
    EXECUTE FUNCTION refresh_last_flat_added();
