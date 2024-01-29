package rabbitmq

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

// var (
// 	host, port, user, pass, vhost string
// )

const (
	TaskEventExchange = "task_event_v1"
	TaskEventQueue    = "task.event"
	TaskRoutingKey    = "task.event.key"
)

var conn *amqp.Connection

func Init() error {
	_, err := GetConnection()
	return err
}

func connect() error {
	if conn == nil || conn.IsClosed() {
		url := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", viper.GetString("rabbitMQ.user"), viper.GetString("rabbitMQ.pass"), viper.GetString("rabbitMQ.host"), viper.GetString("rabbitMQ.port"), viper.GetString("rabbitMQ.vHost"))
		fmt.Println("rabbbitmq end-point- ", url)
		var err error
		conn, err = amqp.Dial(url)
		if err != nil {
			fmt.Println("123rabbbitmq end-point- ", url)
			//log.Println(conn.Config.Vhost, err)
			return err
		}
	}
	return nil
}

// GetConnection ..
func GetConnection() (*amqp.Connection, error) {

	if err := connect(); err != nil {
		return nil, err
	}
	if conn == nil {
		return nil, errors.New("rabbitmq:could not establish connection")
	}
	return conn, nil
}
