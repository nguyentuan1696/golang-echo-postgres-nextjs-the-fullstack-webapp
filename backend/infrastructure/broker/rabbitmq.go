package broker

import amqp "github.com/rabbitmq/amqp091-go"

type AmqpClient struct {
	*amqp.Connection
}

var amqpInstance *AmqpClient
