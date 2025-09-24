package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"crud-by-module/vehiculos/entity"
	"crud-by-module/vehiculos/service"

	"github.com/gorilla/mux"
)

type VehiculoHandler struct {
	service service.VehiculoService
}

func (h *VehiculoHandler) RegisterRoutes(r *mux.Router) {
	vr := r.PathPrefix("/vehiculos").Subrouter()
	vr.HandleFunc("", h.GetAll).Methods("GET")
	vr.HandleFunc("/{id}", h.GetByID).Methods("GET")
	vr.HandleFunc("", h.Create).Methods("POST")
	vr.HandleFunc("/{id}", h.Update).Methods("PUT")
	vr.HandleFunc("/{id}", h.Delete).Methods("DELETE")
}

func NewVehiculoHandler(s service.VehiculoService) *VehiculoHandler {
	return &VehiculoHandler{service: s}
}

// GET /vehiculos
func (h *VehiculoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	vehiculos, err := h.service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(vehiculos)
}

// GET /vehiculos/{id}
func (h *VehiculoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	vehiculo, err := h.service.FindByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(vehiculo)
}

// POST /vehiculos
func (h *VehiculoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var v entity.Vehiculo
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	vehiculo, err := h.service.Save(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(vehiculo)
}

// PUT /vehiculos/{id}
func (h *VehiculoHandler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var v entity.Vehiculo
	err = json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	v.ID = uint(id) // Aseguramos que el ID sea el correcto

	updatedVehiculo, err := h.service.Update(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedVehiculo)
}

// DELETE /vehiculos/{id}
func (h *VehiculoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
