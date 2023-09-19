ALTER TABLE room_restrictions
    ADD CONSTRAINT reservation_id FOREIGN KEY (reservation_id)
        REFERENCES reservations(id)
            ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE room_restrictions
    ADD CONSTRAINT  room_id FOREIGN KEY (room_id)
        REFERENCES rooms(id)
            ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE room_restrictions
    ADD CONSTRAINT restrictions_id FOREIGN KEY (restrictions_id )
        REFERENCES restrictions(id)
            ON DELETE CASCADE ON UPDATE CASCADE;