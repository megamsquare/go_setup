package queue

import (
	"fmt"
	"log"

	config "github.com/megamsquare/go_setup/pkg/env_config"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
}

type RabbitMQConfig struct {
	URI string `config:"AMQP_URI" default:"amqp://guest:guest@localhost:5672/"`
}

func Load_config() *RabbitMQConfig {
	var conf RabbitMQConfig
	config.Load(&conf)
	return &conf
}

func Connect_rabbitMQ(conf *RabbitMQConfig) (*RabbitMQ, error) {
	conn := fmt.Sprintf(conf.URI)

	rbmq, err := amqp.Dial(conn)
	if err != nil {
		return nil, err
	}

	channel, err := rbmq.Channel()
	if err != nil {
		return nil, err
	}

	rabbitMQ := &RabbitMQ {
		conn: rbmq,
		channel: channel,
	}

	return rabbitMQ, nil
}

func (r *RabbitMQ) Close() {
	if r.channel != nil {
		err := r.channel.Close()
		if err != nil {
			log.Printf("Error closing RabbitMQ channel: %v", err)
		}
	}

	if r.conn != nil {
		err := r.conn.Close()
		if err != nil {
			log.Printf("Error closing RabbitMQ connectio: %v", err)
		}
	}
}