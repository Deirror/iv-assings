package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Response struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Handler struct {
	custDb map[int]string

	mux *sync.RWMutex
}

func NewHandler() *Handler {
	accounts := []string{"Alex", "Denis", "Stoyan",
		"Deirror", "Lili", "Bob", "Olq",
		"Jan", "Zahari", "Ivo"}
	db := make(map[int]string)
	for i, acc := range accounts {
		db[i] = acc
	}

	return &Handler{
		custDb: db,
		mux:    &sync.RWMutex{},
	}
}

func (h *Handler) HandleGetCustomerMetadata(w http.ResponseWriter, r *http.Request) {
	r.URL.Query()
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id was empty", http.StatusBadRequest)
		return
	}

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "id was not a number", http.StatusBadRequest)
		return
	}

	h.mux.RLock()
	name, ok := h.custDb[parsedId]
	h.mux.RUnlock()

	if !ok {
		http.Error(w, "id was not found", http.StatusNotFound)
		return
	}

	log.Printf("GET /customer?id=%s\n", id)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Id:   parsedId,
		Name: name,
	})
}

func (h *Handler) HandleGetAllCustomerMetadata(w http.ResponseWriter, r *http.Request) {
	list := make([]Response, len(h.custDb))

	h.mux.RLock()
	defer h.mux.RUnlock()

	idx := 0
	for key, val := range h.custDb {
		resp := Response{
			Id:   key,
			Name: val,
		}
		list[idx] = resp
		idx++
	}

	log.Printf("GET /customers\n")

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func main() {
	addr := "3000"

	mux := &http.ServeMux{}
	handler := NewHandler()

	mux.HandleFunc("GET /customer", handler.HandleGetCustomerMetadata)
	mux.HandleFunc("GET /customers", handler.HandleGetAllCustomerMetadata)

	log.Println("Server listening on port ", addr)
	log.Fatal(http.ListenAndServe(":"+addr, mux))
}
