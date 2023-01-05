package kafkas_test

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	app2 "github.com/kkakoz/pkg/app"
	"github.com/kkakoz/pkg/app/kafkas"
	"github.com/kkakoz/pkg/logger"
	"github.com/spf13/viper"
	"testing"
)

func TestNewConsumer(t *testing.T) {
	viper.SetConfigFile("./conf.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		t.Fatal(err)
	}
	logger.InitLog(viper.GetViper())
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	consumer1, err := kafkas.NewConsumer(viper.GetViper(), func(message *sarama.ConsumerMessage) error {
		fmt.Println("message = ", string(message.Value))
		return nil
	}, config)
	if err != nil {
		t.Fatal(err)
	}

	viper.SetConfigFile("./conf2.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		t.Fatal(err)
	}
	consumer2, err := kafkas.NewConsumer(viper.GetViper(), func(message *sarama.ConsumerMessage) error {
		fmt.Println("message = ", string(message.Value))
		return nil
	}, config)
	if err != nil {
		t.Fatal(err)
	}

	app := app2.NewApp("test", consumer1, consumer2)

	err = app.Start(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
}
