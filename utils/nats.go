package utils

import (
	"log"

	"github.com/nats-io/nats.go"
)

func InitNATS() *nats.Conn {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	return nc
}
