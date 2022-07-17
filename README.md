# **Fampay**

Description

#### Pre-Requisite
- Docker
- Docker Compose
- Golang Version 1.17
- add developer keys in docker-compose.yaml file on line no 38 as below in web service environments comma seperated
```
DEVLOPER_KEYS=AIzaSyA70g-xfN9xzP9SpDibrEWZT78XlNl8eSM,AIzaSyDufGVw6ep1JKBO1vFrvCqkQJ8kr0CmdlI
```

##### features of the service
- Service calls youtube search list api in background continuously for `cricket` term and get vedio details and save in elasticsearch db
- there is a paginated vedios endpoint in service which takes searchterm, offset and pagesize parameters and provide paginated result

#### Run application
```
make run
```

#### Run testcases
```
make test
```

[Postman Documentation](https://documenter.getpostman.com/view/3571564/UzQvtQaQ)


