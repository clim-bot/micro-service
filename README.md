# micro-service
This project is for creating a basic micro service that let's the user register, login and logout. After the user has login. The user can create a new client to be able to attach to an order. The user can update and delete a client. And the order service lets the user create a product and how many units it will be.

### Installation
You just need to CD on the micro service and type `go mod tidy`.

### Running the Service
While on the root folder of the micro service project type `make build` it will build a docker image for the
three services. If you want to start the services type `make up`. Then if you want to stop the docker instance, you need to ype `make down`.

### Curl Commands
#### User Service
```
Register a User:
curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"name":"John Doe", "email":"john@example.com", "password":"password123"}'

Login a User:
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"email":"john@example.com", "password":"password123"}'

Logout a User:
curl -X POST http://localhost:8080/logout -H "Authorization: Bearer <token>"

Update Password:
curl -X POST http://localhost:8080/update-password -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"old_password":"password123", "new_password":"newpassword456"}'

Reset Password:
curl -X POST http://localhost:8080/reset-password -H "Content-Type: application/json" -d '{"email":"john@example.com"}'


Get a User Profile:
curl -X GET http://localhost:8080/profile -H "Authorization: Bearer <token>"
```

#### Order Service
```
Create an Order:
curl -X POST http://localhost:8081/orders -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"client_id":1, "product":"Laptop", "amount":1}'

Get All Orders:
curl -X GET http://localhost:8081/orders -H "Authorization: Bearer <token>"

Get Single Order:
curl -X GET http://localhost:8081/orders/1 -H "Authorization: Bearer <token>"

Delete an Order:
curl -X DELETE http://localhost:8081/orders/1 -H "Authorization: Bearer <token>"
```

#### Client Service
```
Create a Client:
curl -X POST http://localhost:8082/clients -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"name":"Jane Doe", "email":"jane@example.com", "phone":"1234567890", "zipcode":"12345"}'


Update a Client:
curl -X PUT http://localhost:8082/clients/1 -H "Authorization: Bearer <token>" -H "Content-Type: application/json" -d '{"name":"Jane Doe", "email":"jane@example.com", "phone":"0987654321", "zipcode":"54321"}'


Delete a Client:
curl -X DELETE http://localhost:8082/clients/1 -H "Authorization: Bearer <token>"

Get a Single Client:
curl -X GET http://localhost:8082/clients/1 -H "Authorization: Bearer <token>"

Get All Clients:
curl -X GET http://localhost:8082/clients -H "Authorization: Bearer <token>"

```