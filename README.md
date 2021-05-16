## Carbon-Based-Life-Forms
Task: https://github.com/heficed/Carbon-Based-Life-Forms

Application use rabbitMQ and mongoDB service in this case you need to set up docker or your local environment

In Docker approach run:
- docker-compose up -d

Client, Manager and Factory is separated applications that use amqp and http protocols. 

To start client run: 
- go run client_main.go

To start manager run:
- go run manager_main.go

To start factory run:
- go run factory_main.go