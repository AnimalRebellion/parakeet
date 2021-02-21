package proxy

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

type Server struct {
	nc *nats.Conn
}

func (s *Server) Connect() error {
	var err error
	var conn *nats.Conn
	uri := os.Getenv("NATS_URI")

	for i := 0; i < 5; i++ {
		conn, err = nats.Connect(uri)
		if err == nil {
			s.nc = conn
			break
		}
		log.Printf("Waiting before connecting to NATS at: %s", uri)
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		log.Printf("Error establishing connection to NATS: %s", err)
		return err
	}
	log.Printf("Connected to NATS at: %s", s.nc.ConnectedUrl())

	return nil
}

func (s *Server) Send(topic string, payload interface{}) error {
	var err error
	//requestAt := time.Now()
	//response, err := s.nc.Request("tasks", []byte("help please"), 5*time.Second)
	b, err := json.Marshal(payload)
	if err != nil {
		err = fmt.Errorf("Error deserialising: %w", err)
		log.Fatal(err)
		return err
	}
	_, err = s.nc.Request(topic, []byte(b), 5*time.Second)
	if err != nil {
		err = fmt.Errorf("Error making NATS request: %w", err)
		log.Fatal(err)
		return err
	}
	return nil
	//duration := time.Since(requestAt)

	//fmt.Fprintf(w, "Task scheduled in %+v\nResponse: %v\n", duration, string(response.Data))
}

func (s *Server) Receive(topic string) error {
	var err error
	_, err = s.nc.Subscribe(topic, func(m *nats.Msg) {
		var person Person
		err = json.Unmarshal(m.Data, &person)
		if err != nil {
			err = fmt.Errorf("Error deserialising: %w", err)
			log.Println(err)
		}
		log.Printf("Message received on: %s - %s", m.Subject, person.Name)
		err = s.nc.Publish(m.Reply, []byte("Success!!"))
	})
	return fmt.Errorf("Error creating subscription: %w", err)
}
