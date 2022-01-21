# Take home Test

## Requirement
- Go have to be installed
- internet connection

## To run the solution on MAC
```sh
    ./homeTest
```
## To run the solution on any other platform
```go
go run main.go
```

## About my code

I created Struct to describe the data and a method
that will compute the aggregate of successful rate.

I created 100 goroutines workers to manage all the concurrent work being performed.
The worker function takes two arguments a receiver channel and a sender channel.
then I created a buffered channel with a cap of 100, which means i can send an item without 
waiting for a receiver to read the item. This is a performance increase, as it will allow all the goroutines to start immediately.

The additional channel tracks the workers function and to prevent a race condition by ensuring the completion (1000) of all gathered results.

