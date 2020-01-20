package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

var (
	wg sync.WaitGroup
)

func main() {
	consumer, err := sarama.NewConsumer(strings.Split("117.51.148.112:9092", ","), nil)
	if err != nil {
		fmt.Println("failed to start consumer:", err)
		return
	}
	defer consumer.Close()

	partitionList, err := consumer.Partitions("nginx_log")
	if err != nil {
		fmt.Println("Failed to get the list of partitions:", err)
		return
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("nginx_log", int32(partition), sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d:%s\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		go func(partitionConsumer sarama.PartitionConsumer) {
			wg.Add(1)
			for msg := range partitionConsumer.Messages() {
				fmt.Printf("partition:%d Offset:%d Key:%s Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
			wg.Done()
		}(pc)
	}

	time.Sleep(1 * time.Second)
	wg.Wait()
}
