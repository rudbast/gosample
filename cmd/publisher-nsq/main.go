package main

import (
	"encoding/json"
	"log"

	nsq "github.com/nsqio/go-nsq"
)

type Foo struct {
	Bar string `json:"bar"`
}

func main() {
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	foo := Foo{Bar: "bar"}
	byteFoo, _ := json.Marshal(foo)

	err = producer.Publish("random-topic", byteFoo)
	if err != nil {
		log.Println("Error publish:", err)
	}
}
