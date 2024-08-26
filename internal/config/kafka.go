package config

import (
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewKafkaReader(config *viper.Viper, log *logrus.Logger) *kafka.Reader {
	readerConfig := kafka.ReaderConfig{
		Brokers:     config.GetStringSlice("kafka.servers"),
		GroupID:     config.GetString("kafka.group.id"),
		GroupTopics: config.GetStringSlice("kafka.consumers"),
	}
	return kafka.NewReader(readerConfig)
}

func NewKafkaWriter(config *viper.Viper, log *logrus.Logger) *kafka.Writer {
	return &kafka.Writer{
		Addr:                   kafka.TCP(config.GetStringSlice("kafka.servers")...),
		AllowAutoTopicCreation: true,
	}
}
