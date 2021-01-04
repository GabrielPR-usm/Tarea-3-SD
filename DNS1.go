package main

import (
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Consultas"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Compartir"
	"github.com/GabrielPR-usm/Tarea-3-SD/Mensajes/Modificaciones"

)

func consultasServer(lis net.Listener) {
	s := Consulta.Server{}
	grpcServer := grpc.NewServer()
	Consulta.RegisterConsultaBrokerServiceServer(grpcServer, &s)
	fmt.Println("DNS 1 Escuchando a las consultas de los clientes")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve consultas: %s", err)
	}
}

func modificacionesServer(lis net.Listener) {
	s := ModificarZF.Server{}
	grpcServer := grpc.NewServer()
	Consulta.RegisterModificarZFEnDNSServiceServer(grpcServer, &s)
	fmt.Println("DNS 1 Escuchando las modificaciones del admin")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve modificaciones: %s", err)
	}
}

func compartirServer(lis net.Listener) {
	s := CompartirZF.Server{}
	grpcServer := grpc.NewServer()
	Consulta.RegisterCompartirZFServiceServer(grpcServer, &s)
	fmt.Println("DNS 1 Escuchando otros DNS para hacer consistencia")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve consistencia: %s", err)
	}
}

func main() {

	lis, err := net.Listen("tcp", ":1000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func(){
		consultasServer(lis)
	}()

	go func(){
		modificacionesServer(lis)
	}()

	compartirServer(lis)
}
