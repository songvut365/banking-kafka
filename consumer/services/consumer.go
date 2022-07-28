package services

import "github.com/Shopify/sarama"

type consumerHandler struct {
	eventHandler EventHandler
}

func NewConsumerHandler(eventHandler EventHandler) sarama.ConsumerGroupHandler {
	return consumerHandler{eventHandler}
}

// Don't use but have to implement
func (handler consumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Don't use but have to implement
func (handler consumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (handler consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		handler.eventHandler.Handle(msg.Topic, msg.Value)
		session.MarkMessage(msg, "")
	}

	return nil
}
