package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// kafka consumer

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	// sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags) // 启用详细日志记录
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions("topic-erdan-one") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println("分区", partitionList)

	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("topic-erdan-one", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		for msg := range pc.Messages() {
			fmt.Printf("Partition:%d Offset:%d Key:%v Value:%s \n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
		}
		// 异步从每个分区消费信息
		// go func(sarama.PartitionConsumer) {
		// 	for msg := range pc.Messages() {
		// 		fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v", msg.Partition, msg.Offset, msg.Key, msg.Value)
		// 	}
		// }(pc)
	}
}
