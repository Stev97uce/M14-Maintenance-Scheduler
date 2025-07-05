package services

import (
	"context"
	"log"
	"maintenance-scheduler/models"
	"maintenance-scheduler/repository"
	"time"
)

type SchedulerService struct {
	Repo *repository.MaintenanceRepository
}

func NewSchedulerService(repo *repository.MaintenanceRepository) *SchedulerService {
	return &SchedulerService{Repo: repo}
}

func (s *SchedulerService) HandleEvent(machineID, eventType, details string) {
	task := models.MaintenanceTask{
		MachineID: machineID,
		Timestamp: time.Now(),
		Type:      eventType,
		Details:   details,
	}

	_, err := s.Repo.Collection.InsertOne(context.Background(), task)
	if err != nil {
		log.Printf("❌ Error guardando mantenimiento: %v", err)
	} else {
		log.Printf("🛠️ Tarea registrada para máquina %s (%s)", machineID, eventType)
	}
}
