package dominio

import (
	"strconv"
)

// TypeGORM es el modelo de type para GORM de Type
type TypeGORM struct {
	ID          uint   `gorm:"primaryKey:autoIncrement" json:"id"`
	Nombre      string `gorm:"type:varchar(255);not null"`
	Descripcion string `gorm:"type:varchar(255);not null"`
}

// TableName especifica el nombre de la tabla para TypeGORM
func (TypeGORM) TableName() string {
	return "tipos"
}

func (typeGORM *TypeGORM) ToGQL() (*Tipo, error) {

	return &Tipo{
		ID:          strconv.Itoa(int(typeGORM.ID)),
		Nombre:      typeGORM.Nombre,
		Descripcion: typeGORM.Descripcion,
	}, nil
}
