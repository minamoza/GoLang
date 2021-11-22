package http

import (
	"context"
	"encoding/json"
	"fmt"
	"go/internal/models"
	"go/internal/redis_cache"
	"go/internal/store"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Server struct {
	ctx         context.Context
	idleConnsCh chan struct{}
	store       store.Store
	cache       redis_cache.Cache

	Address string
}

func NewServer(ctx context.Context, opts ...ServerOption) *Server {
	srv := &Server{
		ctx:         ctx,
		idleConnsCh: make(chan struct{}),
	}
	for _, opt := range opts {
		opt(srv)
	}

	return srv
}

func (s *Server) basicHandler() chi.Router {
	r := chi.NewRouter()

	dishResourse := NewDishResource(s.store, s.cache)
	r.Mount("/dishes", dishResourse.Routes())

	// REST
	// сущность/идентификатор
	// /electronics/orders
	// /electronics/phones
	r.Post("/orders", func(w http.ResponseWriter, r *http.Request) {
		order := new(models.Order)
		if err := json.NewDecoder(r.Body).Decode(order); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Order().Create(r.Context(), order)
	})
	r.Get("/orders", func(w http.ResponseWriter, r *http.Request) {
		orders, err := s.store.Order().All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, orders)
	})
	r.Get("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		order, err := s.store.Order().ByID(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, order)
	})
	r.Put("/orders", func(w http.ResponseWriter, r *http.Request) {
		order := new(models.Order)
		if err := json.NewDecoder(r.Body).Decode(order); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Order().Update(r.Context(), order)
	})
	r.Delete("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Order().Delete(r.Context(), id)
	})

	r.Post("/restaurants", func(w http.ResponseWriter, r *http.Request) {
		res := new(models.Restaurant)
		if err := json.NewDecoder(r.Body).Decode(res); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Restaurant().Create(r.Context(), res)
	})
	r.Get("/restaurants", func(w http.ResponseWriter, r *http.Request) {
		restaurants, err := s.store.Restaurant().All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, restaurants)
	})
	r.Get("/restaurants/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		restaurant, err := s.store.Restaurant().ByID(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, restaurant)
	})
	r.Put("/restaurants", func(w http.ResponseWriter, r *http.Request) {
		restaurant := new(models.Restaurant)
		if err := json.NewDecoder(r.Body).Decode(restaurant); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Restaurant().Update(r.Context(), restaurant)
	})
	r.Delete("/rerstaurants/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Restaurant().Delete(r.Context(), id)
	})

	r.Post("/categories", func(w http.ResponseWriter, r *http.Request) {
		categorie := new(models.Category)
		if err := json.NewDecoder(r.Body).Decode(categorie); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Categories().Create(r.Context(), categorie)
	})
	r.Get("/categories", func(w http.ResponseWriter, r *http.Request) {
		categories, err := s.store.Categories().All(r.Context())
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, categories)
	})
	r.Get("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		categorie, err := s.store.Categories().ByID(r.Context(), id)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		render.JSON(w, r, categorie)
	})
	r.Put("/categories", func(w http.ResponseWriter, r *http.Request) {
		categorie := new(models.Category)
		if err := json.NewDecoder(r.Body).Decode(categorie); err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Categories().Update(r.Context(), categorie)
	})
	r.Delete("/categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Fprintf(w, "Unknown err: %v", err)
			return
		}

		s.store.Categories().Delete(r.Context(), id)
	})

	return r
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:         s.Address,
		Handler:      s.basicHandler(),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 30,
	}
	go s.ListenCtxForGT(srv)

	log.Println("[HTTP] Server running on", s.Address)
	return srv.ListenAndServe()
}

func (s *Server) ListenCtxForGT(srv *http.Server) {
	<-s.ctx.Done() // блокируемся, пока контекст приложения не отменен

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("[HTTP] Got err while shutting down^ %v", err)
	}

	log.Println("[HTTP] Proccessed all idle connections")
	close(s.idleConnsCh)
}

func (s *Server) WaitForGracefulTermination() {
	// блок до записи или закрытия канала
	<-s.idleConnsCh
}
