package kafkax_test

import (
	"github.com/Shopify/sarama"
	"github.com/kkakoz/pkg/kafkax"
	"github.com/kkakoz/pkg/logger"
	"github.com/spf13/viper"
	"testing"
)

func TestProduct(t *testing.T) {
	viper.SetConfigFile("./conf.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		t.Fatal(err)
	}
	logger.InitLog(viper.GetViper())
	producer, err := kafkax.NewSyncProducer(viper.GetViper())
	if err != nil {
		t.Fatal(err)
	}

	msg := &sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.StringEncoder("hello"),
		//Key:   sarama.StringEncoder(strconv.Itoa(11)),
	}

	_, _, err = producer.SendMessage(msg)
	if err != nil {
		t.Fatal(err)
	}
}
