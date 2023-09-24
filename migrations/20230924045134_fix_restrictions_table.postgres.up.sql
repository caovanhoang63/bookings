ALTER TABLE restrictions
    DROP COLUMN room_id,
    DROP COLUMN reservation_id,
    DROP COLUMN restriction_id,
    ADD COLUMN restriction_name VARCHAR(255) NOT NULL;