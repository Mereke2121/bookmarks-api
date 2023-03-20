package handlers

import (
	"bytes"
	"github.com/bookmarks-api/models"
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

func TestHandler_SignUp(t *testing.T) {
	type mockBehavior func(s *mock_services.MockAuthorization, user models.User)

	testTable := []struct {
		name                      string
		inputBody                 string
		inputUser                 models.User
		mockBehavior              mockBehavior
		expectedStatusCode        int
		expectedStatusRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"user_name": "Test", "email": "test@example.com", "password": "qwerty"}`,
			inputUser: models.User{
				UserName: "Test",
				Email:    "test@example.com",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_services.MockAuthorization, user models.User) {
				s.EXPECT().AddUser(&user).Return(1, nil)
			},
			expectedStatusCode:        200,
			expectedStatusRequestBody: `{"id":1}`,
		},
		{
			name:                      "Empty Fields",
			inputBody:                 `{"user_name": "test", "password": "qwerty"}`,
			mockBehavior:              func(s *mock_services.MockAuthorization, user models.User) {},
			expectedStatusCode:        400,
			expectedStatusRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Service Failure",
			inputBody: `{"user_name": "Test", "email": "test@example.com", "password": "qwerty"}`,
			inputUser: models.User{
				UserName: "Test",
				Email:    "test@example.com",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_services.MockAuthorization, user models.User) {
				s.EXPECT().AddUser(&user).Return(0, errors.New("service failure"))
			},
			expectedStatusCode:        500,
			expectedStatusRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// init deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_services.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			service := &services.Service{Authorization: auth}
			handler := NewHandler(service)

			// test server
			r := gin.New()
			r.POST("/sign-up", handler.SignUp)

			// test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/sign-up", bytes.NewBufferString(testCase.inputBody))

			// perform request
			r.ServeHTTP(w, req)

			// assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedStatusRequestBody, w.Body.String())
		})
	}
}
