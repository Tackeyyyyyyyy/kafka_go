package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "message-log"
	brokerAddress = "localhost:9092"
	partition     = 0
)

func main() {
	ctx := context.Background()
	go produce(ctx)
	consume(ctx)
}

func produce(ctx context.Context) {
	// initialize a counter
	i := 0

	log.New(os.Stdout, "kafka writer: ", 0)

	conn, err := kafka.DialLeader(ctx, "tcp", brokerAddress, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	for {
		m := fmt.Sprintf("send message : %d", i)
		message := kafka.Message{
			Key:   []byte("message"),
			Value: []byte(m),
		}
		_, err := conn.WriteMessages(message,
			kafka.Message{
				Key:   []byte("Key-A"),
				Value: []byte("Hello World!"),
			},
			kafka.Message{
				Key:   []byte("Key-B"),
				Value: []byte("One!"),
			},
			kafka.Message{
				Key:   []byte("Key-C"),
				Value: []byte("Two!"),
			})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		fmt.Println("writes message :", i)
		i++

		time.Sleep(time.Second * 3)
	}
}

func consume(ctx context.Context) {

	l := log.New(os.Stdout, "kafka reader: ", 0)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  l,
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		m := fmt.Sprintf("received  key: %s, value: %s, topic: %s", string(msg.Key), string(msg.Value), msg.Topic)
		fmt.Println(m)
	}

}
