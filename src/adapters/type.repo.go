package adapters

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/pokedexProject/MicroserviceType/database"
	model "github.com/pokedexProject/MicroserviceType/dominio"
	"github.com/pokedexProject/MicroserviceType/ports"

	"gorm.io/gorm"
)

/**
* Es un adaptador de salida
usuario
*/

type typeRepository struct {
	db             *database.DB
	activeSessions map[string]string
}

func NewTypeRepository(db *database.DB) ports.TypeRepository {
	return &typeRepository{
		db:             db,
		activeSessions: make(map[string]string),
	}
}

func ToJSON(obj interface{}) (string, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonData), err
}

// ObtenerTrabajo obtiene un trabajo por su ID.
func (ur *typeRepository) Tipo(id string) (*model.Tipo, error) {
	if id == "" {
		return nil, errors.New("El ID de tipo es requerido")
	}

	var tipoGORM model.TypeGORM
	//result := ur.db.GetConn().First(&usuarioGORM, id)
	result := ur.db.GetConn().First(&tipoGORM, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		log.Printf("Error al obtener el tipo con ID %s: %v", id, result.Error)
		return nil, result.Error
	}

	return tipoGORM.ToGQL()
}

// Tipos obtiene todos los usuarios de la base de datos.
func (ur *typeRepository) Tipos() ([]*model.Tipo, error) {
	var tiposGORM []model.TypeGORM
	result := ur.db.GetConn().Find(&tiposGORM)

	if result.Error != nil {
		log.Printf("Error al obtener los tipos: %v", result.Error)
		return nil, result.Error
	}

	var tipos []*model.Tipo
	for _, tipoGORM := range tiposGORM {
		tipo, _ := tipoGORM.ToGQL()
		tipos = append(tipos, tipo)
	}

	// usuariosJSON, err := json.Marshal(usuarios)
	// if err != nil {
	// 	log.Printf("Error al convertir usuarios a JSON: %v", err)
	// 	return "[]", err
	// }
	// return ToJSON(usuarios)
	return tipos, nil
}
func (ur *typeRepository) CrearTipo(input model.CrearTipoInput) (*model.Tipo, error) {

	tipoGORM :=
		&model.TypeGORM{
			Nombre:      input.Nombre,
			Descripcion: input.Descripcion,
		}
	result := ur.db.GetConn().Create(&tipoGORM)
	if result.Error != nil {
		log.Printf("Error al crear el tipo: %v", result.Error)
		return nil, result.Error
	}

	response, err := tipoGORM.ToGQL()
	return response, err
}

func (ur *typeRepository) ActualizarTipo(id string, input *model.ActualizarTipoInput) (*model.Tipo, error) {
	var tipoGORM model.TypeGORM
	if id == "" {
		return nil, errors.New("El ID de tipo es requerido")
	}

	result := ur.db.GetConn().First(&tipoGORM, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Tipo con ID %s no encontrado", id)
		}
		return nil, result.Error
	}

	// Solo actualiza los campos proporcionados
	if input.Nombre != nil {
		tipoGORM.Nombre = *input.Nombre
	}
	if input.Descripcion != nil {
		tipoGORM.Descripcion = *input.Descripcion
	}
	result = ur.db.GetConn().Save(&tipoGORM)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Printf("Tipo actualizado: %v", tipoGORM)
	return tipoGORM.ToGQL()
}

// EliminarTipo elimina un tipo de la base de datos por su ID.
func (ur *typeRepository) EliminarTipo(id string) (*model.EliminacionTipo, error) {
	// Intenta buscar el tipo por su ID
	var tipoGORM model.TypeGORM
	result := ur.db.GetConn().First(&tipoGORM, id)

	if result.Error != nil {
		// Manejo de errores
		if result.Error == gorm.ErrRecordNotFound {
			// El tipo no se encontró en la base de datos
			response := &model.EliminacionTipo{
				Mensaje: "El tipo no existe",
			}
			return response, result.Error

		}
		log.Printf("Error al buscar el tipo con ID %s: %v", id, result.Error)
		response := &model.EliminacionTipo{
			Mensaje: "Error al buscar el tipo",
		}
		return response, result.Error
	}

	// Elimina el tipo de la base de datos
	result = ur.db.GetConn().Delete(&tipoGORM, id)

	if result.Error != nil {
		log.Printf("Error al eliminar el tipo con ID %s: %v", id, result.Error)
		response := &model.EliminacionTipo{
			Mensaje: "Error al eliminar el tipo",
		}
		return response, result.Error
	}

	// Éxito al eliminar el tipo
	response := &model.EliminacionTipo{
		Mensaje: "Tipo eliminado con éxito",
	}
	return response, result.Error

}
