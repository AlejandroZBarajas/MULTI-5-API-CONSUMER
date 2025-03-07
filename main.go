package main

import (
	"consumer/src/core/rabbit/infrastructureR"
	//"consumer/src/notifications/application"
	//"consumer/src/notifications/domain"
	applicationN "consumer/src/notifications/application"
	infrastructureN "consumer/src/notifications/infrastructure"
	"fmt"
	"log"
	"net/http"
)

func main() {
	rabbitClient, err := infrastructureR.NewRabbitMQ()
	if err != nil {
		log.Fatalf("Error iniciando RabbitMQ: %v", err)
	}
	defer rabbitClient.Close()

	repo := infrastructureN.NewNotificationRepository()
	createNotificationUseCase := applicationN.NewCreateNotification(repo)

	notificationController := infrastructureN.NewNotificationController(createNotificationUseCase, rabbitClient)

	http.HandleFunc("/notifications", notificationController.CreateNewHandler)

	port := ":8080"
	fmt.Printf("Servidor API corriendo en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
