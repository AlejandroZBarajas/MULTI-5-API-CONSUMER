package infrastructureN

import (
	//domainN "consumer/src/notifications/domain"
	notificationEntity "consumer/src/notifications/domain/entity"
	"fmt"
)

type NotificationRepository struct{}

func NewNotificationRepository() *NotificationRepository {
	return &NotificationRepository{}
}

func (r *NotificationRepository) SaveNotification(msg string) error {
	fmt.Println("Notificación guardada:", msg)
	return nil
}

func (r *NotificationRepository) SendNotification(notification *notificationEntity.Notification) error {
	fmt.Println("Notificación enviada:", notification.Message)
	return nil
}

//var _ domainN.NotificationInterface = (*NotificationRepository)(nil)
