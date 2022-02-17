package bootstrap

import (
	"System/app/services"
	"fmt"
)

var isConsumerOpen = false

// ConsumerOpen 开启管道的消费者
func ConsumerOpen() {
	if !isConsumerOpen {
		go services.BookConsumer()
		isConsumerOpen = true
		fmt.Println("已打开消费者")
	}
}
