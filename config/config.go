package config

import (
	"log"
	"os"
)

type Config struct {
	Port              string
	MongoURI          string
	MongoDB           string
	KafkaBroker       string
	KafkaTopic        string
	CassandraHost     string
	CassandraKeyspace string
}

func LoadConfig() Config {
	return Config{
		Port:              getEnv("PORT", "8083"),
		MongoURI:          getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:           getEnv("MONGO_DB", "maintenance_db"),
		KafkaBroker:       getEnv("KAFKA_BROKER", "localhost:9092"),
		KafkaTopic:        getEnv("KAFKA_TOPIC", "machine-events"),
		CassandraHost:     getEnv("CASSANDRA_HOST", "localhost"),
		CassandraKeyspace: getEnv("CASSANDRA_KEYSPACE", "machine_states"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Printf("⚠️  %s no definido, usando valor por defecto: %s", key, fallback)
	return fallback
}
