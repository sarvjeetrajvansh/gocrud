
> # GoCRUD – Local Development Guide

## Prerequisites
- Go 1.25+
- Docker + Docker Compose
- curl or Postman

---

## 1. Clone the repository
```bash
git clone https://github.com/sarvjeetrajvansh/gocrud.git
cd gocrud
```
## 2. Start Server
```bash
docker compose up -d
```

## 3. Start Go server
```bash
go run .
```

## 4. Test 
### I. Create User
```bash
 curl -X POST http://localhost:8080/users \ 
-H "Content-Type: application/json" \
-d '{"name":"abc","email":"abc@mail.com", "age":30}'
```
### II. Get All User
```bash
 curl -X GET http://localhost:8080/users 
```
### III. Update User
```bash
 curl -X PUT http://localhost:8080/users/{ID} \
-H "Content-Type: application/json" \
-d '{"name":"Updated","email":"updated@mail.com","age":31}
```
### IV. Delete User
```bash
 curl -X DELETE http://localhost:8080/users/{ID}
```
### V. Get UserByID
```bash
 curl -X GET http://localhost:8080/users/{ID}
```

## 5. Tracing the Request in Jaeger
>Open http://localhost:16686
> Service: gocrud
> Find traces by coppying the trace id from log
```bash
{"time":"2025-12-30T15:22:48.757502+05:30","level":"INFO","msg":"http request","method":"GET","path":"/favicon.ico", \
"status":404,"duration":11792,"request_id":"000011","trace_id":"7bac37132f2a99b7b6cefee7d242cdd0"}
```
## 6. Stopping Everything
```bash
docker compose down
```