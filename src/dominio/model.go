package dominio

type ActualizarTipoInput struct {
	Nombre      *string `json:"nombre"`
	Descripcion *string `json:"descripcion"`
}

type CrearTipoInput struct {
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

type EliminacionTipo struct {
	Mensaje string `json:"mensaje"`
}

type Tipo struct {
	ID          string `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}

func (Tipo) IsEntity() {}
