package domainN

import notificationEntity "consumer/src/notifications/domain/entity"

type NotificationInterface interface {
	SendNotification(notification *notificationEntity.Notification) error
	SaveNotification(msg string) error
}
