package queue

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"hegelscheduler/internal/config"
	"log"
	"time"
)

const (
	exchange       = "hegel.scheduler"
	reconnectDelay = 5
)

type RabbitClient struct {
	conn        *amqp.Connection
	channel     *amqp.Channel
	exchanges   map[string]string
	queues      map[string]string
	notifyClose chan *amqp.Error
	ctx         context.Context
	bs          *config.BootStrap
}

func NewRabbitProductor(bs *config.BootStrap) (Productor, func(), error) {

	client := &RabbitClient{
		bs:        bs,
		exchanges: make(map[string]string),
		queues:    make(map[string]string),
		ctx:       context.Background(),
	}
	cleanUp := func() {
		client.Close()
	}
	if err := client.Connect(); err != nil {
		return nil, cleanUp, err
	}
	return client, cleanUp, nil
}

func NewRabbitConsumer(bs *config.BootStrap) (Consumer, func(), error) {
	client := &RabbitClient{
		bs:        bs,
		exchanges: make(map[string]string),
		queues:    make(map[string]string),
		ctx:       context.Background(),
	}
	cleanUp := func() {
		client.Close()
	}
	if err := client.Connect(); err != nil {
		return nil, cleanUp, err
	}
	return client, cleanUp, nil
}

func (r *RabbitClient) Connect() error {
	var err error
	r.conn, err = amqp.Dial(r.bs.Data.RabbitMQ)
	if err != nil {
		return err
	}

	r.notifyClose = make(chan *amqp.Error)
	r.conn.NotifyClose(r.notifyClose)

	r.channel, err = r.conn.Channel()
	if err != nil {
		r.conn.Close()
		return err
	}

	go r.handleReconnect()
	return nil
}

func (r *RabbitClient) handleReconnect() {
	for {
		select {
		case <-r.ctx.Done():
			return
		case err := <-r.notifyClose:
			if err != nil {
				log.Printf("Connection closed: %v", err)
			}

			for {
				select {
				case <-r.ctx.Done():
					return
				default:
					log.Println("Attempting to reconnect...")

					if err := r.Connect(); err == nil {
						log.Println("Reconnected successfully")
						return
					}

					time.Sleep(reconnectDelay * time.Second)
				}
			}
		}
	}
}

func (r *RabbitClient) Publish(data any, topic string) error {
	if err := r.exChangeDeclareAndBind(topic); err != nil {
		return err
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := r.channel.Publish(exchange, topic, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Body:         jsonData,
		ContentType:  "application/json",
	}); err != nil {
		return err
	}
	return nil
}

func (r *RabbitClient) exChangeDeclareAndBind(topic string) error {
	if _, exist := r.exchanges[exchange]; !exist {
		if err := r.channel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil); err != nil {
			return err
		}
		r.exchanges[exchange] = exchange
	}
	if _, exist := r.queues[topic]; !exist {
		if _, err := r.channel.QueueDeclare(topic, true, false, false, false, nil); err != nil {
			return err
		}
		if err := r.channel.QueueBind(topic, topic, exchange, false, nil); err != nil {
			return nil
		}
		r.queues[topic] = topic
	}
	return nil
}

func (r *RabbitClient) Subscribe(topic string, handler func(data []byte)) error {
	if err := r.exChangeDeclareAndBind(topic); err != nil {
		return err
	}
	msgs, err := r.channel.Consume(topic, "", true, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for d := range msgs {
			handler(d.Body)
		}
	}()
	return nil
}

func (r *RabbitClient) Close() error {
	r.channel.Close()
	r.conn.Close()
	return nil
}
