package middlewares

import (
	"github.com/matheusmhmelo/FullCycle-rate-limiter/internal/infra/database"
	"net/http"
	"strings"
	"time"
)

const HeaderAPIKey = "API_KEY"

type RateLimiter interface {
	Do(next http.Handler) http.Handler
}

type LimiterConfig struct {
	KeyLimiter             bool
	KeyLimit               int
	IPLimiter              bool
	IPLimit                int
	RequestLimiterDuration time.Duration
	RequestBlockerDuration time.Duration
}

type limiter struct {
	database database.RateInterface
	cfg      *LimiterConfig
}

func NewRateLimiter(db database.RateInterface, cgf *LimiterConfig) RateLimiter {
	return &limiter{
		database: db,
		cfg:      cgf,
	}
}

func (l *limiter) Do(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		key := r.Header.Get(HeaderAPIKey)

		if l.cfg.KeyLimiter && key != "" {
			allowed, err := l.isRequestAllowed(r, key, l.cfg.KeyLimit, database.TypeKey)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("error: " + err.Error()))
				return
			}

			if !allowed {
				handleBlockedRequest(w)
				return
			}
		} else if l.cfg.IPLimiter {
			allowed, err := l.isRequestAllowed(r, ip, l.cfg.IPLimit, database.TypeIP)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("error: " + err.Error()))
				return
			}

			if !allowed {
				handleBlockedRequest(w)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func (l *limiter) isRequestAllowed(
	r *http.Request,
	value string,
	limit int,
	rateType database.LimiterType,
) (bool, error) {
	// Check if request is blocked
	isBlocked, err := l.database.FindBlocker(r.Context(), rateType, value)
	if err != nil {
		return false, err
	}
	if isBlocked {
		return false, nil
	}

	// Get current number of requests
	requests, err := l.database.FindRequests(r.Context(), rateType, value)
	if err != nil {
		return false, err
	}

	// Check if requests reached the limit
	if requests >= limit {
		err = l.database.Block(r.Context(), rateType, value, l.cfg.RequestBlockerDuration)
		if err != nil {
			return false, err
		}

		// Block request because reached the limit
		return false, nil
	}

	// Increment requests number
	err = l.database.NewRequest(r.Context(), rateType, value, l.cfg.RequestLimiterDuration)
	if err != nil {
		return false, err
	}

	// Allow request to proceed
	return true, nil
}

func handleBlockedRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusTooManyRequests)
	w.Write([]byte("you have reached the maximum number of requests or actions allowed within a certain time frame"))
}
