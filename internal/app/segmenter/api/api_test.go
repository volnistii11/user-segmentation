package api

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/volnistii11/user-segmentation/internal/app/segmenter/repository/database"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	return gin.Default()
}

func TestAPI_CreateSegment(t *testing.T) {
	type want struct {
		code int
	}
	type request struct {
		body        []byte
		contentType string
	}
	tests := []struct {
		name        string
		request     string
		body        request
		contentType request
		want        want
	}{
		{
			name:    "wrong content-type",
			request: "/api/segment/create",
			body: request{
				body:        []byte(`{"slug": "AVITO_5"}`),
				contentType: "text/plain",
			},
			want: want{
				code: 400,
			},
		},
		{
			name:    "body is empty",
			request: "/api/segment/create",
			body: request{
				body:        []byte(``),
				contentType: "application/json",
			},
			want: want{
				code: 400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, _ := zap.NewDevelopment()
			repo := database.New(nil)
			router := SetUpRouter()
			handlers := New(logger, repo)

			var body bytes.Buffer
			_ = json.NewEncoder(&body).Encode(tt.body)

			router.POST("/api/segment/create", handlers.CreateSegment)
			req, _ := http.NewRequest(http.MethodPost, tt.request, &body)
			req.Header.Set("content-type", tt.contentType.contentType)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.want.code, w.Code)
		})
	}
}
