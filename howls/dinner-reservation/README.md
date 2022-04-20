## Dinner Reservation
-(just an exercise like a TODO List)

Implementation dinner reservation list, 


> a ``Dinner Reservation`` item represents a client identified by name and group of friends that have a seat at the table



#### Rules
- Client can reserve a table for them and a group of friends
- A table can accommodate different Clients and their group of friends
- At checkIn the number of client group can be bigger than the number that where reserve
  - if the table has space allow group to enter
- when a client leaves the all group leaves


#### Use cases

``/tables``

- ``POST`` add number of tables and seats per table
- ``DELETE`` delete all tables
- ``GET`` list of tables, with number of seats and occupied seats

``/clients``

- ``PUT`` checkIn client to dinner
- ``DELETE`` checkOut client from dinner
- ``GET`` list of all clients having dinner

``/reservation``
- ``GET`` list all reservations
- ``POST`` reserve table for a client


#### API

```
POST /reservation/name
body: 
{
    "table": int,
    "group_size": int
}
response: 
{
    "name": "string"
}
```

#### List reservation's

```
GET /reservation
response: 
{
    "clients": [
        {
            "name": "string",
            "table": int,
            "group_size": int
        }, ...
    ]
}
```

#### CheckIn


```
PUT /clients/name
body:
{
    "group_size": int
}
response:
{
    "name": "string"
}
```

### CheckOut

```
DELETE /clients/name
```

### List Client at the party

```
GET /clients
response: 
{
    "clients": [
        {
            "name": "string",
            "group_size": int,
            "check_in_time": "string"
        }
    ]
}
```


### Implementation: Architecture

This task was implementing in a **Clean Architecture** approach.

It is one of the "evolutionary" Architecture types (like Hexagonal and Onion), This type of code structure focus on the
decouple of the business logic and models from the details of supported frameworks.

- allows easy testing
- Models can change without affecting each others
- easy change/replace frameworks
- easy adapt to change
- loose coupling

https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

### Implementation: Layout

The layout of the project was done in the style of **Screaming Architecture**

This type of layout focus on the Domain entity and Use cases instead of the function and frameworks, mainly supporting
high cohesion.

The source code is structure in folders that have the name of an entity of the business domain. Inside each entity
folder, the code is organized by files that can represent a logic layer of the clean architecture.

- Handler: Represent the client interface, in this task the RestApi endPoints
- Repository: represents the store framework
- Service: the core of the business domain, rules and use cases of each individual domain entity

https://blog.cleancoder.com/uncle-bob/2011/09/30/Screaming-Architecture.html
