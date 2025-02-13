## Golang Middleware Basic Authentication and JWT
**Feature:**
- Level Logging
- Password encryption with bcrypt
- Middleware basic authentication
- Middleware JWT (JSON Web Token)

## Endpoint:
**Register user**
- Method: [POST]
- URL: {base_url}:3000/register

Body:
```sh
{
    "username": "bintang",
    "password": "Bintang1#"
}
```
Response:
```sh
{
    "code": 200,
    "data": {
        "Username": "bintang",
        "Password": "{HASH}"
    },
    "message": "Created successfully"
}
```

**Login user & get token**
- Method: [POST]
- URL: {base_url}:3000/login

Body:
```sh
{
    "username": "bintang",
    "password": "Bintang1#"
}
```
Response:
```sh
{
    "code": 200,
    "message": "Login successfully",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzk0MzU2MzAsInVzZXJuYW1lIjoiYmludGFuZyJ9.r-qmomQfuhvacUlSAvXYk3b0tk-D4qH8q_fpAX7N3To",
    "token_exp": "2025-02-13T16:32:50.467115+07:00"
}
```

**Protected Basic Auth**
- Method: [GET]
- URL: {base_url}:3000/basic/info

Authentication:
- Type: Basic Auth

Body:
```sh
no body
```
Response:
```sh
{
    "code": 200,
    "message": "Basic Auth Successfully"
}
```
Response Error (No Auth):
```sh
{
    "code": 401,
    "message": "Unauthorized"
}
```
Response Error (Invalid):
```sh
{
    "code": 401,
    "message": "Invalid username or password"
}
```

**Protected JWT**
- Method: [GET]
- URL: {base_url}:3000/jwt/notif

Authentication:
- Type: Bearer Token

Body:
```sh
no body
```
Response:
```sh
{
    "code": 200,
    "message": "JWT Successfully"
}
```
Response Error (No Auth):
```sh
{
    "code": 401,
    "message": "Missing authorization header"
}
```
Response Error (Invalid Type):
```sh
{
    "code": 401,
    "message": "Invalid authorization header format"
}
```
Response Error (Invalid Token or Expired):
```sh
{
    "code": 401,
    "message": "Invalid Token"
}
```
