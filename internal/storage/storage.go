package storage

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

// State represents persisted automation state
type State struct {
	SentProfiles []string  `json:"sent_profiles"`
	MessagesSent []string  `json:"messages_sent"`
	LastRun      time.Time `json:"last_run"`
}

var (
	stateFile = "state.json"
	mutex     sync.Mutex
)

// LoadState loads saved state from disk
func LoadState() (*State, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// If file does not exist, return empty state
	if _, err := os.Stat(stateFile); os.IsNotExist(err) {
		return &State{
			SentProfiles: []string{},
			MessagesSent: []string{},
			LastRun:      time.Now(),
		}, nil
	}

	file, err := os.Open(stateFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var state State
	err = json.NewDecoder(file).Decode(&state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

// SaveState saves current state to disk
func SaveState(state *State) error {
	mutex.Lock()
	defer mutex.Unlock()

	state.LastRun = time.Now()

	file, err := os.Create(stateFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(state)
}
