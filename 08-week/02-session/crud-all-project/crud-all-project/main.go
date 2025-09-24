package main

import (
	"crud-all-project/database"
	"crud-all-project/entity"
	"crud-all-project/handler"
	"crud-all-project/repository"
	"crud-all-project/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Ajusta DSN a tu XAMPP: user:password@tcp(127.0.0.1:3306)/nombre_db?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:@tcp(127.0.0.1:3307)/crud-all-project?charset=utf8mb4&parseTime=True&loc=Local"
	if err := database.Connect(dsn); err != nil {
		log.Fatalf("DB connect error: %v", err)
	}

	// Migraciones automáticas (crea tablas si no existen)
	if err := database.DB.AutoMigrate(&entity.Cliente{}, &entity.Servicio{}, &entity.Vehiculo{}); err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}

	// --- wiring: repo -> service -> handler ---
	clienteRepo := repository.NewClienteRepository()
	clienteSvc := service.NewClienteService(clienteRepo)
	clienteH := handler.NewClienteHandler(clienteSvc)

	servicioRepo := repository.NewServicioRepository()
	servicioSvc := service.NewServicioService(servicioRepo)
	servicioH := handler.NewServicioHandler(servicioSvc)

	vehRepo := repository.NewVehiculoRepository()
	vehSvc := service.NewVehiculoService(vehRepo)
	vehH := handler.NewVehiculoHandler(vehSvc)

	// Router
	r := mux.NewRouter()
	// CORS: si lo necesitas, añade middleware CORS aquí (por tu Angular)
	clienteH.RegisterRoutes(r)
	servicioH.RegisterRoutes(r)
	vehH.RegisterRoutes(r)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
