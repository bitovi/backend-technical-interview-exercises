# Bitovi Backend Service Exercise 

This exercise can be completed in NodeJS (w/ Express), Python (w/ FastAPI), or Java (w/ Spring Boot), as the interviewee
prefers.

## Scenario

The Sales Department has requested an API endpoint to fetch a list of sales people and their monthly aggregate sales, 
while optionally ignoring any months with amounts below a parameter value they pass in. Your task is to implement that
parameter.

- Step 1: add a required parameter to the `GET /sales_person` endpoint with appropriate validation
- Step 2: use that parameter value to exclude some data

More details are provided in comments in the source files:
- Node, see `./node/index.js`

- Java, see `./java/src/main/java/com/bitovi/service_exercise/SalesPersonController.java`

- Python, see `./python/main.py`

- Go, see `./go/src/main.go`

## Command line operations

The interviewer will start the server.
The Node & Python servers will automatically reload after changes have been performed.

Make a test request against the server with a command like this:
```
# note the port may vary between the different languages 
curl 'localhost:3002/sales_person'
```

## Interviewer setup instructions

Node:
```
npm start
```

Java:
```
// you will need to have a JDK installed, or run the spring-boot:run Maven target from an IDE

mvnw spring-boot:run
```

Python:
```
// you will need to set up a venv and install requirement.txt

fastapi dev main.py
```

Go:
```
// you will need to have go installed and install go modules

go run main.go
```