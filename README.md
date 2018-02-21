# Payments basic API 

## Requirements

The requirements to fulfill are the following ones:

- Restful API
- API verbs/methods to be implemented    
  - Fetch a payment resource
  - Create, update and delete payment resource
  - List a collection of payment resources
- Persist resource state 

The example *Payment* object structure can be found [here](_docs/single_payment.json).

## Design

### Restful API specification

Given the required methods, the API will look like (note that everything has been grouped under the `v1` path):

- **Authentication:**

Authentication has been simplified for the exercise purposes, it uses a basic plain username/password content sent through Body in JSON format.

  - Endpoint: `POST    /v1/auth`
  - Headers: 
    - `Content-Type: application/json`
  - Body content:
	  ```json
	  {"username": "demo", "password": "pass"}
	  ```
  - Responses:  
    - `200 OK`: Returning the generated JWT token needed for the Authorization header in the other endpoints 
		```json
		{	            
		    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTk0NzQwMTIsImlkIjoiMTAwIiwibmFtZSI6ImRlbW8ifQ.6hZ8NK3sQdlAOeUdb7Nj8wby6BXtBrORd9PatNSUvss"
		}
		```
	- `401 UNAUTHORIZED`: When passing a non expected username/password  
	  Returns `UNAUTHORIZED` format json as content
	- `500 INTERNAL_SERVER_ERROR`: When something goes wrong in the system
	  Returns `INTERNAL_SERVER_ERROR` format json as content

	
- **Healthcheck:**

Used to let external pieces (Load balancers, api gateways, service discoveries, etc...) to know whether the system is healthy or not.
It's a simplified approach always returning `200 OK`, but the healthcheck should also perform basic resources access validation, such as database connectivity and so on to proof that system can handle requests correctly.

  - Endpoint: `GET, HEAD       /healthcheck`
  - Headers: 
    - `Content-Type: application/json`
  - Responses:  
    - `200 OK`: Returning the system healthy as OK and the running service version 
		```text
		OK 1.0
		```
	- `500 INTERNAL_SERVER_ERROR`: When something goes wrong in the system  
      Returns `INTERNAL_SERVER_ERROR` format json as content

- **Fetch a payment resource:**

Used to retrieve and specific Payment information

  - Endpoint: `GET    /v1/payments/{payment_id}`
  - Headers: 
    - `Content-Type: application/json`
    - `Authorization: Bearer {JWT_TOKEN provided by /auth}`
  - Responses:  
    - `200 OK`: Returning the specified payment resource  
      The format can be seen [here](./_docs/single_payment.json)
    - `404 NOT_FOUND`: When asking for a non-existing payment resource  
      Returns `NOT_FOUND` format json as content
    - `401 UNAUTHORIZED`: When passing a non expected username/password  
      Returns `UNAUTHORIZED` format json as content
	- `500 INTERNAL_SERVER_ERROR`: When something goes wrong in the system  
	  Returns `INTERNAL_SERVER_ERROR` format json as content
		
	
- **Create a payment resource:** 

Used to create a new payment resource. It automatically generates all the needed resources in the database.

  - Endpoint: `POST    /v1/payments`
  - Headers: 
    - `Content-Type: application/json`
    - `Authorization: Bearer {JWT_TOKEN provided by /auth}`
  - Body content: The format can be seen [here](./_docs/single_payment.json)
  - Responses:  
    - `200 OK`: Returning the created payment resource
      The format can be seen [here](./_docs/single_payment.json)
    - `400 BAD_REQUEST`: When passing an invalid payment  
      Returns `INVALID_DATA_ERROR` format json as content
    - `401 UNAUTHORIZED`: When passing a non expected username/password   
      Returns `UNAUTHORIZED` format json as content
	- `500 INTERNAL_SERVER_ERROR`: When something goes wrong in the system
	  Returns `INTERNAL_SERVER_ERROR` format json as content

	
- **Update a payment resource:**

Used to update an existing payment resource.

  - Endpoint: `PUT     /v1/payments/{payment_id}`
  - Headers: 
    - `Content-Type: application/json`
    - `Authorization: Bearer {JWT_TOKEN provided by /auth}`
  - Body content: The format can be seen [here](./_docs/single_payment.json)
  - Responses:  
    - `200 OK`: Returning the updated payment resource  
      The format can be seen [here](./_docs/single_payment.json)
    - `400 BAD_REQUEST`: When passing an invalid payment  
      Returns `INVALID_DATA_ERROR` format json as content
    - `404 NOT_FOUND`: When asking for a non-existing payment resource  
      Returns `NOT_FOUND` format json as content
    - `401 UNAUTHORIZED`: When passing a non expected username/password  
      Returns `UNAUTHORIZED` format json as content
	- `500 INTERNAL_SERVER_ERROR`: When something goes wrong in the system.
	  It is also returned when URL `{payment_id}` is different from the body content one  
	  Returns `INTERNAL_SERVER_ERROR` format json as content
 
- **Delete a payment resource:**  

Used to remove a payment resource and all its dependencies.

  - Endpoint: `DELETE    /v1/payments/{payment_id}`
  - Headers: 
    - `Content-Type: application/json`
    - `Authorization: Bearer {JWT_TOKEN provided by /auth}`
  - Responses:  
    - `200 OK`: Returning the deleted payment resource
      The format can be seen [here](./_docs/single_payment.json)
    - `401 UNAUTHORIZED`: When passing a non expected username/password   
      Returns `UNAUTHORIZED` format json as content
    - `404 NOT_FOUND`: When asking for a non-existing payment resource
      Returns `NOT_FOUND` format json as content
	- `500 INTERNAL_SERVER_ERROR`: When something goes wrong in the system
	  Returns `INTERNAL_SERVER_ERROR` format json as content

	
