## Internal Package
**_defines the core domain_**


### What's Inside?

    * domain        -> Entity and Repository
    * Services      -> Core services (e.g. UserService)
    * rest          -> Core http REST handlers (e.g. UserHandler)
    * dto           -> Request and Response DTOs (e.g. UserDTO)
    * middleware    -> Core http middleware (e.g. AuthMiddleware)

### Data flow?

**Incoming** : Client --(JSON)-> REST Handlers --(DTO)-> Service --(Struct)-> RepositoryDB

**Outgoing** : RepositoryDB --(Struct)-> Service --(DTO)-> REST Handlers --(JSON)-> Client