package httpapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/tarusov/sndbx-srvc-24-01/internal/model"
)

type (
	Hander struct {
		service calendarService
	}

	calendarService interface {
		CreateEvent(ctx context.Context, r model.EventCreateRequest) (model.Event, error)
		UpdateEvent(ctx context.Context, r model.EventUpdateRequest) error
		DeleteEvent(ctx context.Context, r model.EventDeleteRequest) error
		GetEvents(ctx context.Context, ef model.EventFilter) ([]model.Event, error)
	}

	resultResp struct {
		Result any `json:"result"`
	}

	errorResp struct {
		Error any `json:"error"`
	}
)

// CTOR
func NewHandler(cs calendarService) *Hander {
	return &Hander{
		service: cs,
	}
}

// CreateEvent
func (h *Hander) CreateEvent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeMethodNotAllowedError(w)
		return
	}

	if err := r.ParseForm(); err != nil {
		writeBadRequestError(w, err)
		return
	}

	reqData, err := parseEventCreateRequest(r.Form)
	if err != nil {
		writeBadRequestError(w, err)
		return
	}

	event, err := h.service.CreateEvent(r.Context(), reqData)
	if err != nil {
		writeInternalServerError(w, err)
		return
	}

	writeResult(w, event)
}

// UpdateEvent
func (h *Hander) UpdateEvent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeMethodNotAllowedError(w)
		return
	}

	if err := r.ParseForm(); err != nil {
		writeBadRequestError(w, err)
		return
	}

	reqData, err := parseEventUpdateRequest(r.Form)
	if err != nil {
		writeBadRequestError(w, err)
		return
	}

	err = h.service.UpdateEvent(r.Context(), reqData)
	if err != nil {
		writeInternalServerError(w, err)
		return
	}

	writeResult(w, "Success")
}

// DeleteEvent
func (h *Hander) DeleteEvent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeMethodNotAllowedError(w)
		return
	}

	if err := r.ParseForm(); err != nil {
		writeBadRequestError(w, err)
		return
	}

	reqData, err := parseEventDeleteRequest(r.Form)
	if err != nil {
		writeBadRequestError(w, err)
		return
	}

	err = h.service.DeleteEvent(r.Context(), reqData)
	if err != nil {
		writeInternalServerError(w, err)
		return
	}

	writeResult(w, "Success")
}

// EventsForDay
func (h *Hander) EventsForDay(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeMethodNotAllowedError(w)
		return
	}

	if err := r.ParseForm(); err != nil {
		writeBadRequestError(w, err)
		return
	}

	date, err := parseDate(r.Form)
	if err != nil {
		writeBadRequestError(w, err)
		return
	}

	result, err := h.service.GetEvents(r.Context(), model.EventFilter{
		From: date,
		To:   date,
	})
	if err != nil {
		writeInternalServerError(w, err)
		return
	}

	writeResult(w, result)
}

// EventsForMonth
func (h *Hander) EventsForMonth(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeMethodNotAllowedError(w)
		return
	}

}

// EventsForWeek
func (h *Hander) EventsForWeek(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeMethodNotAllowedError(w)
		return
	}

}

// writeResult
func writeResult(w http.ResponseWriter, result any) {
	data, _ := json.Marshal(&resultResp{
		Result: result,
	})
	_, _ = w.Write(data)
}

// writeMethodNotAllowedError
func writeMethodNotAllowedError(w http.ResponseWriter) {
	data, _ := json.Marshal(&errorResp{
		Error: http.StatusText(http.StatusMethodNotAllowed),
	})
	writeError(w, data, http.StatusMethodNotAllowed)
}

// writeBadRequestError
func writeBadRequestError(w http.ResponseWriter, err error) {
	data, _ := json.Marshal(&errorResp{
		Error: err.Error(),
	})
	writeError(w, data, http.StatusBadRequest)
}

// writeInternalServerError
func writeInternalServerError(w http.ResponseWriter, err error) {
	data, _ := json.Marshal(&errorResp{
		Error: err.Error(),
	})
	writeError(w, data, http.StatusInternalServerError)
}

// writeError
func writeError(w http.ResponseWriter, body []byte, statusCode int) {
	http.Error(w, string(body), statusCode)
}
