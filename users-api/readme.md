# Users Microservice API

#### Routes

    * GetUserById    GET    /users/:id
    * SearchUsers    GET    /users/search
    * CreateUser     POST   /users
    * UpdateUser     PUT    /users/:id
    * DeleteUser     DELETE /users/:id

#### Design Decisions

    * Hexagonal Architecture.
    * REST Api.
    * Dependency Injection.
    * Structured Error Library.
    * Concurrent Server with Graceful Shutdown.

#### Libraries/Tools/Frameworks

    * Gin
    * Go Standard Library