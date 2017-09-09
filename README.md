# Restful API in Go

This simple api is written in golang.

### Prerequisites
* You've already installed Docker in your local


### the settting to use

1. Install this repository into your local

 `$ git clone https://github.com/Akitsuyoshi/restGo`

2. Build the Docker image

 `$ docker build -t restgo .`

3. Run the image

 `$ docker run -p 80:8080 restgo`

And visit http://localhost/v1/ in your browser, you can see the application run.

Also, when you hit http://localhost/v1/todos/, you see all JSON data.



If you'd like to stop, follow the commands
```
// it shows the container ID, which is running
$ docker ps

// To stop the process with ID
$ docker stop CONTAINER_ID
```

---

The version of api could be controlled in [Router.go](https://github.com/Akitsuyoshi/restGo/blob/master/Router.go)

if necessary, replace v1 to other version.

`r := router.PathPrefix("/v1/").Subrouter()`
