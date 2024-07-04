package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/faisalhardin/sawitpro/internal/entity/mocks"
	model "github.com/faisalhardin/sawitpro/internal/entity/model"
)

var mockUsecase *mocks.MockEstateUsecase

func initMock(t *testing.T) *gomock.Controller {
	ctrl := gomock.NewController(t)
	mockUsecase = mocks.NewMockEstateUsecase(ctrl)

	return ctrl
}

func TestInsertEstateHandler_Success(t *testing.T) {
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

				mockUsecase.EXPECT().InsertEstate(gomock.Any(), gomock.Any()).
					Return(model.InsertEstateResponse{}, nil).Times(1)

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
			r := httptest.NewRequest(http.MethodPost, "/estate", http.NoBody)

			tt.patch()
			h := EstateHandler{
				EstateUsecase: mockUsecase,
			}
			h.InsertEstate(w, r)
			assert.Equal(t, tt.statusCode, w.Code)
		})
	}

}

func Test_GetEstateStatsHandler(t *testing.T) {
	ctrl := initMock(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		url        string
		patch      func()
		statusCode int
	}{
		{
			name: "Success",
			url:  "/UUID/stats",
			patch: func() {
				mockUsecase.EXPECT().
					GetEstateStatsByUUID(gomock.Any(), gomock.Any()).
					Return(model.EstateStats{}, nil).
					Times(1)

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
			r := httptest.NewRequest(http.MethodGet, tt.url, nil)

			tt.patch()
			h := EstateHandler{
				EstateUsecase: mockUsecase,
			}
			h.GetEstateStats(w, r)
			assert.Equal(t, tt.statusCode, w.Code)
		})
	}
}

func Test_GetDronePlanHandler(t *testing.T) {
	ctrl := initMock(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		url        string
		patch      func()
		statusCode int
	}{
		{
			name: "Success",
			url:  "/UUID/drone-plan",
			patch: func() {
				mockUsecase.EXPECT().
					GetDronePlanByEstateUUID(gomock.Any(), gomock.Any()).
					Return(model.EstateDronePlanResponse{}, nil).
					Times(1)

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
			r := httptest.NewRequest(http.MethodGet, tt.url, nil)

			tt.patch()
			h := EstateHandler{
				EstateUsecase: mockUsecase,
			}
			h.GetDronePlan(w, r)
			assert.Equal(t, tt.statusCode, w.Code)
		})
	}
}
