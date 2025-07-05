package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MaintenanceRepository struct {
	Collection *mongo.Collection
}

func NewMaintenanceRepository(mongoURI, dbName string) *MaintenanceRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("❌ Error conectando a MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("❌ MongoDB no responde: %v", err)
	}

	log.Println("✅ Conectado a MongoDB")

	collection := client.Database(dbName).Collection("maintenance_tasks")
	return &MaintenanceRepository{Collection: collection}
}
