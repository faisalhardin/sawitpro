package handler

import (
	"encoding/json"
	"net/http"

	estateUsecase "github.com/faisalhardin/sawitpro/internal/entity/interfaces"
	model "github.com/faisalhardin/sawitpro/internal/entity/model"
	"github.com/pkg/errors"
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
	err := bind(r, &request)
	if err != nil {
		setError(r, w, err)
		return
	}

	resp, err := h.EstateUsecase.InsertEstate(ctx, request)
	if err != nil {
		setError(r, w, err)
		return
	}

	setOKWithData(r, w, resp)
}

func bind(r *http.Request, targetDecode interface{}) error {
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
