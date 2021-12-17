SELECT *
FROM image_author
WHERE username LIKE 'some_username' AND grade > 5
ORDER BY date_added ASC;