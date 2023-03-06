# AUTHSERV-ACME

Api to handle user authorization from jwt token.

## Run API service
```bash
go run authserv.go
```

## Example

The authserv api example is described below.

### Request SignIn

`curl -X POST http://localhost:8081/api/auth/signin  -H 'Content-Type: application/json' -d '{"user":"IU456","password":"password456"}'`

### Response
`{"message":"welcome user: IU456","status":200}`

### Request SignOut

`curl -X PATCH http://localhost:8081/api/auth/signout/IU456`

### Response

`{"message": "signout user: IU456","status": 200}`
