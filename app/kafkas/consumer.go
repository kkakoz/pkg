package kafkas

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/kkakoz/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type consumerOptions struct {
	Address []string
	GroupId string
	Topics  []string
}

type KafkaConsumer struct {
	consumer *cluster.Consumer
	runFunc  ConsumerRunFunc
}

type ConsumerRunFunc func(*sarama.ConsumerMessage) error

func NewConsumer(viper *viper.Viper, runFunc ConsumerRunFunc, config *cluster.Config) (*KafkaConsumer, error) {
	o := &consumerOptions{}
	viper.SetDefault("kafka.address", []string{"127.0.0.1:9092"})
	viper.SetDefault("kafka.topics", []string{"test"})
	err := viper.UnmarshalKey("kafka", o)
	if err != nil {
		return nil, err
	}

	consumer, err := cluster.NewConsumer(o.Address, o.GroupId, o.Topics, config)
	if err != nil {
		return nil, err
	}

	run := &KafkaConsumer{
		consumer: consumer,
		runFunc:  runFunc,
	}

	return run, nil
}

func (c *KafkaConsumer) Start(ctx context.Context) error {
	// consume messages, watch signals
	for {
		select {
		case err := <-c.consumer.Errors():
			logger.Error(fmt.Sprintf("kafka consumer err: %+v\n", err))
		case ntf := <-c.consumer.Notifications():
			logger.Info(fmt.Sprintf("kafka consumer rebalanced: %+v\n", ntf))
		case msg, ok := <-c.consumer.Messages():
			if ok {
				err := c.runFunc(msg)
				if err != nil {
					logger.Error(fmt.Sprintf("consumer msg err: %+v\n", err), zap.String("msg", fmt.Sprintf("%+v", msg)))
				}
				c.consumer.MarkOffset(msg, "") // mark message as processed
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func (c *KafkaConsumer) Stop(ctx context.Context) error {
	return c.consumer.Close()
}
