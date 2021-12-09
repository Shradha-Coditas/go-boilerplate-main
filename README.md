***Build***
```
go install .
```
***Generate Swagger***
```
swag init -g util\server\server.go
```
***Run***
```
go run main.go
```

***Build Docker image***
```
docker build -t golang/boilerplate .
```

***tag it***
```
docker tag golang/boilerplater:latest golang/boilerplate:$version
```

***Push Image to Docker Hub**
```
docker login --username name --password **********
docker push golang/boilerplate:latest
docker push golang/boilerplate:$version
```