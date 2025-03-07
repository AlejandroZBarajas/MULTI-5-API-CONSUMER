package notificationEntity

type Notification struct {
	Message string `json:"message"`
}

func Notify(msg string) *Notification {
	return &Notification{Message: msg}
}
