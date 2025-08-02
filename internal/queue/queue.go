package queue

type Productor interface {
	Publish(data any, topic string) error
}

type Consumer interface {
	Subscribe(topic string, handler func(data []byte)) error
}
