ALTER TABLE room_restrictions
ADD CONSTRAINT reservation_id FOREIGN KEY (reservation_id )
REFERENCES reservation (id)