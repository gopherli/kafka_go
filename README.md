# kafka_go
learn

# 启动测试
- 启动生产者：go run cmd/producer/main.go
- 启动消费者：go run cmd/consumer/main.go
# win10下环境搭建
- 参考文章：https://www.confluent.io/blog/set-up-and-run-kafka-on-windows-linux-wsl-2/
- 相关操作命令(cd kafka...)
    - zookeeper启动：bin/zookeeper-server-start.sh config/zookeeper.properties
    - kafak启动:     bin/kafka-server-start.sh config/server.properties
    - topic创建：    bin/kafka-topics.sh --create --topic quickstart-events --bootstrap-server localhost:9092
    - producer创建生产：bin/kafka-console-producer.sh --topic topic-erdan-one --bootstrap-server localhost:9092
    - consumer创建消费：bin/kafka-console-consumer.sh --topic topic-erdan-one --bootstrap-server localhost:9092
    - topics查看：      bin/kafka-topics.sh --list --zookeeper localhost:2181
- 注意：server.properties 指定的listeners=PLAINTEXT://:9092、advertised.listeners=PLAINTEXT://localhost:9092地址与创建的topic地址一致才能访问通。

# golang操作kafka
- 参考文章：https://www.liwenzhou.com/posts/Go/go_kafka/
- 下载包：go get github.com/Shopify/sarama
- producer生产：
    - 指定好生产的topic:msg.Topic = "topic-erdan-one"
    - 指定好对应服务器的端口：client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
- consumer消费：
    - 指定好对应服务器的端口：consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
    - 指定好从哪个topic取分区：partitionList, err := consumer.Partitions("topic-erdan-one") // 根据topic取到所有的分区
    - 指定好与生产对应的消费topic：pc, err := consumer.ConsumePartition("topic-erdan-one", int32(partition), sarama.OffsetNewest)
- 注意：服务器的端口三个一致：server.properties地址、topic的地址、go指定的地址三个要一一致
