package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"canteen-app/internal/adapter/http/api/mocks"
	domAuth "canteen-app/internal/domain/auth"
	"canteen-app/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupRouterWithAuthUseCase(authUC *mocks.AuthUseCase, refreshTTL time.Duration) *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	NewAuthHandler(r, authUC, refreshTTL)

	return r
}

func TestAuthHandler_Register(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    map[string]string
		setupAuthUC    func(m *mocks.AuthUseCase)
		wantStatusCode int
		wantErrorText  string
	}{
		{
			name: "success",
			requestBody: map[string]string{
				"login":    "the_real_slim_shady",
				"password": "sdf3kJIS2FgiwefiJCiSJ5#@KJFKj",
				"name":     "Slim",
				"surname":  "Shady",
				"role":     "admin",
			},

			setupAuthUC: func(m *mocks.AuthUseCase) {
				m.On("Register", "the_real_slim_shady", "sdf3kJIS2FgiwefiJCiSJ5#@KJFKj", "Slim", "Shady", "admin").Return(
					&domAuth.Tokens{
						AccessToken:  "access_token",
						RefreshToken: "refresh_token",
					}, nil).Once()
			},

			wantStatusCode: http.StatusCreated,
			wantErrorText:  "",
		},

		{
			name: "missing required field",
			requestBody: map[string]string{
				"login":   "the_real_slim_shady",
				"name":    "Slim",
				"surname": "Shady",
				"role":    "admin",
			},

			wantStatusCode: http.StatusBadRequest,
			wantErrorText:  "invalid request",
		},

		{
			name: "user exists error",
			requestBody: map[string]string{
				"login":    "the_real_slim_shady",
				"password": "sdf3kJIS2FgiwefiJCiSJ5#@KJFKj",
				"name":     "Slim",
				"surname":  "Shady",
				"role":     "admin",
			},

			setupAuthUC: func(m *mocks.AuthUseCase) {
				m.On("Register", "the_real_slim_shady", "sdf3kJIS2FgiwefiJCiSJ5#@KJFKj", "Slim", "Shady", "admin").Return(
					func(login, password, name, surname, role string) (*domAuth.Tokens, error) {
						return &domAuth.Tokens{}, usecase.ErrUserExists
					},
				).Once()
			},

			wantStatusCode: http.StatusConflict,
			wantErrorText:  "user already exists",
		},

		{
			name: "internal server error",
			requestBody: map[string]string{
				"login":    "the_real_slim_shady",
				"password": "sdf3kJIS2FgiwefiJCiSJ5#@KJFKj",
				"name":     "Slim",
				"surname":  "Shady",
				"role":     "admin",
			},

			setupAuthUC: func(m *mocks.AuthUseCase) {
				m.On("Register", "the_real_slim_shady", "sdf3kJIS2FgiwefiJCiSJ5#@KJFKj", "Slim", "Shady", "admin").Return(
					func(login, password, name, surname, role string) (*domAuth.Tokens, error) {
						return &domAuth.Tokens{}, errors.New("error")
					},
				).Once()
			},

			wantStatusCode: http.StatusInternalServerError,
			wantErrorText:  "internal server error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			authUC := mocks.NewAuthUseCase(t)

			if tc.setupAuthUC != nil {
				tc.setupAuthUC(authUC)
			}

			router := setupRouterWithAuthUseCase(authUC, time.Duration(30))

			bodyBytes, err := json.Marshal(tc.requestBody)
			req, err := http.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewReader(bodyBytes))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.wantStatusCode, w.Code)

			var resp map[string]interface{}
			err = json.Unmarshal(w.Body.Bytes(), &resp)
			require.NoError(t, err)

			if tc.wantErrorText != "" {
				assert.Equal(t, tc.wantErrorText, resp["error"])
			} else {
				assert.Equal(t, "access_token", resp["access_token"])

				cookies := w.Result().Cookies()
				require.NotEmpty(t, cookies)

				cookie := cookies[0]

				assert.Equal(t, false, cookie.Secure)
				assert.Equal(t, true, cookie.HttpOnly)
				assert.Equal(t, "/", cookie.Path)
				assert.Equal(t, "", cookie.Domain)
				assert.Equal(t, http.SameSiteLaxMode, cookie.SameSite)
				assert.NotEmpty(t, cookie.Expires)
				assert.Equal(t, "refresh_token", cookie.Name)
				assert.Equal(t, "refresh_token", cookie.Value)
			}

			authUC.AssertExpectations(t)
		})
	}
}
