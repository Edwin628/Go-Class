package models

import (
	"errors"
	"sync"
	"time"
)

// Room represents a video chat room.
type Room struct {
	Name         string    `json:"name"`
	Host         string    `json:"host"`
	Participants []string  `json:"participants"`
	CreatedAt    time.Time `json:"created_at"`
}

var (
	rooms = make(map[string]*Room) // In-memory storage for rooms
	mu    sync.Mutex               // Mutex to handle concurrent access
)

// CreateRoom creates a new room and adds it to the rooms map.
func CreateRoom(name, host string) (*Room, error) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := rooms[name]; exists {
		return nil, errors.New("room already exists")
	}

	room := &Room{
		Name:         name,
		Host:         host,
		Participants: []string{host},
		CreatedAt:    time.Now(),
	}

	rooms[name] = room
	return room, nil
}

// GetRoom retrieves a room by name.
func GetRoom(name string) (*Room, error) {
	mu.Lock()
	defer mu.Unlock()

	room, exists := rooms[name]
	if !exists {
		return nil, errors.New("room not found")
	}

	return room, nil
}

// GetAllRooms returns a list of all rooms.
func GetAllRooms() []*Room {
	mu.Lock()
	defer mu.Unlock()

	roomList := make([]*Room, 0, len(rooms))
	for _, room := range rooms {
		roomList = append(roomList, room)
	}

	return roomList
}

// DeleteRoom deletes a room by name.
func DeleteRoom(name string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := rooms[name]; !exists {
		return errors.New("room not found")
	}

	delete(rooms, name)
	return nil
}
