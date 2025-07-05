package kafka

import (
	"context"
	"encoding/json"
	"log"
	"maintenance-scheduler/services"

	"github.com/segmentio/kafka-go"
)

type Event struct {
	MachineID string                 `json:"machine_id"`
	Type      string                 `json:"type"`
	Data      map[string]interface{} `json:"data"`
}

func StartConsumer(broker, topic string, handler *services.SchedulerService) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "maintenance-group",
	})

	log.Println("📡 Escuchando eventos desde Kafka...")

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("❌ Error leyendo mensaje: %v", err)
			continue
		}

		var event Event
		err = json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Printf("❌ Error parseando evento: %v", err)
			continue
		}

		log.Printf("📥 Evento recibido: %s para %s", event.Type, event.MachineID)
		handler.HandleEvent(event.MachineID, event.Type, "Evento crítico recibido")
	}
}
