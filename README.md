### go-workout-01 ###
Go's workout 01, based on [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

#### Go commands ####
To run test
```console
go test
```
To run test with coverage
```console
go test -cover
```
To run test (verbose)
```console
go test -test.v
```
To run a (particular) test
```console
go test -run <TestName>
go test -run TestArea
```
To run test with benchmark
```console
go test -bench=<LOCATION>
go test -bench=.
```

To run test with the [race detector](https://blog.golang.org/race-detector)
```console
go test -race
```