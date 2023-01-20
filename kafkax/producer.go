package kafkax

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/kkakoz/pkg/logger"
	"github.com/spf13/viper"
)

type producerOptions struct {
	Address []string
}

func NewSyncProducer(viper *viper.Viper) (sarama.SyncProducer, error) {
	o := &producerOptions{}
	viper.SetDefault("kafka.address", []string{"127.0.0.1:9092"})
	err := viper.UnmarshalKey("kafka", o)
	if err != nil {
		return nil, err
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal        // ack确认机制
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 选择分区-随机分区
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	//config.Producer.Idempotent = true                         // 幂等性, 重复数据只持久化一条
	sarama.Logger = Logger{}

	// 连接kafka
	client, err := sarama.NewSyncProducer(o.Address, config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

//func SendSyncMsgByte(client sarama.SyncProducer, topic string, data []byte) error {
//	msg := &sarama.ProducerMessage{
//		Topic: topic,
//		Value: sarama.ByteEncoder(data),
//	}
//	_, _, err := client.SendMessage(msg)
//	return errors.Wrap(err, "发送消息失败")
//}

type Logger struct {
}

func (l Logger) Print(v ...interface{}) {
	logger.Error(fmt.Sprint(v...))
}

func (l Logger) Printf(format string, v ...interface{}) {
	logger.Error(fmt.Sprintf(format, v...))
}

func (l Logger) Println(v ...interface{}) {
	logger.Error(fmt.Sprintln(v...))
}
