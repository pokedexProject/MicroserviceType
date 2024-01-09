package ports

import (
	model "github.com/pokedexProject/MicroserviceType/dominio"
)

// puerto de salida
type TypeRepository interface {
	CrearTipo(input model.CrearTipoInput) (*model.Tipo, error)
	Tipo(id string) (*model.Tipo, error)
	ActualizarTipo(id string, input *model.ActualizarTipoInput) (*model.Tipo, error)
	EliminarTipo(id string) (*model.EliminacionTipo, error)
	Tipos() ([]*model.Tipo, error)
}
