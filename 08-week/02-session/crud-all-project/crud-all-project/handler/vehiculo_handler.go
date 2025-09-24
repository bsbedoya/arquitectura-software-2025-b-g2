package handler

import (
	"crud-all-project/entity"
	"crud-all-project/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type VehiculoHandler struct {
	Service service.VehiculoService
}

func NewVehiculoHandler(s service.VehiculoService) *VehiculoHandler {
	return &VehiculoHandler{Service: s}
}

func (h *VehiculoHandler) RegisterRoutes(r *mux.Router) {
	sr := r.PathPrefix("/vehiculos").Subrouter()
	sr.HandleFunc("", h.GetAll).Methods("GET")
	sr.HandleFunc("/{id}", h.GetByID).Methods("GET")
	sr.HandleFunc("", h.Create).Methods("POST")
	sr.HandleFunc("/{id}", h.Update).Methods("PUT")
	sr.HandleFunc("/{id}", h.Delete).Methods("DELETE")
}

func (h *VehiculoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.Service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(list)
}

func (h *VehiculoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)
	item, err := h.Service.FindByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *VehiculoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var payload entity.Vehiculo
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

func (h *VehiculoHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)
	existing, err := h.Service.FindByID(uint(id))
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	var payload entity.Vehiculo
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	existing.Placa = payload.Placa
	existing.TipoVehiculo = payload.TipoVehiculo
	existing.IDCliente = payload.IDCliente

	updated, err := h.Service.Update(existing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (h *VehiculoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idStr)
	if err := h.Service.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
