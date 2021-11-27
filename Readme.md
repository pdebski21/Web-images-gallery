# Web Art Gallery

## User resource
---------------------------------

**GET all resource object:**

> curl localhost:8080/users

**GET single resource object:**

> curl localhost:8080/users/1

**POST single object:**

> curl -XPOST localhost:8080/users -H 'application/json' -d '{"UserID": 100, "Username": "piwo fan", "Password": "lol ale dobre piwo"}'

**UPDATE single object:**

> curl -X PUT localhost:8080/users -H 'application/json' -d '{"UserID": 1, "Username": "piwo fan", "Password": "lol ale dobre piwo"}'

**DELETE single object:**

> curl -X DELETE localhost:8080/users/1

## Image resource
---------------------------------

**GET all resource object:**

> curl localhost:8080/images

**GET single resource object:**

>curl localhost:8080/images/1

**POST single object:**

**UPDATE single object:**

**DELETE single object:**

> curl -X DELETE localhost:8080/images/1

## Comment resource
---------------------------------

**GET all resource object:**

> curl localhost:8080/comments

**GET single resource object:**

> curl localhost:8080/comments/1

**POST single object:**

**UPDATE single object:**

**DELETE single object:**

> curl -X DELETE localhost:8080/comments/1

## Favourite resource
---------------------------------

**GET all resource object:**

> curl localhost:8080/favourites

**GET single resource object:**

> curl localhost:8080/favourites/1

**POST single object:**

**UPDATE single object:**

**DELETE single object:**

> curl -X DELETE localhost:8080/favourites/1