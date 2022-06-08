# answer-service
answer-map REST web service

## Setting up a local database and service
- Run `docker compose up`.
- Run `psql -d postgres -U user1 -a -f .\sql\create_schema.sql`.
- Enter password: pass1.
- The service will be available on port 8080.
- answer-map.postman_collection.json can be imported to connect to your local instance on 8080

## Additional questions:

1 How would you support multiple users?
- 
1 How would you support answers with types other than string
1 What are the main bottlenecks of your solution?
1 How would you scale the service to cope with thousands of requests?