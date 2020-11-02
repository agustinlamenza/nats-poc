package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	for {
		nc.Publish("test", []byte("Hello World"))

		time.Sleep(600 * time.Millisecond)

		resp, err := nc.Request("request", []byte("veremos que onda"), 1*time.Second)
		if err != nil {
			log.Println(err)
		}

		log.Println(string(resp.Data))
	}
}
