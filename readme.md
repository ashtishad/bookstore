## Bookstore Microservice(V2)

## Design Decisions

 1. Database per service pattern
    * Separate database for each service, one service won't communicate directly to another service's database.
    * Why?
        * Separation of concerns(Each service to run independently).
        * Database schema/structure of other service that might change unexpectedly won't affect another.
        * There won't be a single point of failure would increase Site Reliability.
        * Some services might function more efficiently with different types of DB's (sql vs nosql).
        * Easy to scale,test,manage,maintain and audit.


2. How to exchange data between services?
    * Asynchronous Data Communication(Event Driven).
    * Will use Event Bus to exchange data(eg: Kafka/RabbitMQ/NATS).
    * Why Async Communication?
      * Zero dependency on other services.
      * No need to wait for other services to be ready.
      * Addition of new services is easy and service operations will be extremely fast.
    * Downside? - Data duplication.
    
    
## Libraries/Tools/Frameworks

    * Gin.
    * Docker
    * Kubernetes.
    * PostgreSQL and MongoDB

## Data Exchange Format

    1. JSON (REST)
    2. CSV  (Database Bulk Import)
    3. DTO  (User Level Data Transfer)

## Data Flow(Hexagonal Architecture)

    Incoming : Client --(JSON)-> REST Handlers --(DTO)-> Service --(Domain Object)-> RepositoryDB

    Outgoing : RepositoryDB --(Domain Object)-> Service --(DTO)-> REST Handlers --(JSON)-> Client


[Project status : Ongoing]
