package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"crud-by-module/clientes/entity"
	"crud-by-module/clientes/service"

	"github.com/gorilla/mux"
)

type ClienteHandler struct {
	service service.ClienteService
}

func (h *ClienteHandler) RegisterRoutes(r *mux.Router) {
	s := r.PathPrefix("/clientes").Subrouter()
	s.HandleFunc("", h.GetAll).Methods("GET")              // GET /clientes
	s.HandleFunc("/{id:[0-9]+}", h.GetByID).Methods("GET") // GET /clientes/{id}
	s.HandleFunc("", h.Create).Methods("POST")             // POST /clientes
}

func NewClienteHandler(s service.ClienteService) *ClienteHandler {
	return &ClienteHandler{service: s}
}

// GET /clientes
func (h *ClienteHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	clientes, err := h.service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(clientes)
}

// GET /clientes/{id}
func (h *ClienteHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	cliente, err := h.service.FindByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(cliente)
}

// POST /clientes
func (h *ClienteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c entity.Cliente
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	cliente, err := h.service.Save(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cliente)
}
