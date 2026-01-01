### About
In this repository, my aim is to learn Go Basics and then use Go to create a Utility.
Each folder will contain the concepts that I learned.
I am using this [book](https://www.amazon.com/Mastering-Go-professional-utilities-concurrent-dp-1801079315/dp/1801079315).

### Build/Run Go Program
- `go build fileName`
  - It will produce a statically linked binary file that can be executed in shell.
- `go run fileName`
  - It will simply run the program.
Note: Make sure the `main` method file should contains `package main` in first line. 

### How to execute the Project
- `go get .` to get all deps
- `go run .` to run the project (8080)
- `go test .` to run the tests

### Interaction with Project
- `curl http://localhost:8080/albums` to get a list of albums
- To post an album
```
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "<string>,"title": "<string>","artist": "<string>","price": <number>}'
```