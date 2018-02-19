# Payments basic API


TODO:
- tests
- docker with postgres to execute everything
- makefile completed
- design document
- look for TODOs in the code

- Decisions taken:
	- Used relational database to try it out with Go, I could have used a nosql one to store the entire json and would have made it easier,
	but it depends on architectural decisions out of the scope of this exercise.
	- Not using the exact same response format for the list of payments, but a one with pagination information instead
	- Simplified date field to string since the POST is considered to be done with a processing_date input already populated (otherwise changing to Time is easy and ussing current_timestamp when inserting the standard approach used)

- Things to be improved/done:
  - Simplify db schema, or use a proper one, not invented by me (I may be missing lot of external references and so on)
  - In the paginated response, the previous and next link could be built automatically in the response
  - Model validation on POST, knowing all the required fields, there is a Validate method which has not been implemented (dtos/payment.go::Validate())
  
- Implemented
  - healthcheck in GET, HEAD
  - logger
  - cors enabled
  - all db interactions are transactions
  - auth returning a jwt used as Authorization Bearer token in all the requests
  - everything under the v1 common path group
  - structured log (good for storing easy searcheable data)