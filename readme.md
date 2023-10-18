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