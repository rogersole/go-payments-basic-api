# Payments basic API


TODO:
- date not stored property into PostgreSQL
- daos/payment.go
	- update
- populate db with some test inserts	  	
- tests
- look for TODOs in the code
- docker with postgres to execute everything
- makefile completed
- design document

- Decisions taken:
	- Used relational database to try it out with Go, I could have used a nosql one to store the entire json and would have made it easier,
	but it depends on architectural decisions out of the scope of this exercise.
	- Not using the exact same response format for the list of payments, but a one with pagination information instead

- Things to be improved/done:
  - Simplify db schema, or use a proper one, not invented by me (I may be missing lot of external references and so on)
  - In the paginated response, the previous and next link could be built automatically in the response
  