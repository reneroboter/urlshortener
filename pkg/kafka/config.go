package kafka

import (
	"os"

	"github.com/twmb/franz-go/pkg/kgo"
)

func NewKafkaConfig() []kgo.Opt {
	return []kgo.Opt{
		kgo.SeedBrokers(os.Getenv("KAFKA_HOST")),
		kgo.ClientID(os.Getenv("KAFKA_CLIENT")),
		kgo.AllowAutoTopicCreation(),
	}
}
