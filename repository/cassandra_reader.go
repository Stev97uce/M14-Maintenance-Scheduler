package repository

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

type CassandraReader struct {
	Session *gocql.Session
}

func NewCassandraReader(host, keyspace string) *CassandraReader {
	cluster := gocql.NewCluster(host)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("❌ Error conectando a Cassandra: %v", err)
	}

	log.Println("✅ Conectado a Cassandra")
	return &CassandraReader{Session: session}
}

type MachineEvent struct {
	EventID   gocql.UUID
	MachineID string
	Type      string
	Timestamp time.Time
	Data      string
}

func (r *CassandraReader) GetRecentEvents(since time.Time) ([]MachineEvent, error) {
	var events []MachineEvent

	iter := r.Session.Query(`
		SELECT event_id, machine_id, type, timestamp, data 
		FROM machine_events 
		WHERE timestamp > ? ALLOW FILTERING`, since).Iter()

	for {
		var e MachineEvent
		if !iter.Scan(&e.EventID, &e.MachineID, &e.Type, &e.Timestamp, &e.Data) {
			break
		}
		events = append(events, e)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return events, nil
}
