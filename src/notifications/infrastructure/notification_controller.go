package infrastructureN

import (
	"consumer/src/core/rabbit/infrastructureR"
	applicationN "consumer/src/notifications/application"
	"encoding/json"
	"fmt"
	"net/http"
)

type NotificationController struct {
	CreateNotificationUseCase *applicationN.CreateNotification
	RabbitClient              *infrastructureR.RabbitMQ
}

func NewNotificationController(
	create *applicationN.CreateNotification,
	rabbitClient *infrastructureR.RabbitMQ,
) *NotificationController {
	return &NotificationController{
		CreateNotificationUseCase: create,
		RabbitClient:              rabbitClient,
	}
}

func (nc *NotificationController) CreateNewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	var notification struct {
		Msg string `json:"msg"`
	}

	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	fmt.Printf("Datos recibidos: %v\n", notification.Msg)

	err = nc.CreateNotificationUseCase.Run(notification.Msg)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al registrar evento: %v", err), http.StatusInternalServerError)
		return
	}

	eventNotification := map[string]interface{}{
		"device_name": notification.Msg,
		"message":     "Evento registrado desde dispositivo",
	}

	eventNotificationJSON, err := json.Marshal(eventNotification)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al convertir mensaje a JSON: %v", err), http.StatusInternalServerError)
		return
	}

	err = nc.RabbitClient.PublishMessage("event_queue", eventNotificationJSON)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al publicar mensaje en RabbitMQ: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Evento registrado")))

}
