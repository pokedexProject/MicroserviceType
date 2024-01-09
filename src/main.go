package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	repository "github.com/pokedexProject/MicroserviceType/adapters"
	service "github.com/pokedexProject/MicroserviceType/aplication"
	"github.com/pokedexProject/MicroserviceType/database"
	pb "github.com/pokedexProject/MicroserviceType/proto"
	"google.golang.org/grpc"
)

func main() {

	db := database.Connect()
	database.EjecutarMigraciones(db.GetConn())
	typeRepository := repository.NewTypeRepository(db)
	typeService := service.NewTypeService(typeRepository)
	// Configura el servidor gRPC
	//este servidor está escuchando en el puerto 50052
	//y se encarga de registrar el servicio de entrenadores
	grpcServe := grpc.NewServer()
	// Registra el servicio de entrenadores en el servidor gRPC
	pb.RegisterTypeServiceServer(grpcServe, typeService)

	// Define el puerto en el que se ejecutará el servidor gRPC
	port := "50053"
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Printf("Server listening on port %s...\n", port)

	// Inicia el servidor gRPC en segundo plano
	go func() {
		if err := grpcServe.Serve(listen); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Espera una señal para detener el servidor gRPC
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch

	fmt.Println("Shutting down the server...")

	// Detén el servidor gRPC de manera segura
	grpcServe.GracefulStop()
}
