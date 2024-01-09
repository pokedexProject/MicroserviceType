package dominio

type ActualizarTypeInput struct {
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

type CrearTypeInput struct {
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

type EliminacionType struct {
	Mensaje string `json:"mensaje"`
}

type Type struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

func (Type) IsEntity() {}
