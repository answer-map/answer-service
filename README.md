# answer-service

answer-map REST web service

## Setting up a local database and service

- Run `docker compose up`.
- Run `psql -d postgres -U user1 -a -f .\sql\create_schema.sql`.
- Enter password: pass1.
- The service will be available on port 8080.
- answer-map.postman_collection.json can be imported to connect to your local instance on 8080

## Additional questions:

1. How would you support multiple users?
  - there is an answer_user table with a primary key to identify each answer_user (user_id)
  - this user_id can be referenced in the answer_map table and with a unique constraint to make each answer_map unique to the user_id and answer_key
1. How would you support answers with types other than string?
  - Suppose that the answer value type is not known, then we can avoid decoding it from json by defining it as [AnswerValue *json.RawMessage](https://pkg.go.dev/encoding/json#RawMessage) in the AnswerMap struct. That raw message can be marshalled into a byte array before being stored as a byte array in the database.
1. What are the main bottlenecks of your solution?
  - one [replica](https://docs.docker.com/compose/compose-file/deploy/#replicas) of the docker image that runs the service
  - uses a sql docker image with limited resources (no volumes, small storage, not load-balanced)
1. How would you scale the service to cope with thousands of requests?
  - deploy more replicas of the docker image that runs the service and use a load-balancer
  - use a [datasource](https://cloud.google.com/sql/docs/) that can autoscale resources
