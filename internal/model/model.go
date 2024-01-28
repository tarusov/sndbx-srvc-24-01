package model

import (
	"fmt"
	"time"
)

type (
	Date time.Time

	Event struct {
		ID          int    `json:"id"`
		UserID      int    `json:"user_id"`
		Date        Date   `json:"date"`
		Description string `json:"description"`
	}

	EventCreateRequest struct {
		UserID      int
		Date        Date
		Description string
	}

	EventUpdateRequest struct {
		ID          int
		UserID      int
		Date        Date
		Description string
	}

	EventDeleteRequest struct {
		ID int
	}

	EventFilter struct {
		ID     int
		UserID int
		From   Date
		To     Date
	}
)

// MarshalJSON send date as 2006-01-02
func (d Date) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`"%s"`, time.Time(d).Format("2006-01-02"))
	return []byte(s), nil
}

func (ef EventFilter) Match(e Event) bool {

	if ef.ID != 0 {
		if ef.ID != e.ID {
			return false
		}
	}

	if ef.UserID != 0 {
		if ef.UserID != e.UserID {
			return false
		}
	}

	f := time.Time(ef.From)
	if !f.IsZero() {
		if f.After(time.Time(e.Date)) {
			return false
		}
	}

	t := time.Time(ef.To)
	if !t.IsZero() {
		if t.Before(time.Time(e.Date)) {
			return false
		}
	}

	return true
}
