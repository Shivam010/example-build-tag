# example.com/build/tags
This tests go build tags 

### Run using
Both build `main` functions, one in [main.go](main.go) and other in [directory.go](directory.go)
```sh
go run example.com/build/tags
```
Or to Run only `directory.go` `main` function
```sh
go run -tags dir example.com/build/tags
```
