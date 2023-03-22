package handlers

import (
	"fmt"
	"github.com/bookmarks-api/pkg/services"
	mock_services "github.com/bookmarks-api/pkg/services/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehavior func(s *mock_services.MockAuthorization, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer Token",
			token:       "Token",
			mockBehavior: func(s *mock_services.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return("1", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "1",
		},
		{
			name:                 "No Header",
			headerName:           "",
			token:                "Token",
			mockBehavior:         func(s *mock_services.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"get authorization token from header"}`,
		},
		{
			name:                 "Invalid Bearer",
			headerName:           "Authorization",
			headerValue:          "Bearr Token",
			token:                "",
			mockBehavior:         func(s *mock_services.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"auth token in wrong format"}`,
		},
		{
			name:                 "Invalid Token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "",
			mockBehavior:         func(s *mock_services.MockAuthorization, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"auth token in wrong format"}`,
		},
		{
			name:        "Service Failure",
			headerName:  "Authorization",
			headerValue: "Bearer Token",
			token:       "Token",
			mockBehavior: func(s *mock_services.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return("1", errors.New("failed to parse token"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"failed to parse token"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// init deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_services.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.token)

			service := &services.Service{Authorization: auth}
			handler := NewHandler(service)

			// test server
			r := gin.New()
			r.GET("/protected", handler.UserIdentity, func(c *gin.Context) {
				id, _ := c.Get("userId")
				c.String(200, fmt.Sprintf("%d", id.(int)))
			})

			// test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			// make request
			r.ServeHTTP(w, req)

			// assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}
