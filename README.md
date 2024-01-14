# Go/Golang server with one endpoint for search of the passed argument in the slice of ints - which should be passed into the stdin for go program to read in and serve

run and test locally: 
```
cat input.txt | go run main.go
```

```
curl localhost:8000/endpoint/1
```

```
curl localhost:8000/endpoint/-1
```
