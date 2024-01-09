package database

import (
	"log"

	model "github.com/pokedexProject/MicroserviceType/src/dominio"
	"gorm.io/gorm"
)

// EjecutarMigraciones realiza todas las migraciones necesarias en la base de datos.
func EjecutarMigraciones(db *gorm.DB) {

	db.AutoMigrate(&model.TypeGORM{})

	log.Println("Migraciones completadas")
}
