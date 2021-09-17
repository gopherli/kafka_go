package main

import (
	"fmt"
	"runtime"

	"github.com/Shopify/sarama"
)

// 基于sarama第三方库开发的kafka client

func main() {
	config := sarama.NewConfig()
	config.ClientID = "kafak-erdan"
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	// sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags) // 启用详细日志记录
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "topic-erdan-one"
	msg.Value = sarama.StringEncoder("this is a test log")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config) //113.89.35.146
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()

	// 发送消息
	done := make(chan string)
	sendCount := 0
	for {
		go func() {
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
			done <- "stop"
		}()
		sendCount++
		if sendCount > 1000 {
			break
		}
	}
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())
	<-done
}
