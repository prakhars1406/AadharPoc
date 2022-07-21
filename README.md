# Aadhar POC

## About the project

The POC exposes two endpoints
- POST /aadhar  to insert user aadhar details
- GET /aadhar/:id to fetch user aadhar details in XML format

### To run the server
```
go run app/main.go
```

### Database

Mongo db is used. Check the config.go for database name and collection name.
Create index on id field:
```
db.aadharDetails.createIndex({ "id": 1 })
```