- **List a collection of payment resources:**

Used to get the list of payments. It has been implemented with `pagination`, therefore, specific query params are  
available.

  - Endpoint: `GET    /v1/payments[?page={page_number}&per_page={elements_per_page}]`
  - Headers: 
    - `Content-Type: application/json`
    - `Authorization: Bearer {JWT_TOKEN provided by /auth}`
  - Responses:  
    - `200 OK`: Returning the list of available payments, with information regarding page and total elements  
      The format can be seen [here](./_docs/payments_list.json)
    - `401 UNAUTHORIZED`: When passing a non expected username/password  
      Returns `UNAUTHORIZED` format json as content
	- `500 INTERNAL_SERVER_ERROR`: When something goes wrong in the system  
	  Returns `INTERNAL_SERVER_ERROR` format json as content


### Errors responses

The system responses are unified and defined as an external `.yaml` file.
It can be found [here](./config/errors.yaml), and the responses examples are the following ones:  
And have the following format:

- `INTERNAL_SERVER_ERROR`
	```json
	{
	    "error_code": "INTERNAL_SERVER_ERROR",
	    "message": "we have encountered an internal server error",
	    "developer_message": "internal server error: invalid character '}' looking for beginning of value"
	}
	```

- `INVALID_DATA`
	```json
	{
	    "error_code": "INVALID_DATA",
	    "message": "there is some problem with the data you submitted. See \"details\" for more information",
	    "details": [
	        {
	            "field": "attributes",
	            "error": "cannot be blank"
	        }
	    ]
	}
	```

- `UNAUTHORIZED`
	```json
	{
	    "error_code": "UNAUTHORIZED",
	    "message": "authentication failed",
	    "developer_message": "authentication failed: Unauthorized"
	}
	```

- `NOT_FOUND`
	```json
	{
	    "error_code": "NOT_FOUND",
	    "message": "the requested resource was not found"
	}
	```

### Database tables design

The storage system chosen is `postgres`. Due to its popularity, open source approach and good performance, the election  
was clear and easy. Another approach would have been to go for a NoSQL database to simply store the JSON resources as  
documents and retrieve them all, but given that the Payment resource is a composition of different parts (which seem  
to be managed by independent resources as well), the relation approach seemed more scalable to solve this exercise.

To design the database, the provided [example resource](./_docs/single_payment.json) has been analysed and the design
and the resulting schema is the following one:

![Database design](./_docs/db_design.png)

Note:
 - a `charges information` can have multiple `sender charge` resources.
 - a `payment attribute` have exactly 3 `party` resources (Beneficiary, Debtor and Sponsor) 

## Implementation

### Programming language chosen and why

The service has been developed in `Go`, version `1.10`.
Even I feel comfortable about using Java, and I've been using a lot during my last years, I found interesting to solve  
this exercise in Go to deep learn things related to language.

### Frameworks/libraries used

- `go-ozzo/ozzo-dbx`: for abstracting database communications. It supports several RDBMS to connect with. 
  All database interactions are encapsulated in transactions. The commit is performed automatically, and the rollback  
  only on an error triggered situations. 
- `go-ozzo/ozzo-routing`: for handling requests routing and handlers.  
   It also allows to enable CORS easily (done in the exercise)
- `go-ozzo/ozzo/validation`: for handling requests validations
- `satori/go.uuid`: for handling UUIDs management
- `sirupsen/logrus`: for logging. Structured logging is used. It can simplify searches in a documental storage system
- `spf13/viper`: for configuration management/reading
- `stretchr/testify`: for testing
- `dgrijalva/jwt-go`: for implementing JWT on authentication

### How to run tests locally

```bash
make test
```

### How to run the service locally
 
```bash
make run
```

### Code structure

The code has a simple and completely isolated structure, based on MVC, using DAOs (Data Access Object) to communicate  
with the storing system, DTOs (Data Transfer Object) to communicate with clients and Services to abstract the service   
definition.

Given this structure, adding another resource would only imply adding things into:

- `apis/resource_x`: to process body content and convert into DTOs data and pass them to service
- `daos/resource_x`: to process process DTOs, transform them into DAOs and interact with storage system
- `dtos/resource_x`: to define how the expected and returned json objects are mapped into objects
- `services/resource_x`: to validate DTOs input and pass the to DAO

### Decisions taken

- Not using the exact same response format for the list of payments, but a one with pagination information instead.
- Simplified date field to string since the POST is considered to be done by storage system once the insert is  
  performed. But it depends on architectural designs not related to this exercise.
- On `PUT` method, if resource is not found, it is not automatically created, as would be expected with `PUT` operation,   
  it could be done, but it returns a not found and does nothing with the stored data.

## Further work

- Build a proper db schema knowing all the involved data scope
- In the paginated responses, build the `previous` and `next` links in the headers or in the body content
- Model validation is only applied to check the `payment attributes` field is not null. It could be extended to all the  
  model validation or change it for a json schema validation to force an specific data format when being sent to the  
  server.
- Change `account type` field from `int 0` to something that does not map as a Null or default value in the programming   
  language.
- Extend healthcheck to perform storage checks or more complex functions
