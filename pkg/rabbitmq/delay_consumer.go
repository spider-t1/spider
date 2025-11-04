package rabbitmq

import (
	"fmt"
	"spider/pkg/logger"
	"time"
)

// ConsumeDelayQueue 消费延迟消息（从目标队列消费）
func (r *RabbitMQ) ConsumeDelayQueue(qName string, handler func([]byte) error) error {
	if err := r.CheckAndReconnect(); err != nil {
		return err
	}
	// 声明目标队列（确保存在）
	_, err := r.ch.QueueDeclare(
		qName, // 目标队列名称
		true,  // 持久化
		false, // 不自动删除
		false, // 非排他性
		false, // 不等待响应
		nil,   // 无额外参数
	)
	if err != nil {
		return fmt.Errorf("声明目标队列失败: %w", err)
	}

	logger.Logger.Info(fmt.Sprintf("开始监听目标队列: %s", qName))

	// 创建消费者（autoAck设为false，手动确认）
	// 生成有效的消费者标签，避免特殊字符
	consumerTag := fmt.Sprintf("delay-consumer-%s-%d", qName, time.Now().Unix())
	msgs, err := r.ch.Consume(
		qName,       // 队列名称
		consumerTag, // 消费者标签
		false,       // ✅ 关键修正：手动确认消息
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("创建消费者失败: %w", err)
	}

	// 启动协程处理消息
	go func() {
		for d := range msgs {
			// 获取MessageId
			messageId := d.MessageId
			if messageId == "" {
				messageId = "unknown"
			}

			// 处理消息
			if err := handler(d.Body); err != nil {
				logger.Logger.Error(fmt.Sprintf("处理消息失败: MessageId=%s, 错误: %v, 内容: %s", messageId, err, string(d.Body)))
				// 拒绝消息并重新入队
				d.Ack(false)
				//	 改成重新入队
				r.PublishDelayMessage(qName, d.Body, 5000)
			} else {
				// 手动确认消息
				d.Ack(false)
				logger.Logger.Info(fmt.Sprintf("处理消息成功: MessageId=%s, 内容: %s", messageId, string(d.Body)))
			}
		}
		logger.Logger.Info(fmt.Sprintf("消费者协程已关闭，队列: %s", qName))
	}()

	return nil
}
