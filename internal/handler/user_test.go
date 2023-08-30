package handler

import (
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"time"
	"userSegmentation/internal/entity"
	"userSegmentation/internal/service"
	mock_service "userSegmentation/internal/service/mocks"
	"userSegmentation/internal/utils"
)

func TestHandler_createUser(t *testing.T) {

	type mockBehavior func(r *mock_service.MockUser, ctx context.Context, user entity.User)
	tests := []struct {
		name                 string
		inputBody            string
		inputUser            entity.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: `{"username": "test-username"}`,
			inputUser: entity.User{
				Username: "test-username",
			},
			mockBehavior: func(r *mock_service.MockUser, ctx context.Context, user entity.User) {
				r.EXPECT().CreateUser(ctx, user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}` + "\n",
		},
		{
			name:                 "incorrect user data",
			inputBody:            `fdnkgj4387gt`,
			mockBehavior:         func(r *mock_service.MockUser, ctx context.Context, user entity.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"incorrect user data: bad request"}` + "\n",
		},
	}
	utils.CreateLogger()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockUser(c)
			tt.mockBehavior(repo, context.Background(), tt.inputUser)

			services := &service.Service{User: repo}
			handler := New(services)

			r := echo.New()
			r.POST("/user/", handler.createUser)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/user/", bytes.NewBufferString(tt.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
		})
	}
}

func TestHandler_operations(t *testing.T) {

	type mockBehavior func(r *mock_service.MockUser, ctx context.Context, userOperations entity.UserOperations)
	tests := []struct {
		name                 string
		inputBody            string
		inputUser            entity.UserOperations
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: `{"user_id": 1, "month":8, "year":2023}`,
			inputUser: entity.UserOperations{
				UserId: 1,
				Month:  8,
				Year:   2023,
			},
			mockBehavior: func(r *mock_service.MockUser, ctx context.Context, userOperations entity.UserOperations) {
				r.EXPECT().Operations(ctx, userOperations).Return([]entity.Operation{
					{
						UserId: 1, SegmentName: "test-segment-name", Operation: "created",
						Datetime: time.Date(2023, time.August, 8, 12, 0, 0, 0, time.UTC),
					},
					{
						UserId: 1, SegmentName: "test-segment-name", Operation: "deleted",
						Datetime: time.Date(2023, time.August, 20, 12, 0, 0, 0, time.UTC),
					},
				}, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: `[{"user_id":1,"segment_name":"test-segment-name","operation":"created","datetime":"2023-08-08T12:00:00Z"},` +
				`{"user_id":1,"segment_name":"test-segment-name","operation":"deleted","datetime":"2023-08-20T12:00:00Z"}]` + "\n",
		},
		{
			name:                 "incorrect input",
			inputBody:            `ggvk`,
			mockBehavior:         func(r *mock_service.MockUser, ctx context.Context, userOperations entity.UserOperations) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"incorrect input data: bad request"}` + "\n",
		},
	}
	utils.CreateLogger()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockUser(c)
			tt.mockBehavior(repo, context.Background(), tt.inputUser)

			services := &service.Service{User: repo}
			handler := New(services)

			r := echo.New()
			r.POST("/user/operations", handler.operations)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/user/operations", bytes.NewBufferString(tt.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
		})
	}
}
