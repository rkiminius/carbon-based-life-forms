## Carbon-Based-Life-Forms
Solution of task: https://github.com/heficed/Carbon-Based-Life-Forms

Application use rabbitMQ and mongoDB service in this case you need to set up docker or your local environment

In Docker approach run:
```
docker-compose up -d
```

Client, Manager and Factory is separated applications that use AMQP and HTTP protocols. 

### Client
To start client run: 
```
go run client_main.go
```

HTTP request:
```
Find all clients minerals
GET: localhost:1323/client/minerals

Find client minerals using uuid
GET: localhost:1323/client/minerals/{uuid}

Ask to perform action on mineral
POST: localhost:1323/client/order
{
    "mineralId": "5bd0e70b9db9ea0011519bd5",
    "action": "FRACTURE"
}
```
### Manager
To start manager run:
```
go run manager_main.go
```

HTTP request:
```
Find all mineral types
GET: localhost:1324/manager/mineralType/all

Add new mineral type
POST: localhost:1324/manager/mineralType/new

Remove mineral type by id
DELETE: localhost:1324/manager/mineralType/{mineral_type_id}

Find task by id
GET: localhost:1324/manager/task/{task_id}

Get all tasks
GET: localhost:1324/manager/task
```

### Factory

To start factory run:
```
go run factory_main.go
```
