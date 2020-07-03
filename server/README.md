# inventory

### What is this repository for? ###
This is backend for inventory management system

### How do I get set up? ###
* Make production binary
    
    `make production-build`
    
* Copy production binary into release/ and run
    
    `docker build . -t inventory`
    
    `docker run -d --publish 15888:15888 --name docker-inventory inventory `
    
### How do I get run unit test? ###
* Using go test

    `go test ./... `