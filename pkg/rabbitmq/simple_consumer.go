package rabbitmq

import (
	"fmt"
	"spider/pkg/logger"
	"time"
)

// Consumer 消费普通队列中的消息
func (r *RabbitMQ) Consumer(qName string, handler func([]byte) error) (err error) {
	if err = r.CheckAndReconnect(); err != nil {
		return err
	}
	// 声明队列（确保队列存在）
	_, err = r.ch.QueueDeclare(
		qName, // 队列名称
		true,  // 非持久化（根据需求修改）
		false, // 不自动删除
		false, // 排他性
		false, // 不等待响应
		nil,   // 无额外参数
	)
	if err != nil {
		return fmt.Errorf("声明队列失败: %w", err)
	}

	logger.Logger.Info(fmt.Sprintf("开始监听队列: %s", qName))

	// 创建消费者（autoAck设为false，手动确认）
	// 生成有效的消费者标签，避免特殊字符
	consumerTag := fmt.Sprintf("consumer-%s-%d", qName, time.Now().Unix())
	msgs, err := r.ch.Consume(
		qName,       // 队列名称
		consumerTag, // 消费者标签
		false,       // 手动确认消息
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("创建消费者失败: %v", err))
		return err
	}

	// 启动协程处理消息
	go func() {
		for d := range msgs {
			time.Sleep(1 * time.Second)

			// 获取MessageId
			messageId := d.MessageId
			if messageId == "" {
				messageId = "unknown"
			}

			// 处理消息
			if err := handler(d.Body); err != nil {
				logger.Logger.Error(fmt.Sprintf("处理消息失败: MessageId=%s, 错误: %v, 内容: %s", messageId, err, string(d.Body)))
				// 拒绝消息并重新入队
				d.Nack(false, true)
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
