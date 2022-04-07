# REST API WITH ECHO

## TOOLS

- [GO Programming Language](https://go.dev/dl/)
- [ECHO Framework](https://echo.labstack.com/guide/)
- [JWT](https://github.com/dgrijalva/jwt-go)
- [Mysql](https://ubuntu.com/server/docs/databases-mysql)
- [Migrate](https://github.com/golang-migrate/migrate)
---
## REPOSITORY
- [GITHUB REPOSITORY](https://github.com/ichungelo/movie_app_server.git)
---

## MIGRATION DATABASE

### MIGRATE UP
```
migrate -path migrations/ -database 'mysql://{user}:{password}@tcp({host}:{port})/dbname?query' -verbose up
```
### MIGRATE DOWN
```
migrate -path migrations/ -database 'mysql://{user}:{password}@tcp({host}:{port})/dbname?query' -verbose down
```
### SEED DATABASE
```
mysql -u user -p movie_app < migrations/seeds/*.sql
```
### STEP BY STEP
- create db name movie_app
- then run command below :
``` shell
$ mysql -u {dbUser} -p movie_app < migrations/seeds/*.sql
$ go run .
```
---

## RUN SERVER
```
go run .
```
## RUN SWAGGER
[Click Here](http://localhost:5000/swaggerui)
---
## API SPEC

### Register Account
#### Request:
- Endpoint: `/api/auth/register`
- Method: POST
- Body:
  ```json
  {
      "username": "string",
      "email": "string",
      "first_name": "string",
      "last_name": "string",
      "password": "string",
      "confirm_password": "string"
  }
  ```
#### Response:
- body:
  ```json
  {
      "success": "boolean",
      "message": "string"
  }
  ```

### Login Account
#### Request:
- Endpoint: `/api/auth/login`
- Method: POST
- Body:
  ```json
  {
      "username": "string",
      "password": "string"
  }
  ```
#### Response:
- body:
  ```json
  {
      "success": "boolean",
      "message": "string",
      "token": "JWT"
  }
  ```


### Get All Movies
#### Request:
- Endpoint: `/api/movies`
- Method: GET

#### Response:
- body:
  ```json
  {
      "success": "boolean",
      "message": [
        {
          "movie_id": "int"
          "title": "string"
          "release_year": "year"
          "production": "endpoint"
          "overview": "text"
        },
        {
          "movie_id": "int"
          "title": "string"
          "release_year": "year"
          "production": "endpoint"
          "overview": "text"
        },...
      ]
  }
  ```

### Post Movie By Id
#### Request:
- Endpoint: `/api/movies/{movieId}`
- Method: GET

#### Response:
- body:
  ```json
  {
    "success": "boolean",
    "message": {
        "movie_id": "int",
        "title": "string",
        "year": "year",
        "poster": "endpoint",
        "overview": "string",
        "reviews": [
          {
            "review_id": "int",
            "username": "string",
            "review": "string",
          },
          {
            "review_id": "int",
            "username": "string",
            "review": "string",
          },...
        ]
      }
  }
  ```

### Post Review
#### Request:
- Endpoint: `/api/movies/{movieId}/review`
- Method: POST
- Body:
  ```json
  {
      "review": "string",
  }
  ```
#### Response:
- body:
  ```json
  {
      "success": "boolean",
      "message": "string"
  }
  ```

### Update Review
#### Request:
- Endpoint: `/api/movies/{movieId}/review/{reviewId}`
- Method: PUT
- Body:
  ```json
  {
      "review": "string",
  }
  ```
#### Response:
- body:
  ```json
  {
      "success": "boolean",
      "message": "string"
  }
  ```

### Delete Review
#### Request:
- Endpoint: `/api/movies/{movieId}/review/{reviewId}`
- Method: DELETE

#### Response:
- body:
  ```json
  {
      "success": "boolean",
      "message": "string"
  }
  ```