# Fiesta.Pics Frontend - Social media platform for parties!

## Frontend - https://github.com/alex-305/fiestafrontend

## About

The Fiesta.Pics RESTful API is a collection of endpoints that allow you to retrieve information from the database, insert records, and update records. Additionally, I implemented my own user authentication system which leverages the power of JWT tokens along with bcrypt's hashing function to store user password in a hashed form on the database. The application is split into 3 main packages, **handlers**, **auth**, and **db**.

### Handlers
The handlers package is where my API endpoints' routes are defined. There are 20 endpoints with 9 of those being POST requests, 8 being GET requests, and 3 being DELETE requests. These are all stateless endpoints, all of which, communicate with the db package to insert, select, or delete records. Additionally, it communicates with the auth package to validate user authentication.

### Auth
The auth package is where everything related to user authentication and user authorization is handled. This package utilizes JWT tokens for token management and bcrypt to hash passwords. It validates JWT tokens by parsing it with the secret key then checking to see that the claims are accurate by comparing it against database records.

### DB
The db package is where the database is connected to using a DSN. It is responsible for making all the queries to the database including inserts, selection, and deletion.

### Additional Packages

#### Models
The models package is where my structs are defined. These mainly consist of types that correspond to my database schema.

#### Helpers
The helpers package is where common tasks used throughout my application are abstracted away into a simple function.