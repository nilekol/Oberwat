
#Starting redis server
redis-server &

# Running 3 instances of the Go server on different ports
go run main.go &
PORT=8081 go run main.go &
PORT=8082 go run main.go &

wait
