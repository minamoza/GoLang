package main

import (
	"context"
	"go/internal/http"
	"go/internal/redis_cache"
	"go/internal/store/postgres"
	_ "github.com/go-redis/redis/v8"
)

func main() {
	urlExample := "postgres://localhost:5431/mina"
	store := postgres.NewDB()
	if err := store.Connect(urlExample); err != nil {
		panic(err)
	}
	defer store.Close()

	cache := redis_cache.NewRedisCache("localhost:6379", 1, 1800)

	srv := http.NewServer(
		context.Background(),
		http.WithAddress(":8080"),
		http.WithStore(store),
		http.WithCache(cache),
	)
	if err := srv.Run(); err != nil {
		panic(err)
	}

	srv.WaitForGracefulTermination()
}
