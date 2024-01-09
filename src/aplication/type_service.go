package service

import (
	"context"
	"fmt"

	model "github.com/pokedexProject/MicroserviceType/dominio"
	repository "github.com/pokedexProject/MicroserviceType/ports"
	pb "github.com/pokedexProject/MicroserviceType/proto"
)

// este servicio implementa la interfaz TypeServiceServer
// que se genera a partir del archivo proto
type TypeService struct {
	pb.UnimplementedTypeServiceServer
	repo repository.TypeRepository
}

func NewTypeService(repo repository.TypeRepository) *TypeService {
	return &TypeService{repo: repo}
}

func (s *TypeService) CreateType(ctx context.Context, req *pb.CreateTypeRequest) (*pb.CreateTypeResponse, error) {

	crearTipoInput := model.CrearTipoInput{
		Nombre:      req.GetNombre(),
		Descripcion: req.GetDescripcion(),
	}
	u, err := s.repo.CrearTipo(crearTipoInput)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Tipo creado: %v", u)
	response := &pb.CreateTypeResponse{
		Id:          u.ID,
		Nombre:      u.Nombre,
		Descripcion: u.Descripcion,
	}
	fmt.Printf("Tipo creado: %v", response)
	return response, nil
}

func (s *TypeService) GetType(ctx context.Context, req *pb.GetTypeRequest) (*pb.GetTypeResponse, error) {
	u, err := s.repo.Tipo(req.GetId())
	if err != nil {
		return nil, err
	}
	response := &pb.GetTypeResponse{
		Id:          u.ID,
		Nombre:      u.Nombre,
		Descripcion: u.Descripcion,
	}
	return response, nil
}

func (s *TypeService) ListTypes(ctx context.Context, req *pb.ListTypesRequest) (*pb.ListTypesResponse, error) {
	tipos, err := s.repo.Tipos()
	if err != nil {
		return nil, err
	}
	var response []*pb.Type
	for _, u := range tipos {
		tipo := &pb.Type{
			Id:          u.ID,
			Nombre:      u.Nombre,
			Descripcion: u.Descripcion,
		}
		response = append(response, tipo)
	}

	return &pb.ListTypesResponse{Types: response}, nil
}

func (s *TypeService) UpdateTipo(ctx context.Context, req *pb.UpdateTypeRequest) (*pb.UpdateTypeResponse, error) {
	nombre := req.GetNombre()
	descripcion := req.GetDescripcion()

	fmt.Printf("Nombre: %v", nombre)
	actualizarTipoInput := &model.ActualizarTipoInput{
		Nombre:      &nombre,
		Descripcion: &descripcion,
	}
	fmt.Printf("Tipo actualizado input: %v", actualizarTipoInput)
	u, err := s.repo.ActualizarTipo(req.GetId(), actualizarTipoInput)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Tipo actualizado: %v", u)
	response := &pb.UpdateTypeResponse{
		Id:          u.ID,
		Nombre:      u.Nombre,
		Descripcion: u.Descripcion,
	}
	return response, nil
}

func (s *TypeService) DeleteType(ctx context.Context, req *pb.DeleteTypeRequest) (*pb.DeleteTypeResponse, error) {
	respuesta, err := s.repo.EliminarTipo(req.GetId())
	if err != nil {
		return nil, err
	}
	response := &pb.DeleteTypeResponse{
		Mensaje: respuesta.Mensaje,
	}
	return response, nil
}
