package http

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go/internal/models"
	"go/internal/redis_cache"
	"go/internal/store"
	"net/http"
	"strconv"
)

type DishResource struct {
	store store.Store
	cache redis_cache.Cache
}

func NewDishResource(store store.Store, cache redis_cache.Cache) *DishResource {
	return &DishResource{
		store: store,
		cache: cache,
	}
}

func (cr *DishResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", cr.CreateDish)
	r.Get("/", cr.AllDishes)
	r.Get("/{id}", cr.ByID)
	r.Put("/", cr.UpdateDish)
	r.Delete("/{id}", cr.DeleteDish)

	return r
}

func (cr *DishResource) CreateDish(w http.ResponseWriter, r *http.Request) {
	Dish := new(models.Dish)
	if err := json.NewDecoder(r.Body).Decode(Dish); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if err := cr.store.Dish().Create(r.Context(), Dish); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}



	w.WriteHeader(http.StatusCreated)
}

func (cr *DishResource) AllDishes(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	filter := &models.DishFilter{}

	searchQuery := queryValues.Get("query")
	if searchQuery != "" {
		DishesFromCache, err := cr.cache.Dish().Get(r.Context(), searchQuery)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Cache err: %v", err)
			return
		}

		if DishesFromCache != nil {
			render.JSON(w, r, DishesFromCache)
			return
		}

		filter.Query = &searchQuery
	}

	dishes, err := cr.store.Dish().All(r.Context(), filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if searchQuery != "" && len(dishes) > 0 {
		err = cr.cache.Dish().Set(r.Context(), searchQuery, dishes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Cache err: %v", err)
			return
		}
	}

	render.JSON(w, r, dishes)
}

func (cr *DishResource) ByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	dish, err := cr.store.Dish().ByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	render.JSON(w, r, dish)
}

func (cr *DishResource) UpdateDish(w http.ResponseWriter, r *http.Request) {
	dish := new(models.Dish)
	if err := json.NewDecoder(r.Body).Decode(dish); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	err := cr.store.Dish().Update(r.Context(), dish)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}
}

func (cr *DishResource) DeleteDish(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	err = cr.store.Dish().Delete(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}
}