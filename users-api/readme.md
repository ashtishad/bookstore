# Users Microservice API

### Routes

    * GetUserById           GET     /users/:id
    * SearchUsers           GET     /users/search
    * CreateUser            POST    /users
    * UpdateUser            PUT     /users/:id
    * DeleteUser            DELETE  /users/:id
    * GetAll(paginated)     GET     /users

### Design Decisions

    * Hexagonal Architecture.
    * Domain Driven Design.
    * REST API.
    * Postgresql Database.
    * Dependency Injection.
    * Structured Error Library.
    * Concurrent Server with Graceful Shutdown.

### Libraries/Tools/Frameworks

    * Gin.
    * Go Standard Library.
    * Postgresql.
    * Docker.

### Data Exchange Format

    1. JSON (REST)
    2. CSV  (Database Bulk Import)
    3. DTO  (User Level Data Transfer)

### Data Model
    users:
        * id            bigserial,pk
        * name          varchar
        * gender        varchar
        * dateOfBirth   date
        * email         varchar
        * city          varchar
        * created_at    varchar
        * status        smallint
