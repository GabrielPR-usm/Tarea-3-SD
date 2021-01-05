package main

import (
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Modificaciones"

)



func main() {

	lis, err := net.Listen("tcp", ":9500")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := ModificarZF.Server{}

	grpcServer := grpc.NewServer()

	//REGISTRO DE SERVICIOS
	ModificarZF.RegisterMOdificarZFRequestServiceServer(grpcServer, &s)
	//FIN REGISTRO DE SERVICIOS
	fmt.Println("BROKER ESCUCHANDO")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
