⭐️---------- Current Task ----------⭐️

1. Cookies (Left with middleware to authenticate when rendering a new page + redirect logic.)

------ Future things to be implemented if time persists------

1. Circuit breaker - https://levelup.gitconnected.com/circuit-breaker-example-in-golang-e6459c87eaeb

2. Graphql

⭐️-------------- How to run this service? --------------⭐️

After cloning this repo. Go to the makefile and check for the commands.

1. Start up docker desktop

- colima start (to start up dock)

2. type "make start-auto" command
   Depending on the setting in .dev file, the app will either run on host machine or in container.

Now you could try the API with Postman.

**Other steps for reference in case you want more granular operaton** 2. In this project CLI use the followind make commands in order:
❗️ Use this commands if it is your first time running

- make postgres (make postgres container for db connection)
- make createdb (make the database)
- make migrateup (make the database based on the schema in migration file)
- make build (to build this project image)

❗️ Use this commands only after you have created the containers before
💥(USE 2 Seperate CLI for each make command)💥

- make run (to run this project image)
- make start (to run service)

❗️ If container not running

- docker ps -a (Returns all container)
- docker start (Container ID)

⭐️-------------- Testing API Endpoints --------------⭐️

1. Copy the URL and paste in Postman
2. Select POST
3. Select BODY, Raw > JSON in blue dropdown
4. Copy and paste into editor

Create User Endpoint 🐣
Endpoint - http://localhost:8080/createUser (Create User, Unique Username, Email and Mobile)

Sample input
{
"first_name": "John",
"last_name": "Doe",
"user_name": "johndoe",
"user_password": "password123",
"email": "johndoe@example.com",
"mobile": 1234567890
}

User Login Endpoint 🐣 (Store cookie in db and local storage)
Endpoint - http://localhost:8080/userLogin (User Login)

Sample input
{
"user_name": "johndoe",
"user_password": "password123"
}

User Logout Endpoint 🐣 (Removes cookie in db and local storage)
Endpoint - http://localhost:8080/userLogin (User Logout)

Sample input
{
"user_name": "johndoe",
}

Additional info:

Docker container name of the service: login9
Postgres container name: postgres12
Database name: loginMicroservice9
Runs on: http://localhost:8080/
