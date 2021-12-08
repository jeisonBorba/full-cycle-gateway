package kafka

import (
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jeisonborba/full-cycle-gateway/adapter/presenter/transaction"
	"github.com/jeisonborba/full-cycle-gateway/domain/entity"
	"github.com/jeisonborba/full-cycle-gateway/usecase/process_transaction"
	"github.com/stretchr/testify/assert"
)

func TestProducerPublish(t *testing.T) {
	expextedOutput := process_transaction.TransactionDtoOutput{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "you dont have limit for this transaction",
	}
	// outputJson, _ := json.Marshal(expextedOutput)

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}
	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())
	err := producer.Publish(expextedOutput, []byte("1"), "test")
	assert.Nil(t, err)
}
