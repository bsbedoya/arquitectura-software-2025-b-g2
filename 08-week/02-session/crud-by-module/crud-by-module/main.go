package main

import (
	clienteEntity "crud-by-module/clientes/entity"
	"crud-by-module/clientes/handler"
	clienteRepoPkg "crud-by-module/clientes/repository"
	clienteServicePkg "crud-by-module/clientes/service"

	vehiculoEntity "crud-by-module/vehiculos/entity"
	vehiculoHandlerPkg "crud-by-module/vehiculos/handler"
	vehiculoRepoPkg "crud-by-module/vehiculos/repository"
	vehiculoServicePkg "crud-by-module/vehiculos/service"

	servicioEntity "crud-by-module/servicios/entity"
	servicioHandlerPkg "crud-by-module/servicios/handler"
	servicioRepoPkg "crud-by-module/servicios/repository"
	servicioServicePkg "crud-by-module/servicios/service"

	"crud-by-module/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// 1. Conectar a la BD (sin return, usa log.Fatal dentro de Connect si falla)
	dsn := "root:@tcp(127.0.0.1:3307)/crud-?charset=utf8mb4&parseTime=True&loc=Local"
	database.Connect(dsn)

	// 2. Migraciones automáticas (crea tablas si no existen)
	if err := database.DB.AutoMigrate(
		&clienteEntity.Cliente{},
		&vehiculoEntity.Vehiculo{},
		&servicioEntity.Servicio{},
	); err != nil {
		log.Fatalf("AutoMigrate error: %v", err)
	}

	// ================== CLIENTES ==================
	clienteRepo := clienteRepoPkg.NewClienteRepository()
	clienteService := clienteServicePkg.NewClienteService(clienteRepo)
	clienteHandler := handler.NewClienteHandler(clienteService)

	// ================== VEHICULOS ==================
	vehiculoRepo := vehiculoRepoPkg.NewVehiculoRepository(database.DB)
	vehiculoService := vehiculoServicePkg.NewVehiculoService(vehiculoRepo)
	vehiculoHandler := vehiculoHandlerPkg.NewVehiculoHandler(vehiculoService)

	// ================== SERVICIOS ==================
	servicioRepo := servicioRepoPkg.NewServicioRepository()
	servicioService := servicioServicePkg.NewServicioService(servicioRepo)
	servicioHandler := servicioHandlerPkg.NewServicioHandler(servicioService)

	// 3. Configurar router
	r := mux.NewRouter()
	clienteHandler.RegisterRoutes(r)
	vehiculoHandler.RegisterRoutes(r)
	servicioHandler.RegisterRoutes(r)

	// 4. Iniciar servidor
	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
