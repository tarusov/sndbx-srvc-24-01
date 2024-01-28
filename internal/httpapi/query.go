package httpapi

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/tarusov/sndbx-srvc-24-01/internal/model"
)

const (
	queryParamEventID     = "id"
	queryParamUserID      = "user_id"
	queryParamDate        = "date"
	queryParamDescription = "description"

	queryDateFormat = "2006-01-02"
)

type (
	eventParams struct {
		UserID      int
		Date        model.Date
		Description string
	}
)

// parseEventCreateRequest
func parseEventCreateRequest(vals url.Values) (empty model.EventCreateRequest, err error) {

	ep, err := parseEventParams(vals)
	if err != nil {
		return empty, err
	}

	return model.EventCreateRequest{
		UserID:      ep.UserID,
		Date:        ep.Date,
		Description: ep.Description,
	}, nil
}

// parseEventUpdateRequest
func parseEventUpdateRequest(vals url.Values) (empty model.EventUpdateRequest, err error) {

	id, err := parseEventID(vals)
	if err != nil {
		return empty, err
	}

	ep, err := parseEventParams(vals)
	if err != nil {
		return empty, err
	}

	return model.EventUpdateRequest{
		ID:          id,
		UserID:      ep.UserID,
		Date:        ep.Date,
		Description: ep.Description,
	}, nil
}

// parseEventDeleteRequest
func parseEventDeleteRequest(vals url.Values) (empty model.EventDeleteRequest, err error) {

	id, err := parseEventID(vals)
	if err != nil {
		return empty, err
	}

	return model.EventDeleteRequest{
		ID: id,
	}, nil
}

// parseEventID
func parseEventID(vals url.Values) (int, error) {

	eventID, err := strconv.Atoi(vals.Get(queryParamEventID))
	if err != nil {
		return 0, fmt.Errorf("invalid id format: %w", err)
	}

	return eventID, nil
}

// parseEventParams
func parseEventParams(vals url.Values) (empty eventParams, err error) {

	userID, err := strconv.Atoi(vals.Get(queryParamUserID))
	if err != nil {
		return empty, fmt.Errorf("invalid user_id format: %w", err)
	}

	date, err := time.Parse(queryDateFormat, vals.Get(queryParamDate))
	if err != nil {
		return empty, fmt.Errorf("invalid date format: %w", err)
	}

	return eventParams{
		UserID:      userID,
		Date:        model.Date(date),
		Description: vals.Get(queryParamDescription),
	}, nil
}

// parseDate
func parseDate(vals url.Values) (model.Date, error) {

	date, err := time.Parse(queryDateFormat, vals.Get(queryParamDate))
	if err != nil {
		return model.Date{}, fmt.Errorf("invalid date format: %w", err)
	}

	return model.Date(date), nil
}
