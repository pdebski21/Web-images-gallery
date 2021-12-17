CREATE VIEW image_author AS
  SELECT images.id,
    images.name,
    users.username,
    images.date_added,
    images.value,
    images.extension,
    avg(grades.grade) AS grade
   FROM images
     JOIN users ON images.author_id = users.id
     LEFT JOIN grades ON images.id = grades.image_id
  GROUP BY images.id, images.name, users.username, images.date_added, images.value, images.extension;