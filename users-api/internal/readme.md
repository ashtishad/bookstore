## internal package
**_defines the core domain_**


### Contains

    * domain        -> Entity and Repository
    * Services      -> Core services (e.g. UserService)
    * rest          -> Core http REST handlers (e.g. UserHandler)
    * dto           -> Request and Response DTOs (e.g. UserDTO)
    * middleware    -> Core http middleware (e.g. AuthMiddleware)

### Data Flow

    Incoming : Client --(JSON)-> REST Handlers --(DTO)-> Service --(Domain Object)-> RepositoryDB

    Outgoing : RepositoryDB --(Domain Object)-> Service --(DTO)-> REST Handlers --(JSON)-> Client