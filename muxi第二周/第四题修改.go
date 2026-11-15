
package main

import (
	"fmt"
	"time"
)

type message struct {
	Topic     string
	Partition int32
	Offset    int64
}

type FeedEventDM struct {
	Type    string
	UserID  int
	Title   string
	Content string
}

type MSG struct {
	ms        message
	feedEvent FeedEventDM
}

const ConsumeNum = 5

func main() {
	var consumeMSG []MSG
	var lastConsumeTime time.Time

	msgs := make(chan MSG)

	// 生产消息
	go func() {
		for i := 0; ; i++ {
			msgs <- MSG{
				ms: message{
					Topic:     "消费主题",
					Partition: 0,
					Offset:    0,
				},
				feedEvent: FeedEventDM{
					Type:    "grade",
					UserID:  i,
					Title:   "成绩提醒",
					Content: "您的成绩是xxx",
				},
			}

			// 模拟生产消息的间隔
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// 不断接收消息
	for msg := range msgs {

		// 将新消息添加到切片
		consumeMSG = append(consumeMSG, msg)

		// 达到5条，批量消费
		if len(consumeMSG) >= ConsumeNum {

			// 复制出独立的一批数据，避免异步消费时发生数据竞争
			m := make([]MSG, ConsumeNum)
			copy(m, consumeMSG[:ConsumeNum])

			// 异步消费
			go fn(m)

			// 移除已经消费的5条消息
			consumeMSG = consumeMSG[ConsumeNum:]

			// 更新消费时间
			lastConsumeTime = time.Now()

		} else if !lastConsumeTime.IsZero() &&
			time.Since(lastConsumeTime) > 5*time.Minute {

			// 超过5分钟，有未消费的消息
			if len(consumeMSG) > 0 {

				// 剩余多少条就消费多少条
				m := make([]MSG, len(consumeMSG))
				copy(m, consumeMSG)

				go fn(m)

				// 清空待消费切片
				consumeMSG = consumeMSG[:0]

				lastConsumeTime = time.Now()
			}
		}
	}
}

// 消费消息
func fn(m []MSG) {
	fmt.Printf("本次消费了%d条消息\n", len(m))
}
