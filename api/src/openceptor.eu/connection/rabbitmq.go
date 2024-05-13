package connection

import (
	"fmt"
	"sync"

	"openceptor.eu/config"

	amqp "github.com/rabbitmq/amqp091-go"

	"net/url"
)

var lockRabbitMq = &sync.Mutex{}

type rabbitmq struct {
	dsn string
	conn *amqp.Connection
}

var rabbitmqInstance *rabbitmq

func buildRabbitMqDsn(c *config.Config) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/%s",
		url.QueryEscape(c.Queue.Username),
		url.QueryEscape(c.Queue.Password),
		c.Queue.Host,
		c.Queue.Port,
		url.QueryEscape(c.Queue.Vhost))
}

func openRabbitMq(dsn string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func GetRabbitMqInstance(c *config.Config) *amqp.Connection {
	if rabbitmqInstance != nil && !rabbitmqInstance.conn.IsClosed() {
		fmt.Println("rabbitmqInstance already exist, re-use")

		return rabbitmqInstance.conn;
	}

	lockRabbitMq.Lock()
	defer lockRabbitMq.Unlock()

	if rabbitmqInstance != nil && !rabbitmqInstance.conn.IsClosed() {
		fmt.Println("rabbitmqInstance already exist, re-use (after lock)")

		return rabbitmqInstance.conn;
	}

	fmt.Println("rabbitmqInstance not exist, create it !")
	rabbitmqInstance = &rabbitmq{}

	rabbitmqInstance.dsn = buildRabbitMqDsn(c)

	conn, err := openRabbitMq(rabbitmqInstance.dsn)
	if err != nil {
		panic(err)
	}

	rabbitmqInstance.conn = conn
	
	return rabbitmqInstance.conn
}
