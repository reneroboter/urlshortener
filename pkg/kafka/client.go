package kafka

import (
	"github.com/twmb/franz-go/pkg/kgo"
)

func NewKafkaClient() (*kgo.Client, error) {
	return kgo.NewClient(NewKafkaConfig()...)
}
