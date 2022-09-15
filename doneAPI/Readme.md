# doneAPI Services #

Required Tools: 

- Golang v1.19
- Docker latest  
  

- Packages:  
  
  - docker pull postgres
  - brew install kyleconroy/sqlc/sqlc
  - brew install golang-migrate

Setup:
  
- Open golang project in Goland or similiar IDE
- run the following make commands from project root
  - make postgres : create Docker postgres container
  - make createdb : created the database in postgres container
  - make migrateup : creates required database objects  
  - run the main.go file in project root to start server


Routes:  
&emsp; This API server supports the following routes  
&emsp;&emsp; http://localhost:8000/api/v1  - Server healthCheck return 200 if server is running  
&emsp;&emsp; http://localhost:8000/api/v1/registrations  (GET) returns all registrations  
&emsp;&emsp; http://localhost:8000/api/v1/registrations  (POST) Adds new registration record  
&emsp;&emsp; http://localhost:8000/api/v1/registrations/:id  (GET) returns specific registration  
  

For the simplification of this challenge in a non production example the following has been omitted:  
  
- User Authentication
- Token management
- Use of Environment Variables for database parameters
- Robust error checking for response status codes other than 200, 404, 500

