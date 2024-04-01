package tests

import (
	kafka "api-gateway/api/api_test/kafkatest"
	"testing"

	"github.com/IBM/sarama/mocks"
)

func TestProducer(t *testing.T) {
	mockProducer := mocks.NewSyncProducer(t, nil)
	defer func() {
		if err := mockProducer.Close(); err != nil {
			t.Error(err)
		}
	}()
	producer := kafka.NewMockKafka(mockProducer)
	messageToProduce := "Kafka mock test"
	mockProducer.ExpectSendMessageAndSucceed()
	err := producer.SendMessageToKafka(messageToProduce)
	if err != nil {
		t.Errorf("Error is sending message to Kafka: %v", err)
	}
}
