package storage

import (
	"encoding/json"
	"os"
	"sync"
)

// State represents persisted automation state
type State struct {
	SentConnections map[string]bool `json:"sent_connections"`
	SentMessages    map[string]bool `json:"sent_messages"`
}

// Store handles loading and saving state
type Store struct {
	File  string
	State State
	mu    sync.Mutex
}

// NewStore loads or creates state storage
func NewStore(file string) *Store {
	s := &Store{
		File: file,
		State: State{
			SentConnections: make(map[string]bool),
			SentMessages:    make(map[string]bool),
		},
	}
	s.load()
	return s
}

func (s *Store) load() {
	data, err := os.ReadFile(s.File)
	if err != nil {
		return
	}
	_ = json.Unmarshal(data, &s.State)
}

func (s *Store) save() {
	data, _ := json.MarshalIndent(s.State, "", "  ")
	_ = os.WriteFile(s.File, data, 0644)
}

// MarkConnectionSent persists a sent connection
func (s *Store) MarkConnectionSent(url string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.State.SentConnections[url] = true
	s.save()
}

// MarkMessageSent persists a sent message
func (s *Store) MarkMessageSent(url string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.State.SentMessages[url] = true
	s.save()
}
