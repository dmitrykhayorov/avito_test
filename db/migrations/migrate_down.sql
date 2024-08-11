DO $$
    DECLARE
        ids RECORD;
    BEGIN
        FOR ids IN
            SELECT id as h_id
            FROM house
        LOOP
                EXECUTE 'DROP SEQUENCE IF EXISTS ' || 'flat_seq_' || ids.h_id  || ' CASCADE';
        END LOOP;
    END $$;

DROP TRIGGER IF EXISTS trg_assign_flat_id ON flat;
DROP TRIGGER trg_create_flat_sequence ON house;


DROP FUNCTION assign_flat_id();
DROP FUNCTION refresh_last_flat_added()

DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "flat";
DROP TABLE IF EXISTS "house";






-- DROP DATABASE margertf;