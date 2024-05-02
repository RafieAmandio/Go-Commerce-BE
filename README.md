# go-starter

Go base project using [gorilla/mux](https://github.com/gorilla/mux) request router and inspired from the Clean code Architecture.


## Folder Structure

here how golang-started manage directory:
```
|-- config
|-- library
|   |-- middleware
|   |-- response
|   |-- logger
|-- server
|-- src
|   |-- Model
|   |-- Dto
|   |-- Handler
|   |-- Service
|   |-- Repository
|-- storage
|   |-- database
```

all global setting will place in library folder.
all module will place in src.

# Config
Config modules for the app

# library
library with their folders are using for the whole apps that will be needed to be a global function that will call in other modules. example:

we have a example modules with their controller, repo, service, in the service we need a jwt for verification the jwt token. the jwt token function will be a callable function on other moduless, hence the jwt must be a global function

# Server
Server will have server config and method releated to server like start and shutdown server and routes 

# src
This boilerplate based on modules. eg: example modules, user modules, product modules etc. and the modules have their structure based on clean architecture. controllers, models, repositories, services. the flow is handler -> service -> repository.

## models
models is a struct to store any objects from database, example you have a message in your tables. then you must create your messages.go in the models folder, it's use to decode the data from database query and store it as a object

## dto
it's similar with models but the differents are:
<ul>
    <li>don't define dto in repository, the repository is the responsibility of the models</li>
    <li>dto is usefull to get your request body, request params, or response your data to your rest api</li>
</ul>
example you have message.go as models and its models called as messages, and messages_response.go as the dto and its dto called as MessageResponseBody. the repositories will return the data as Messages and the services will return it too, but if you want to sent your data as a json response, you should have used the MessageResponseBody to response it.

## Handlers
Handlers usefull to encode raw body, query params, make a response to client. note: make the controller based on the protocols eg: HTTP Handlers, RabbiqMQ controller, Grpc controller.
and it also useful to convert models to dto for responding to the data

## Service
Service contains all of business logic, like creating a messages, creating a users, validation products, validation for products quantity, run you transaction. 
note that if you have relations based on the modules example you have a product modules with order modules, when user make an order then product quantity must be reduced based on order quantity, so you must make a transaction in service

## Repository
Repository contains all of database query

# Storage
Storage store the storage layer, such as database, database stored in storage/db folder. If you use mysql, then just create a file called mysql.go and you must define your mysql modules here (it's a dependency).
If you using redis in your project, create a cached folder then create a file called redis.go, and make your modules again. all examples was attached in this repo


## License
This package is licensed under MIT license. See LICENSE for details.
