## Ecommerce Microservice

#### Environment Setup

###### Clone using ssh protocol `git clone git@github.com:ashtishad/ecommerce.git`

To run the application, you have to define the environment variables, default values of the variables are defined inside `start.sh`

- SERVER_ADDRESS    `[IP Address of the machine]` : `localhost`
- SERVER_PORT       `[Port of the machine]` : `8000` `only 5000,5001,8000,8001 are allowed for google auth callback`
- DB_USER           `[Database username]` : `root`
- DB_PASSWD         `[Database password]`: `root`
- DB_ADDR           `[IP address of the database]` : `localhost`
- DB_PORT           `[Port of the database]` : `3306`
- DB_NAME           `[Name of the database]` : `users`

###### MySQL Database Setup
* Make the changes to your `start.sh` file for modifying default db configurations.
* `docker-compose.yml` file. This contains the database migrations scripts. You just need to bring the container up.
* `docker-compose down
  docker volume rm ecommerce_mysqldata` to wipe up a database and remove applied migrations.
  To start the docker container, run the `docker-compose up`.

###### Run the application
* Run the application with `./start.sh` command from project root. or, if you want to run it from IDE, please set
  environment variables by executing command from start.sh on your terminal.
* (optional) Make the Script Executable: You must give the script execute permissions before you can run it. Use the following command:
    `chmod +x start.sh`

#### Tools Used

* Language used: GoLang
* Database Used: MySQL
* Design       : Domain driven design
  * Libraries Used:
    * [Gin](https://github.com/gin-gonic/gin)
    * [Go SQL Driver](https://github.com/go-sql-driver/mysql)
    * [Google OAuth Library](golang.org/x/oauth2/google)

#### Project Structure
```

├── cmd
│   └── app
│       └── app.go                  <-- Define routes, logger setup, wire up handler, start server
│       ├── app_helpers.go          <-- Sanity check, writeResponse
│       ├── app_helpers_test.go     <-- Unit tests of Sanity check, validate input, writeResponse
│       └── db_connection.go        <-- MySQL db connection with dsn
│       └── handlers.go             <-- User handlers for app endpoints
│       └── google_auth_handlers.go <-- Google auth handlers for app endpoints
├── domain
│     └── user.go                   <-- User struct based on database schema
│     ├── user_dto.go               <-- User level data with hiding sensitive fields
│     ├── user_repository.go        <-- Includes core repository interface
│     └── user_repository_db.go     <-- Repository interface implementation with db
│     └── user_sql_queries.go       <-- SQL queries written seperately here
├── service   
│     └── user_service.go           <-- Generate salt,hash pass, covert dto to domain and vice versa
│     └── service_helpers.go.go     <-- Included user input validation
├── migrations                      <-- Database schema migrations scripts
├── docker-compose.yml              <-- Docker setup
├── start.sh                        <-- Builds app with exporting environment variables
├── readme.md                       <-- Self explanetory
├── main.go                         <-- Self explanetory

```


#### Data Flow (Hexagonal architecture)

    Incoming : Client --(JSON)-> REST Handlers --(DTO)-> Service --(Domain Object)-> RepositoryDB

    Outgoing : RepositoryDB --(Domain Object)-> Service --(DTO)-> REST Handlers --(JSON)-> Client


#### Example Requests

###### Create a user

```
curl --location 'localhost:8000/users' \
--header 'Content-Type: application/json' \
--data-raw '{"email":"keanu_reeves@gmail.com","full_name":"Keanu Reeves","password":"secrpsswrd","phone":"1234567890","sign_up_option":"general"}'
```

###### Update a user

```
curl --location --request PUT 'localhost:8000/users/{user_id}' \
--header 'Content-Type: application/json' \
--data-raw '{
	"email": "keanu_reeves@gmail.com",
	"full_name": "John Wick",
    "phone": "1234567890"
	
}
'
```

###### Check Existing User by Email and Password

```

curl --location 'localhost:8000/existing-user' \
--header 'Content-Type: application/json' \
--data-raw '{
"email": "keanu_reeves@gmail.com",
"password": "seepasword"
}'

```

#### Testing Google Auth

```
1. From browser go to, localhost:port/google-auth/login, login with your google account, select/enter your gmail account.
2. After redirecting to callback url, you will see a user created with sign_up_option= google.
3. Please use only 8000,8001,5000,5001 ports as SERVER_PORT in environemnt varibales..


As only backend is implemented for it, so I have skipped usual steps:
Frontend obtains Google OAuth token
Frontend sends OAuth token to backend
Backend verifies the token
```

#### Design Decisions

###### 1. Handle password with salt mechanism

* generates new salt + hashed-password on user creation.
* on update, it also updates the salt value corresponding to user_id.
* used database transactions for multi table update, insert.


###### Hexagonal Architecture

![hexagonal_architecture.png](assets%2Fimages%2Fhexagonal_architecture.png)
