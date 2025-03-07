package applicationN

import (
	domainN "consumer/src/notifications/domain"
	notificationEntity "consumer/src/notifications/domain/entity"
	"fmt"
)

type CreateNotification struct {
	repo domainN.NotificationInterface
}

func NewCreateNotification(repo domainN.NotificationInterface) *CreateNotification {
	return &CreateNotification{repo: repo}
}

func (cn *CreateNotification) Run(msg string) error {
	notification := notificationEntity.Notify(msg)

	err := cn.repo.SendNotification(notification)

	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}
