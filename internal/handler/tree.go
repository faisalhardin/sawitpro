package handler

import (
	"net/http"

	model "github.com/faisalhardin/sawitpro/internal/entity/model"
	"github.com/go-chi/chi/v5"
)

func (h *EstateHandler) InsertTree(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	request := model.InsertNewTreeRequest{}
	err := bindFunc(r, &request)
	if err != nil {
		setError(r, w, err)
		return
	}

	request.EstateUUID = chi.URLParam(r, "uuid")

	resp, err := h.EstateUsecase.InsertNewTree(ctx, request)
	if err != nil {
		setErrorFunc(r, w, err)
		return
	}

	setOKWithDataFunc(r, w, resp)
}
