package handler

import (
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/service"
	mock_service "userSegmentation/internal/service/mocks"
	"userSegmentation/internal/utils"
)

func TestHandler_createSegment(t *testing.T) {

	type mockBehavior func(r *mock_service.MockSegment, ctx context.Context, segment entity.Segment)
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
			mockBehavior: func(r *mock_service.MockSegment, ctx context.Context, segment entity.Segment) {
				r.EXPECT().CreateSegment(ctx, segment).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}` + "\n",
		},
		{
			name:                 "incorrect segment data",
			inputBody:            `fdnjvl349h`,
			mockBehavior:         func(r *mock_service.MockSegment, ctx context.Context, segment entity.Segment) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"incorrect segment data: bad request"}` + "\n",
		},
	}
	utils.CreateLogger()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockSegment(c)
			tt.mockBehavior(repo, context.Background(), tt.inputSegment)

			services := &service.Service{Segment: repo}
			handler := New(services)

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
	type mockBehavior func(r *mock_service.MockSegment, ctx context.Context, segment entity.Segment)
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
			mockBehavior: func(r *mock_service.MockSegment, ctx context.Context, segment entity.Segment) {
				r.EXPECT().DeleteSegment(ctx, segment.Name).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"success"}` + "\n",
		},
		{
			name:                 "incorrect segment data",
			inputBody:            `fdnjvl349h`,
			mockBehavior:         func(r *mock_service.MockSegment, ctx context.Context, segment entity.Segment) {},
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
			mockBehavior: func(r *mock_service.MockSegment, ctx context.Context, segment entity.Segment) {
				r.EXPECT().DeleteSegment(ctx, segment.Name).Return(errors.Wrap(utils.ErrNotFound,
					"Segment test-segment-name not found"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"Segment test-segment-name not found: not found"}` + "\n",
		},
	}
	utils.CreateLogger()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockSegment(c)
			tt.mockBehavior(repo, context.Background(), tt.inputSegment)

			services := &service.Service{Segment: repo}
			handler := New(services)

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
