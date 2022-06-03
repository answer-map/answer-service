# answer-service
answer-map REST web service

## Setting up a local database and service
- Run `docker compose up`.
- Run `psql -d postgres -U user1 -a -f .\sql\create_schema.sql`.
- Enter password: pass1.
- The service will be available on port 8080.
