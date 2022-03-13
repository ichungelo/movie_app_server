# REST API WITH ECHO

## TOOLS

- [GO Programming Language](https://go.dev/dl/)
- [ECHO Framework](https://echo.labstack.com/guide/)
- [JWT](https://github.com/dgrijalva/jwt-go)
- [PostgreSQL]()

## REPOSITORY
- [GITHUB REPOSITORY]()

## MIGRATION DATABASE
```
lorem ipsum dolor sit amet
```

## RUN SERVER
```
go run .
```

## API SPEC

### Register Account
Request:
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
Response:
- body:
  ```json
  {
      "success": "boolean",
      "message": "string"
  }
  ```

### Login Account
Request:
- Endpoint: `/api/auth/login`
- Method: POST
- Body:
  ```json
  {
      "username": "string",
      "password": "string"
  }
  ```
Response:
- body:
  ```json
  {
      "success": "boolean",
      "message": "string",
      "token": "JWT"
  }
  ```