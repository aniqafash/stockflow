package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"stockflow/internal/dto"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestProductHandler_List(t *testing.T) {
	tests := []struct {
		name            string
		request         string
		search          string
		mockError       error
		expectedSvcCall int
		wantStatus      int
		wantCode        string
	}{
		{
			name:            "valid request",
			request:         "/products",
			search:          "",
			mockError:       nil,
			expectedSvcCall: 1,
			wantStatus:      http.StatusOK,
		},
		{
			name:            "valid search request",
			request:         "/products?search=Flatshoes",
			search:          "Flatshoes",
			mockError:       nil,
			expectedSvcCall: 1,
			wantStatus:      http.StatusOK,
		},
		{
			name:            "valid page number request",
			request:         "/products?page=2",
			search:          "",
			mockError:       nil,
			expectedSvcCall: 1,
			wantStatus:      http.StatusOK,
		},
		{
			name:            "invalid page number request",
			request:         "/products?page=abc",
			search:          "",
			mockError:       errors.New(""),
			expectedSvcCall: 0,
			wantStatus:      http.StatusBadRequest,
			wantCode:        "BAD_PRODUCTS_REQUEST",
		},
		{
			name:            "valid page limit request",
			request:         "/products?limit=2",
			search:          "",
			mockError:       nil,
			expectedSvcCall: 1,
			wantStatus:      http.StatusOK,
		},
		{
			name:            "invalid page limit request",
			request:         "/products?limit=abc",
			search:          "",
			mockError:       errors.New(""),
			expectedSvcCall: 0,
			wantStatus:      http.StatusBadRequest,
			wantCode:        "BAD_PRODUCTS_REQUEST",
		},
		{
			name:            "error lookup",
			request:         "/products",
			search:          "",
			mockError:       dto.ErrLookUp,
			expectedSvcCall: 1,
			wantStatus:      http.StatusInternalServerError,
			wantCode:        "ERROR_LOOKUP",
		},
		{
			name:            "other service error",
			request:         "/products",
			search:          "",
			mockError:       errors.New("service error"),
			expectedSvcCall: 1,
			wantStatus:      http.StatusInternalServerError,
			wantCode:        "ERROR_INTERNAL_SERVER",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockService := NewMockProductService(ctrl)

			mockService.EXPECT().
				List(gomock.Any(), tt.search, gomock.Any()).
				Return([]dto.Product{}, tt.mockError).
				Times(tt.expectedSvcCall)

			w := httptest.NewRecorder()
			c, r := gin.CreateTestContext(w)
			h := NewProductHandler(mockService)
			r.GET("/products", h.List)

			req := httptest.NewRequest(http.MethodGet, tt.request, nil)
			c.Request = req
			r.ServeHTTP(w, req)

			require.Equal(t, tt.wantStatus, w.Result().StatusCode)
			if tt.wantStatus != http.StatusOK {
				var resp dto.FailResponse
				err := json.NewDecoder(w.Body).Decode(&resp)
				require.NoError(t, err)
				require.Equal(t, tt.wantCode, resp.Error.Code)
			}
		})
	}
}
