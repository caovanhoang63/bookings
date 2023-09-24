ALTER TABLE restrictions
ADD COLUMN room_id integer,
ADD COLUMN reservation_id integer,
ADD COLUMN restriction_id integer,
ADD COLUMN start_date date,
ADD COLUMN end_date date,
DROP COLUMN restriction_name VARCHAR(255) NOT NULL;