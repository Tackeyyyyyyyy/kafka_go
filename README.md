# kafka_go


### Steps to run

- Clone this repository
`````
git clone git@github.com:Tackeyyyyyyyy/kafka_go.git
`````

- Install dependencies 
`````
go mod tidy && go mod vendor
`````

- Run kafka broker
`````
docker-compose up
`````

- Make sure Kafka is running
`````
go run main.go
`````

If you run this code, you’ll see an output similar to this:
`````
writes message : 0
received  key: message, value: send message : 0, offset 0, topic: queue-message-log
received  key: Key-A, value: Hello World!, offset 1, topic: queue-message-log
received  key: Key-B, value: One!, offset 2, topic: queue-message-log
received  key: Key-C, value: Two!, offset 3, topic: queue-message-log
writes message : 1
received  key: message, value: send message : 1, offset 4, topic: queue-message-log
received  key: Key-A, value: Hello World!, offset 5, topic: queue-message-log
received  key: Key-B, value: One!, offset 6, topic: queue-message-log
received  key: Key-C, value: Two!, offset 7, topic: queue-message-log
writes message : 2
received  key: message, value: send message : 2, offset 8, topic: queue-message-log
received  key: Key-A, value: Hello World!, offset 9, topic: queue-message-log
received  key: Key-B, value: One!, offset 10, topic: queue-message-log
received  key: Key-C, value: Two!, offset 11, topic: queue-message-log
.
.
.

`````
