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
	Username string `config:"RABBITMQ_USER" default:"rabbitmq"`
	Password string `config:"RABBITMQ_PASSWORD" default:"rabbitmq"`
	Host string `config:"RABBITMQ_HOST" default:"rabbitmq"`
	Port int `config:"RABBITMQ_PORT" default:"5672"`
}

func Load_config() *RabbitMQConfig {
	var conf RabbitMQConfig
	config.Load(&conf)
	return &conf
}

func Connect_rabbitMQ(conf *RabbitMQConfig) (*RabbitMQ, error) {
	rabbitMQURL := fmt.Sprintf("amqp://%s:%s@%s:%d", conf.Username, conf.Password, conf.Host, conf.Port)

	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	rabbitMQ := &RabbitMQ {
		conn: conn,
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