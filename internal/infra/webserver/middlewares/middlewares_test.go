package middlewares

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/matheusmhmelo/FullCycle-rate-limiter/internal/infra/database"
	"github.com/matheusmhmelo/FullCycle-rate-limiter/internal/infra/database/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLimiter_Do(t *testing.T) {
	tests := []struct {
		name       string
		req        func() *http.Request
		assertMock func(T *testing.T) *mock.MockRateInterface
		expectNext bool
	}{
		{
			name: "success with IP",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
				require.NoError(t, err)
				req.RemoteAddr = "127.0.0.1"
				return req
			},
			assertMock: func(t *testing.T) *mock.MockRateInterface {
				ctrl := gomock.NewController(t)
				m := mock.NewMockRateInterface(ctrl)

				m.EXPECT().
					FindBlocker(gomock.Any(), database.TypeIP, "127.0.0.1").
					Return(false, nil)
				m.EXPECT().
					FindRequests(gomock.Any(), database.TypeIP, "127.0.0.1").
					Return(0, nil)
				m.EXPECT().
					NewRequest(gomock.Any(), database.TypeIP, "127.0.0.1", gomock.Any()).
					Return(nil)

				return m
			},
			expectNext: true,
		},
		{
			name: "IP blocked",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
				require.NoError(t, err)
				req.RemoteAddr = "127.0.0.1"
				return req
			},
			assertMock: func(t *testing.T) *mock.MockRateInterface {
				ctrl := gomock.NewController(t)
				m := mock.NewMockRateInterface(ctrl)

				m.EXPECT().
					FindBlocker(gomock.Any(), database.TypeIP, "127.0.0.1").
					Return(true, nil)

				return m
			},
			expectNext: false,
		},
		{
			name: "find IP request limit",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
				require.NoError(t, err)
				req.RemoteAddr = "127.0.0.1"
				return req
			},
			assertMock: func(t *testing.T) *mock.MockRateInterface {
				ctrl := gomock.NewController(t)
				m := mock.NewMockRateInterface(ctrl)

				m.EXPECT().
					FindBlocker(gomock.Any(), database.TypeIP, "127.0.0.1").
					Return(false, nil)
				m.EXPECT().
					FindRequests(gomock.Any(), database.TypeIP, "127.0.0.1").
					Return(5, nil)
				m.EXPECT().
					Block(gomock.Any(), database.TypeIP, "127.0.0.1", gomock.Any()).
					Return(nil)

				return m
			},
			expectNext: false,
		},
		{
			name: "error to get block IP",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
				require.NoError(t, err)
				req.RemoteAddr = "127.0.0.1"
				return req
			},
			assertMock: func(t *testing.T) *mock.MockRateInterface {
				ctrl := gomock.NewController(t)
				m := mock.NewMockRateInterface(ctrl)

				m.EXPECT().
					FindBlocker(gomock.Any(), database.TypeIP, "127.0.0.1").
					Return(false, errors.New("error"))

				return m
			},
			expectNext: false,
		},
		{
			name: "error to find requests by IP",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
				require.NoError(t, err)
				req.RemoteAddr = "127.0.0.1"
				return req
			},
			assertMock: func(t *testing.T) *mock.MockRateInterface {
				ctrl := gomock.NewController(t)
				m := mock.NewMockRateInterface(ctrl)

				m.EXPECT().
					FindBlocker(gomock.Any(), database.TypeIP, "127.0.0.1").
					Return(false, nil)
				m.EXPECT().
					FindRequests(gomock.Any(), database.TypeIP, "127.0.0.1").
					Return(0, errors.New("error"))

				return m
			},
			expectNext: false,
		},
		{
			name: "error to add request by IP",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
				require.NoError(t, err)
				req.RemoteAddr = "127.0.0.1"
				return req
			},
			assertMock: func(t *testing.T) *mock.MockRateInterface {
				ctrl := gomock.NewController(t)
				m := mock.NewMockRateInterface(ctrl)

				m.EXPECT().
					FindBlocker(gomock.Any(), database.TypeIP, "127.0.0.1").
					Return(false, nil)
				m.EXPECT().
					FindRequests(gomock.Any(), database.TypeIP, "127.0.0.1").
					Return(0, nil)
				m.EXPECT().
					NewRequest(gomock.Any(), database.TypeIP, "127.0.0.1", gomock.Any()).
					Return(errors.New("error"))

				return m
			},
			expectNext: false,
		},
		{
			name: "error to block IP",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
				require.NoError(t, err)
				req.RemoteAddr = "127.0.0.1"
				return req
			},
			assertMock: func(t *testing.T) *mock.MockRateInterface {
				ctrl := gomock.NewController(t)
				m := mock.NewMockRateInterface(ctrl)

				m.EXPECT().
					FindBlocker(gomock.Any(), database.TypeIP, "127.0.0.1").
					Return(false, nil)
				m.EXPECT().
					FindRequests(gomock.Any(), database.TypeIP, "127.0.0.1").
					Return(5, nil)
				m.EXPECT().
					Block(gomock.Any(), database.TypeIP, "127.0.0.1", gomock.Any()).
					Return(errors.New("error"))

				return m
			},
			expectNext: false,
		},
		{
			name: "success with API Key",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
				require.NoError(t, err)
				req.Header.Add(HeaderAPIKey, "abc123")
				return req
			},
			assertMock: func(t *testing.T) *mock.MockRateInterface {
				ctrl := gomock.NewController(t)
				m := mock.NewMockRateInterface(ctrl)

				m.EXPECT().
					FindBlocker(gomock.Any(), database.TypeKey, "abc123").
					Return(false, nil)
				m.EXPECT().
					FindRequests(gomock.Any(), database.TypeKey, "abc123").
					Return(0, nil)
				m.EXPECT().
					NewRequest(gomock.Any(), database.TypeKey, "abc123", gomock.Any()).
					Return(nil)

				return m
			},
			expectNext: true,
		},
		{
			name: "IP blocked",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
				require.NoError(t, err)
				req.Header.Add(HeaderAPIKey, "abc123")
				return req
			},
			assertMock: func(t *testing.T) *mock.MockRateInterface {
				ctrl := gomock.NewController(t)
				m := mock.NewMockRateInterface(ctrl)

				m.EXPECT().
					FindBlocker(gomock.Any(), database.TypeKey, "abc123").
					Return(true, nil)

				return m
			},
			expectNext: false,
		},
		{
			name: "find IP request limit",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
				require.NoError(t, err)
				req.Header.Add(HeaderAPIKey, "abc123")
				return req
			},
			assertMock: func(t *testing.T) *mock.MockRateInterface {
				ctrl := gomock.NewController(t)
				m := mock.NewMockRateInterface(ctrl)

				m.EXPECT().
					FindBlocker(gomock.Any(), database.TypeKey, "abc123").
					Return(false, nil)
				m.EXPECT().
					FindRequests(gomock.Any(), database.TypeKey, "abc123").
					Return(5, nil)
				m.EXPECT().
					Block(gomock.Any(), database.TypeKey, "abc123", gomock.Any()).
					Return(nil)

				return m
			},
			expectNext: false,
		},
		{
			name: "error to get block IP",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
				require.NoError(t, err)
				req.Header.Add(HeaderAPIKey, "abc123")
				return req
			},
			assertMock: func(t *testing.T) *mock.MockRateInterface {
				ctrl := gomock.NewController(t)
				m := mock.NewMockRateInterface(ctrl)

				m.EXPECT().
					FindBlocker(gomock.Any(), database.TypeKey, "abc123").
					Return(false, errors.New("error"))

				return m
			},
			expectNext: false,
		},
		{
			name: "error to find requests by IP",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
				require.NoError(t, err)
				req.Header.Add(HeaderAPIKey, "abc123")
				return req
			},
			assertMock: func(t *testing.T) *mock.MockRateInterface {
				ctrl := gomock.NewController(t)
				m := mock.NewMockRateInterface(ctrl)

				m.EXPECT().
					FindBlocker(gomock.Any(), database.TypeKey, "abc123").
					Return(false, nil)
				m.EXPECT().
					FindRequests(gomock.Any(), database.TypeKey, "abc123").
					Return(0, errors.New("error"))

				return m
			},
			expectNext: false,
		},
		{
			name: "error to add request by IP",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
				require.NoError(t, err)
				req.Header.Add(HeaderAPIKey, "abc123")
				return req
			},
			assertMock: func(t *testing.T) *mock.MockRateInterface {
				ctrl := gomock.NewController(t)
				m := mock.NewMockRateInterface(ctrl)

				m.EXPECT().
					FindBlocker(gomock.Any(), database.TypeKey, "abc123").
					Return(false, nil)
				m.EXPECT().
					FindRequests(gomock.Any(), database.TypeKey, "abc123").
					Return(0, nil)
				m.EXPECT().
					NewRequest(gomock.Any(), database.TypeKey, "abc123", gomock.Any()).
					Return(errors.New("error"))

				return m
			},
			expectNext: false,
		},
		{
			name: "error to block IP",
			req: func() *http.Request {
				req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1", nil)
				require.NoError(t, err)
				req.Header.Add(HeaderAPIKey, "abc123")
				return req
			},
			assertMock: func(t *testing.T) *mock.MockRateInterface {
				ctrl := gomock.NewController(t)
				m := mock.NewMockRateInterface(ctrl)

				m.EXPECT().
					FindBlocker(gomock.Any(), database.TypeKey, "abc123").
					Return(false, nil)
				m.EXPECT().
					FindRequests(gomock.Any(), database.TypeKey, "abc123").
					Return(5, nil)
				m.EXPECT().
					Block(gomock.Any(), database.TypeKey, "abc123", gomock.Any()).
					Return(errors.New("error"))

				return m
			},
			expectNext: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := tt.assertMock(t)
			l := limiter{
				database: db,
				cfg: &LimiterConfig{
					KeyLimiter:             true,
					KeyLimit:               5,
					IPLimiter:              true,
					IPLimit:                5,
					RequestLimiterDuration: time.Minute,
					RequestBlockerDuration: time.Minute,
				},
			}

			var nextCalled bool
			next := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				nextCalled = true
			})

			middlewareFunc := l.Do(next)
			middlewareFunc.ServeHTTP(httptest.NewRecorder(), tt.req())
			require.Equal(t, tt.expectNext, nextCalled)
		})
	}
}
