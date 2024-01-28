package service

import (
	"context"

	"github.com/tarusov/sndbx-srvc-24-01/internal/model"
)

type (
	Service struct {
		storage eventStorage
	}

	eventStorage interface {
		CreateEvent(context.Context, model.EventCreateRequest) (int, error)
		UpdateEvent(context.Context, model.EventUpdateRequest) error
		DeleteEvent(context.Context, model.EventDeleteRequest) error
		GetEvents(context.Context, model.EventFilter) ([]model.Event, error)
	}
)

// New
func New(es eventStorage) *Service {
	return &Service{
		storage: es,
	}
}

func (s *Service) CreateEvent(ctx context.Context, r model.EventCreateRequest) (model.Event, error) {

	id, err := s.storage.CreateEvent(ctx, r)
	if err != nil {
		return model.Event{}, err
	}

	return model.Event{
		ID:          id,
		UserID:      r.UserID,
		Date:        r.Date,
		Description: r.Description,
	}, nil
}

func (s *Service) UpdateEvent(ctx context.Context, r model.EventUpdateRequest) error {
	return s.storage.UpdateEvent(ctx, r)
}

func (s *Service) DeleteEvent(ctx context.Context, r model.EventDeleteRequest) error {
	return s.storage.DeleteEvent(ctx, r)
}

func (s *Service) GetEvents(ctx context.Context, ef model.EventFilter) ([]model.Event, error) {
	return s.storage.GetEvents(ctx, ef)
}
