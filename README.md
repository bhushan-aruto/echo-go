Echo-Go Simple Usage 🚀

This is a basic Echo framework project in Go that demonstrates handling:

Simple GET requests

Path parameters

Query parameters

📌 API Endpoints & Examples

1️⃣ Get a Simple Message

Endpoint: GET /hello

Response:

{"message": "Hi, you are using the Echo framework"}

Test in Browser or Postman:http://localhost:8080/hello

2️⃣ Get a Message with Query Parameter

Endpoint: GET /queryhello?name=bhushan

Response:

{"message": "Hi, this is Echo framework, and I am bhushan"}

Test in Browser or Postman:http://localhost:8080/queryhello?name=bhushan

3️⃣ Get a Message with Path Parameter

Endpoint: GET /paramhello/{name}

Example URL: GET /paramhello/bhushan

Response:

{"message": "Hi, my name is bhushan"}

Test in Browser or Postman:http://localhost:8080/paramhello/bhushan

