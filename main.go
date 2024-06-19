package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/matheusmhmelo/FullCycle-rate-limiter/internal/infra/webserver/middlewares"
	"github.com/redis/go-redis/v9"
	"net/http"

	"github.com/matheusmhmelo/FullCycle-rate-limiter/configs"
	"github.com/matheusmhmelo/FullCycle-rate-limiter/internal/infra/database"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.DBHost, config.DBPort),
		Password: config.DBPassword,
		DB:       0, // use default DB
	})
	rateDB := database.NewRateLimitStorage(rdb)
	rateMiddleware := middlewares.NewRateLimiter(rateDB, &middlewares.LimiterConfig{
		KeyLimiter:             config.KeyLimiter,
		KeyLimit:               config.KeyLimit,
		IPLimiter:              config.IPLimiter,
		IPLimit:                config.IPLimit,
		RequestLimiterDuration: config.RequestLimiterDuration,
		RequestBlockerDuration: config.RequestBlockerDuration,
	})

	r := chi.NewRouter()
	r.Use(rateMiddleware.Do)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	})

	fmt.Println("Running server at port 8080...")
	http.ListenAndServe(":8080", r)
}
