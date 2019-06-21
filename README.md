# Momotaro

a CLI client for [Kibidango](https://github.com/tomocy/kibidango) (a container from scratch)

### Environment
```
docker run -it -v $GOPATH:/go -w /go/src/github.com/tomocy/momotaro --rm --privileged golang:alpine
```

### Usage
- create a kibidango
    ```
    go run main.go create
    ```