package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect abre la conexión a MySQL usando DSN.
// Ejemplo DSN para XAMPP (ajusta user/password/dbname):
// root:password@tcp(127.0.0.1:3306)/carwash_db?charset=utf8mb4&parseTime=True&loc=Local
func Connect(dsn string) error {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("gorm.Open error: %w", err)
	}
	return nil
}
