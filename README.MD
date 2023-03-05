
---------- ❗️Current Task❗️ ----------

1. HTTP endpoints 

2. Git CI/CD pipeline

3. Deply to aws

------ Future things to be implemented ------

1. Authentication - https://www.sohamkamani.com/golang/password-authentication-and-storage/
Constant check for session cookies

2. Circuit breaker - https://levelup.gitconnected.com/circuit-breaker-example-in-golang-e6459c87eaeb

3. Include zap logger

4. Data dog?

5 Graphql

---------- Guidelines for endpoints ---------- 😎
1. Define clear and descriptive routes: Your routes should be clear and descriptive, making it easy for clients to understand how to call the endpoint. Use meaningful URL paths and HTTP methods (e.g. GET, POST, PUT, etc.).

2. Use appropriate HTTP status codes: Return the appropriate HTTP status code for each response. For example, 200 OK for successful requests, 404 Not Found for missing resources, and 500 Internal Server Error for server-side errors.

3. Handle errors: Make sure to handle and return meaningful error messages for unexpected conditions.

4. Validate inputs: Validate inputs from the client to prevent malicious attacks and ensure data integrity.

5. Secure the endpoint: Make sure to secure the endpoint by using secure connections (e.g. HTTPS), implementing authentication and authorization, and sanitizing inputs.

6. Document the endpoint: Provide clear and concise documentation for the endpoint, including the URL, HTTP method, parameters, and expected response format. This will make it easier for clients to understand how to use the endpoint.

7. Test the endpoint: Test the endpoint thoroughly to make sure it works as expected and can handle different types of requests and responses.

8. Monitor and maintain the endpoint: Regularly monitor the endpoint to ensure it is performing well and quickly address any issues that arise. Keep the endpoint up-to-date with the latest security patches and upgrades.


-- How to run this service? -- 
 After cloning this repo

1. Start up docker desktop

2. In this project CLI use the followind make commands in order:
- make postgres (make postgres container for db connection)
- make createdb (make the database)
- make migrateup (make the database based on the schema in migration file)
- make build (to build this project image)
- make run (to run this project image)
- go run cmd/main.go

3. Testing for endpoint on postman
Eg. http://localhost:8080/test1 (Create user)

Copy the URL and paste.
Select POST
Select BODY, Raw > JSON in blue dropdown

Copy and paste into editor
Sample input

{
  "first_name": "John",
  "last_name": "Doe",
  "user_name": "johndoe",
  "user_password": "password123",
  "email": "johndoe@example.com",
  "mobile": 1234567890
}
