-- CREATE DATABASE margertf;

CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    email_address VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    role VARCHAR(50) NOT NULL,
    date_created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "house" (
    id SERIAL PRIMARY KEY,
    address VARCHAR(255) NOT NULL UNIQUE,
    year INT NOT NULL,
    developer VARCHAR(255),
    date_created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    last_flat_added TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "flat" (
    id INT,
    house_id INT NOT NULL REFERENCES house(id),
    price INT NOT NULL,
    rooms INT DEFAULT 1,
    status VARCHAR(50) NOT NULL DEFAULT 'created',
    date_created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id, house_id)
);

CREATE OR REPLACE FUNCTION create_flat_sequence()
    RETURNS TRIGGER AS $$
BEGIN
    EXECUTE format('CREATE SEQUENCE flat_seq_%s START 1', NEW.id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER trg_create_flat_sequence
    AFTER INSERT ON house
    FOR EACH ROW
    EXECUTE FUNCTION create_flat_sequence();


CREATE OR REPLACE FUNCTION assign_flat_id()
RETURNS TRIGGER AS $$
BEGIN
    NEW.id := nextval('flat_seq_' || NEW.house_id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER trg_assign_flat_id
    BEFORE INSERT ON flat
    FOR EACH ROW
    EXECUTE FUNCTION assign_flat_id();


CREATE OR REPLACE FUNCTION refresh_last_flat_added()
    RETURNS TRIGGER AS $$
BEGIN
    UPDATE house
    SET last_flat_added = CURRENT_TIMESTAMP
    WHERE id = NEW.house_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER trg_after_flat_insert
    AfTER INSERT ON flat
    FOR EACH ROW
    EXECUTE FUNCTION refresh_last_flat_added();


CREATE OR REPLACE FUNCTION delete_house_seq()
RETURNS TRIGGER AS $$
BEGIN
    EXECUTE 'DROP SEQUENCE IF EXISTS ' || 'flat_seq_' || OLD.id  || ' CASCADE';
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE TRIGGER trg_after_delete_house
    AFTER DELETE ON house
    FOR EACH ROW
    EXECUTE FUNCTION delete_house_seq();
