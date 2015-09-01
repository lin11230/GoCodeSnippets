package main

import (
	"ClearWatch2/libs/queue"
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("copyqueue start...")
	hostname, _ := os.Hostname()
	fmt.Println("Hostname is ", hostname)

	// production RabbitMQ
	// var rabbitProd = "amqp://clearwatch:dkqYW7T1rp@cw-prod-rabbit:5672//clearwatch"

	// cw2-rabbit1 RabbitMQ
	// copyRabbit = "amqp://clearwatch:dkqYW7T1rp@cw2-rabbit1:5672/clearwatch"

	// cw2-uat RabbitMQ
	// copyRabbit = "amqp://clearwatch:dkqYW7T1rp@cw2-uat:5672/clearwatch"
	var rabbitProd = "amqp://guest:guest@localhost:5672/"

	// queue && routing key
	var prodRouteKey = "worker.tracker"

	var subCopy <-chan amqp.Delivery

	// connect to prodction RabbitMQ
	mqProd, errProd := queue.NewRabbitMQ(rabbitProd)
	if errProd != nil {
		fmt.Errorf("Not able to create Production's RabbitMQ %s", errProd)
	}

	// multiple-binding to Prodction RabbitMQ
	prodRouteKey = "livefeed.C-bd25d318-24b0-4a39-bcf3-12b163999d0f"
	subCopy, errCopy := mqProd.Subscribe(prodRouteKey, "test.key")
	if errCopy != nil {
		fmt.Errorf("Subscribe New Queue for tracker  error %s", errCopy)
	}

	for res := range subCopy {

		fmt.Println("\n*** test rabbitmq binding\n", string(res.Body), "\n***\n")

	}
}
