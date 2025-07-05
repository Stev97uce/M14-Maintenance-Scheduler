package models

import "time"

type MaintenanceTask struct {
	MachineID string    `bson:"machine_id"`
	Timestamp time.Time `bson:"timestamp"`
	Type      string    `bson:"type"`
	Details   string    `bson:"details"`
}
