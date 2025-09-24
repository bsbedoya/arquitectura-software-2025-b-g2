package handler

import (
	"crud-all-project/entity"
	"crud-all-project/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ServicioHandler struct {
	Service service.ServicioService
}

func NewServicioHandler(s service.ServicioService) *ServicioHandler {
	return &ServicioHandler{Service: s}
}

func (h *ServicioHandler) RegisterRoutes(r *mux.Router) {
	sr := r.PathPrefix("/servicios").Subrouter()
	sr.HandleFunc("", h.GetAll).Methods("GET")
	sr.HandleFunc("/{id}", h.GetByID).Methods("GET")
	sr.HandleFunc("", h.Create).Methods("POST")
	sr.HandleFunc("/{id}", h.Update).Methods("PUT")
	sr.HandleFunc("/{id}", h.Delete).Methods("DELETE")
}

func (h *ServicioHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.Service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(list)
}

func (h *ServicioHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)
	item, err := h.Service.FindByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *ServicioHandler) Create(w http.ResponseWriter, r *http.Request) {
	var payload entity.Servicio
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	created, err := h.Service.Save(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *ServicioHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)
	existing, err := h.Service.FindByID(uint(id))
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	var payload entity.Servicio
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	existing.NombreServicio = payload.NombreServicio
	existing.Descripcion = payload.Descripcion
	existing.Precio = payload.Precio
	existing.DuracionEstimada = payload.DuracionEstimada

	updated, err := h.Service.Save(existing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (h *ServicioHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)
	if err := h.Service.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
