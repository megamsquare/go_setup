package queue

import (
	"fmt"

	config "github.com/megamsquare/go_setup/pkg/env_config"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
}

type RabbitMQConfig struct {
	Username string `config:"RABBITMQ_USERNAME" default:"rabbitmq"`
	Password string `config:"RABBITMQ_PASSWORD" default:"rabbitmq"`
	Port int `config:"RABBITMQ_PORT" default:"5672"`
}

func Load_config() *RabbitMQConfig {
	var conf RabbitMQConfig
	config.Load(&conf)
	return &conf
}

func Connect_rabbitMQ(conf *RabbitMQConfig) (*RabbitMQ, error) {
	conn := fmt.Sprintf("")

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
		
	}

	if r.conn != nil {

	}
}