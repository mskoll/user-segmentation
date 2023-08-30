package handler

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/lib/errTypes"
	"userSegmentation/internal/lib/logger"
	"userSegmentation/internal/service"
	mock_service "userSegmentation/internal/service/mocks"
)

func TestHandler_createSegment(t *testing.T) {

	type mockBehavior func(r *mock_service.MockSegment, segment entity.Segment)
	tests := []struct {
		name                 string
		inputBody            string
		inputSegment         entity.Segment
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: `{"name": "test-segment-name", "percent": 0}`,
			inputSegment: entity.Segment{
				Name:    "test-segment-name",
				Percent: 0,
			},
			mockBehavior: func(r *mock_service.MockSegment, segment entity.Segment) {
				r.EXPECT().CreateSegment(segment).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}` + "\n",
		},
		{
			name:                 "incorrect segment data",
			inputBody:            `fdnjvl349h`,
			mockBehavior:         func(r *mock_service.MockSegment, segment entity.Segment) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"incorrect segment data: bad request"}` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockSegment(c)
			tt.mockBehavior(repo, tt.inputSegment)

			services := &service.Service{Segment: repo}
			handler := New(services, logger.CreateLogger())

			r := echo.New()
			r.POST("/segment/", handler.createSegment)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/segment/", bytes.NewBufferString(tt.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())

		})
	}
}

func TestHandler_deleteSegment(t *testing.T) {
	type mockBehavior func(r *mock_service.MockSegment, segment entity.Segment)
	tests := []struct {
		name                 string
		inputBody            string
		inputSegment         entity.Segment
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: `{"name": "test-segment-name"}`,
			inputSegment: entity.Segment{
				Name:    "test-segment-name",
				Percent: 0,
			},
			mockBehavior: func(r *mock_service.MockSegment, segment entity.Segment) {
				r.EXPECT().DeleteSegment(segment.Name).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"success"}` + "\n",
		},
		{
			name:                 "incorrect segment data",
			inputBody:            `fdnjvl349h`,
			mockBehavior:         func(r *mock_service.MockSegment, segment entity.Segment) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"incorrect segment data: bad request"}` + "\n",
		},
		{
			name:      "segment not found ",
			inputBody: `{"name": "test-segment-name"}`,
			inputSegment: entity.Segment{
				Name:    "test-segment-name",
				Percent: 0,
			},
			mockBehavior: func(r *mock_service.MockSegment, segment entity.Segment) {
				r.EXPECT().DeleteSegment(segment.Name).Return(errors.Wrap(errTypes.ErrNotFound,
					"Segment test-segment-name not found"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"Segment test-segment-name not found: not found"}` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockSegment(c)
			tt.mockBehavior(repo, tt.inputSegment)

			services := &service.Service{Segment: repo}
			handler := New(services, logger.CreateLogger())

			r := echo.New()
			r.DELETE("/segment/", handler.deleteSegment)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/segment/", bytes.NewBufferString(tt.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())

		})
	}
}
