package mq

import (
	"context"
	"encoding/json"

	"github.com/TranThang-2804/infrastructure-engine/internal/domain"
	"github.com/TranThang-2804/infrastructure-engine/internal/shared/log"
	"github.com/TranThang-2804/infrastructure-engine/internal/utils"
)

type CompositeResourcePublisher struct {
	messageQueue MessageQueue
}

func NewCompositeResourcePublisher(mq MessageQueue) domain.CompositeResourceEventPublisher {
	return &CompositeResourcePublisher{
		messageQueue: mq,
	}
}

func (cr *CompositeResourcePublisher) PublishToPendingSubject(
	c context.Context,
	compositeResource domain.CompositeResource,
) error {
	logger := log.BaseLogger.WithFields("infrastructure", utils.GetStructName(cr))

	// Publish message to queue
	messageData, err := json.Marshal(compositeResource)
	if err != nil {
		logger.Error("Error marshalling composite resource to JSON", "error", err)
		return err
	}
	cr.messageQueue.Publish("composite-resource.pending", messageData)
	logger.Info("Publish to pending subject", "compositeResourceId", compositeResource.Id)
	return nil
}

func (cr *CompositeResourcePublisher) PublishToProvisioningSubject(
	c context.Context,
	compositeResource domain.CompositeResource,
) error {
	logger := log.BaseLogger.WithFields("infrastructure", utils.GetStructName(cr))

	// Publish message to queue
	messageData, err := json.Marshal(compositeResource)
	if err != nil {
		logger.Error("Error marshalling composite resource to JSON", "error", err)
		return err
	}
	return cr.messageQueue.Publish("composite-resource.provisioning", messageData)
}

func (cr *CompositeResourcePublisher) PublishToDeletingSubject(
	c context.Context,
	compositeResource domain.CompositeResource,
) error {
	logger := log.BaseLogger.WithFields("infrastructure", utils.GetStructName(cr))

	// Publish message to queue
	messageData, err := json.Marshal(compositeResource)
	if err != nil {
		logger.Error("Error marshalling composite resource to JSON", "error", err)
		return err
	}
	return cr.messageQueue.Publish("composite-resource.deleting", messageData)
}
