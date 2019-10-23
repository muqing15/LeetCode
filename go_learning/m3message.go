package go_learning

import (
	"fmt"
	"time"
)

type Producer interface {
	Product()
}

type Consumer interface {
	Consume()
}

type StringProducer struct {
	msgChan chan string
}

type StringConsumer struct {
	msgChan chan string
}

func NewStringProducer(ch chan string) Producer {
	return &StringProducer{
		msgChan: ch,
	}
}

func NewStringConsumer(ch chan string) Consumer {
	return &StringConsumer{
		msgChan: ch,
	}
}

func (s *StringProducer) Product() {
	tic := time.NewTicker(time.Second)
	for t := range tic.C {
		s.msgChan <- fmt.Sprintf("product message time:%v", t)
	}
}

func (s *StringConsumer) Consume() {
	toc := time.NewTicker(time.Second * 3)
	for t := range toc.C {
		str := ""
		for i := 0; i < 3; i++ {
			str += fmt.Sprintf("consume message time:%v message%d:%s\n", t, i+1, <-s.msgChan)
		}
		fmt.Println(str)
	}
}
