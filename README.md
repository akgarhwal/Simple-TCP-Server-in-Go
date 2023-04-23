# Simple-TCP-Server-in-Go
Simple TCP Server in Go


## Start TCP Server
```sh
go run main.go
```

## Start new terminal and send http request to server
```sh
curl http://localhost:8080/\?name\=Abhinesh 
```

![Simple Example](./res/simple-example.png)


## Server with only one thread
![Single Thread Example](./res/single-thread.png)


## Server with multiple threads
![Multiple Thread Example](./res/multi-thread.png)

## Improvements

* [TODO] Support for Worker Pool since server will not habe enough resources.
* [TODO] Support for time out if request is taking more time than defined threshold