package main

import (
	"fmt"
	"os"
	rabbitmq "task-manger-service/client/rabbitmq/configuration"
	"task-manger-service/client/rabbitmq/consumer"
	"task-manger-service/db"
	"task-manger-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")

	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		os.Exit(0)
	}
	router := gin.Default()
	routes.InitRoutes(router)
	err := db.InitMySQL()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	// initializing rabbitMq
	err = rabbitmq.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	consumer.StartConsumers()

	err = router.Run(":" + viper.GetString("server.port"))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
