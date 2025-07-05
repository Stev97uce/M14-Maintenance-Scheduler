package main

import (
	"log"
	"time"

	"maintenance-scheduler/config"
	"maintenance-scheduler/repository"
	"maintenance-scheduler/services"
)

func main() {
	cfg := config.LoadConfig()

	repo := repository.NewMaintenanceRepository(cfg.MongoURI, cfg.MongoDB)
	service := services.NewSchedulerService(repo)

	cassandra := repository.NewCassandraReader(cfg.CassandraHost, cfg.CassandraKeyspace)

	log.Println("🟢 Maintenance Scheduler iniciado")

	go func() {
		lastCheck := time.Now().Add(-1 * time.Minute) // consulta eventos del último minuto

		for {
			events, err := cassandra.GetRecentEvents(lastCheck)
			if err != nil {
				log.Printf("❌ Error leyendo de Cassandra: %v", err)
				continue
			}

			for _, ev := range events {
				if ev.Type == "machine_error" || ev.Type == "machine_started" {
					service.HandleEvent(ev.MachineID, ev.Type, "Evento leído desde Cassandra")
				}
			}

			lastCheck = time.Now()
			time.Sleep(30 * time.Second)
		}
	}()

	select {}
}
