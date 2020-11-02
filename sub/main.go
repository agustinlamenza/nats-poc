package main

import (
	"fmt"
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

	nc.Subscribe("request", func(m *nats.Msg) {
		str := fmt.Sprintf("Esta es la respuesta %v", time.Now())
		nc.Publish(m.Reply, []byte(str))
	})

	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe("test", ch)

	for e := range ch {
		log.Println(string(e.Data))
	}

	sub.Unsubscribe()
}
