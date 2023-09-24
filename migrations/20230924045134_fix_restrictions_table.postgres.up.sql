ALTER TABLE restrictions
DROP COLUMN room_id,
DROP COLUMN reservation_id,
DROP COLUMN restriction_id,
DROP COLUMN start_date,
DROP COLUMN end_date,
ADD COLUMN restriction_name VARCHAR(255) NOT NULL;