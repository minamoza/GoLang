package http

import (
	"go/internal/redis_cache"
	"go/internal/store"
)

type ServerOption func(srv *Server)

func WithAddress(address string) ServerOption {
	return func(srv *Server) {
		srv.Address = address
	}
}

func WithStore(store store.Store) ServerOption {
	return func(srv *Server) {
		srv.store = store
	}
}

func WithCache(cache redis_cache.Cache) ServerOption {
	return func(srv *Server) {
		srv.cache = cache
	}
}