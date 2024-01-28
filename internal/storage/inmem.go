package storage

import (
	"context"
	"errors"

	"github.com/tarusov/sndbx-srvc-24-01/internal/model"
)

var ErrNotFound = errors.New("not found")

type (
	InMemStorage struct {
		data map[int]model.Event
	}
)

func NewInMemStorage() *InMemStorage {
	return &InMemStorage{
		data: make(map[int]model.Event),
	}
}

// CreateEvent
func (s *InMemStorage) CreateEvent(_ context.Context, r model.EventCreateRequest) (int, error) {

	id := len(s.data) + 1 // TODO: skip 0 counter for filter after.
	s.data[id] = model.Event{
		ID:          id,
		UserID:      r.UserID,
		Date:        r.Date,
		Description: r.Description,
	}

	return id, nil
}

// UpdateEvent
func (s *InMemStorage) UpdateEvent(_ context.Context, r model.EventUpdateRequest) error {

	if _, ok := s.data[r.ID]; !ok {
		return ErrNotFound
	}

	s.data[r.ID] = model.Event(r)

	return nil
}

// DeleteEvent
func (s *InMemStorage) DeleteEvent(_ context.Context, r model.EventDeleteRequest) error {

	if _, ok := s.data[r.ID]; !ok {
		return ErrNotFound
	}

	delete(s.data, r.ID)

	return nil
}

// GetEvents
func (s *InMemStorage) GetEvents(_ context.Context, ef model.EventFilter) ([]model.Event, error) {

	result := make([]model.Event, 0)
	for _, v := range s.data {
		if ef.Match(v) {
			result = append(result, v)
		}
	}

	return result, nil
}
