
-- Watch out for golang-added comments - it screws with the sequence.
INSERT INTO comments (image_id, user_id, text, date_added)
VALUES (1, 1, 'dupa', current_timestamp);