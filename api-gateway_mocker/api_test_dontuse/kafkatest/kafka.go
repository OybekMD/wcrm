package kafkatest

import (
	"errors"
	"github.com/IBM/sarama"
)

type MockKafka struct {
	producer sarama.SyncProducer
}

func NewMockKafka(producer sarama.SyncProducer) *MockKafka {
	return &MockKafka{producer: producer}
}

func (m *MockKafka) SendMessageToKafka(message string) error {
	if m.producer == nil {
		return errors.New("producer is not initialized")
	}

	msg := &sarama.ProducerMessage{
		Topic: "kafka-p-mock",
		Value: sarama.StringEncoder(message),
		Partition: -2,
	}

	// Sending
	_, _, err := m.producer.SendMessage(msg)
	return err
}
