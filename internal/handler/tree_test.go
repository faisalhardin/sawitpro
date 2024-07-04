package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	model "github.com/faisalhardin/sawitpro/internal/entity/model"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInsertTreeHandler_Success(t *testing.T) {
	ctrl := initMock(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		patch      func()
		statusCode int
	}{
		{
			name: "Success",
			patch: func() {
				bindFunc = func(r *http.Request, targetDecode interface{}) error {
					return nil
				}

				mockUsecase.EXPECT().InsertNewTree(gomock.Any(), gomock.Any()).
					Return(model.InsertNewTreeResponse{}, nil).Times(1)

				setOKWithDataFunc = func(r *http.Request, w http.ResponseWriter, data interface{}) (err error) {
					return nil
				}

			},
			statusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/estate/{uuid}/tree", http.NoBody)
			chi.NewRouteContext().URLParams.Add("uuid", "some-uuid")

			tt.patch()
			h := EstateHandler{
				EstateUsecase: mockUsecase,
			}
			h.InsertTree(w, r)
			assert.Equal(t, tt.statusCode, w.Code)
		})
	}
}
