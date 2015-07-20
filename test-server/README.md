# Test Server

Learn from https://github.com/ant0ine/go-json-rest#countries

Demonstrate simple POST GET and DELETE operations

curl demo:
```
curl -i -H 'Content-Type: application/json' \
    -d '{"Distribution":"FR","Arch":"France"}' http://127.0.0.1:8080/os
curl -i -H 'Content-Type: application/json' \
    -d '{"Distribution":"US","Arch":"United States"}' http://127.0.0.1:8080/os
curl -i http://127.0.0.1:8080/os/FR
curl -i http://127.0.0.1:8080/os/US
curl -i http://127.0.0.1:8080/os
curl -i -X DELETE http://127.0.0.1:8080/os/FR
curl -i http://127.0.0.1:8080/os
curl -i -X DELETE http://127.0.0.1:8080/os/US
curl -i http://127.0.0.1:8080/os
```
