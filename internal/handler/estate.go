package handler

import (
	"encoding/json"
	"net/http"

	estateUsecase "github.com/faisalhardin/sawitpro/internal/entity/interfaces"
	model "github.com/faisalhardin/sawitpro/internal/entity/model"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
)

var (
	setErrorFunc      = setError
	setOKWithDataFunc = setOKWithData
	bindFunc          = bind
)

type EstateHandler struct {
	EstateUsecase estateUsecase.EstateUsecase
}

func NewEstateHandler(h *EstateHandler) *EstateHandler {
	return h
}

func (h *EstateHandler) InsertEstate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	request := model.InsertEstateRequest{}
	err := bindFunc(r, &request)
	if err != nil {
		setErrorFunc(r, w, err)
		return
	}

	resp, err := h.EstateUsecase.InsertEstate(ctx, request)
	if err != nil {
		setErrorFunc(r, w, err)
		return
	}

	setOKWithDataFunc(r, w, resp)
}

func (h *EstateHandler) GetEstateStats(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	estateID := chi.URLParam(r, "uuid")
	resp, err := h.EstateUsecase.GetEstateStatsByUUID(ctx, estateID)
	if err != nil {
		setErrorFunc(r, w, err)
		return
	}

	setOKWithDataFunc(r, w, resp)
}

func (h *EstateHandler) GetDronePlan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	estateID := chi.URLParam(r, "uuid")
	params := model.GetDronePlanParams{}
	err := decodeSchemaRequest(r, &params)
	if err != nil {
		return
	}

	params.UUID = estateID

	resp, err := h.EstateUsecase.GetDronePlanByEstateUUID(ctx, params)
	if err != nil {
		setErrorFunc(r, w, err)
		return
	}

	setOKWithDataFunc(r, w, resp)
}

func decodeSchemaRequest(r *http.Request, val interface{}) error {
	sourceDecode := r.URL.Query()
	decoder := schema.NewDecoder()
	if err := decoder.Decode(val, sourceDecode); err != nil {
		return err
	}
	// return BindQuery(sourceDecode, val, ignoreStatus...)
	return nil
}

func bind(r *http.Request, targetDecode interface{}) error {
	if r.Method == http.MethodGet {
		if err := decodeSchemaRequest(r, targetDecode); err != nil {
			return err
		}
		return nil
	}
	bodyDecode := json.NewDecoder(r.Body)
	err := bodyDecode.Decode(targetDecode)
	if err != nil {
		return errors.Wrap(err, "bind")
	}

	return err
}

func setError(r *http.Request, w http.ResponseWriter, errInput error) (err error) {

	switch errCause := errors.Cause(errInput).(type) {
	case *model.Response:
		writeErrJSON(w, errCause.Code, errCause)
	default:
		writeErrJSON(w, http.StatusNotFound, model.Response{
			Data:    "Not Found",
			ErrName: "404",
		})
	}
	return err
}

func writeErrJSON(w http.ResponseWriter, status int, data interface{}) (int, error) {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeLen, writeErr := w.Write([]byte(`{"errors":["Internal Server Error"]}`))
		if writeErr != nil {
			return writeLen, writeErr
		}
		return writeLen, err
	}

	w.WriteHeader(status)
	return w.Write(b)
}

func setOKWithData(r *http.Request, w http.ResponseWriter, data interface{}) (err error) {
	_, err = write(w, r, http.StatusOK, data)
	return err
}
func write(w http.ResponseWriter, r *http.Request, status int, data interface{}) (int, error) {

	w.Header().Set("Content-Type", "application/json")
	responseInBytes, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		wLen, wErr := w.Write([]byte(`{"errors:["Internal Server Error"]}`))
		if wErr != nil {
			return wLen, wErr
		}
	}

	w.WriteHeader(status)
	return w.Write(responseInBytes)
}
